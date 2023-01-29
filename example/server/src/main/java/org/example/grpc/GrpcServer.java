package org.example.grpc;

import com.thingworx.types.InfoTable;
import com.thingworx.types.collections.ValueCollection;
import example.protobuf.ExampleGrpc;
import example.protobuf.ExampleOuterClass;
import io.grpc.Server;
import io.grpc.ServerBuilder;
import io.grpc.Status;
import io.grpc.stub.StreamObserver;
import org.example.data.DataProcessor;
import org.example.data.ThingworxDataInterface;

import java.io.IOException;
import java.util.UUID;
import java.util.concurrent.ThreadLocalRandom;
import java.util.concurrent.TimeUnit;

public class GrpcServer {
    private static Server server;

    private static DataProcessor dataProcessor;

    public GrpcServer(ThingworxDataInterface thingworxDataInterface) {
        dataProcessor = new DataProcessor(thingworxDataInterface);
    }

    public void startGrpcServer() throws InterruptedException, IOException {
        int port = 50051;
        server = ServerBuilder
                .forPort(port)
                .addService(new ExampleImpl(dataProcessor))
                .maxConnectionIdle(120, TimeUnit.SECONDS)
                .keepAliveTime(60, TimeUnit.SECONDS)
                .keepAliveTimeout(10, TimeUnit.SECONDS)
                .build()
                .start();

        Runtime.getRuntime().addShutdownHook(new Thread(() -> {
            System.out.println("*** shutting down gRPC server since JVM is shutting down");
            try{
                stopGrpcServer();
            } catch (InterruptedException e) {
                e.printStackTrace(System.err);
                System.out.println(e.getMessage());
            }
            System.out.println("*** server shut down");
        }));
    }

    public void stopGrpcServer() throws InterruptedException {
        if(server != null) {
            System.out.println("gRPC Server stopped");
            server.shutdown().awaitTermination(30, TimeUnit.SECONDS);
        }
    }

    public void blockUntilShutdown() throws InterruptedException {
        if(server != null) {
            server.awaitTermination();
        }
    }

    public boolean isServerShutdown() {
        return server.isShutdown();
    }

    static class ExampleImpl extends ExampleGrpc.ExampleImplBase {
        DataProcessor dataProcessor;

        ExampleImpl(DataProcessor dataProcessor) {
            this.dataProcessor = dataProcessor;
        }
        @Override
        public void getIoTData(ExampleOuterClass.GetIoTDataReq request, StreamObserver<ExampleOuterClass.GetIoTDataRes> responseObserver) {
            if(request.getDevice().getName().isEmpty()) {
                responseObserver.onError(Status.INTERNAL
                        .withDescription("Get IoT Data. Missing device name")
                        .asRuntimeException());
            }
            ExampleOuterClass.GetIoTDataRes.Builder response = ExampleOuterClass.GetIoTDataRes.newBuilder();
            System.out.println("Get Iot Data for device name: " + request.getDevice().getName());
            try {
                InfoTable infoTable = dataProcessor.exampleService(request.getDevice().getName());
                for(int i = 0; i < infoTable.getRowCount(); i++) {
                    ValueCollection row = infoTable.getRow(i);
                    ExampleOuterClass.IoTData.Builder responseStruct = ExampleOuterClass.IoTData.newBuilder();
                    responseStruct.setAnyBoolVal((Boolean) row.getValue("anyBoolVal"));
                    responseStruct.setAnyInt32Val((Integer) row.getValue("anyInt32Val"));
                    responseStruct.setAnyInt64Val((Long) row.getValue("anyInt64Val"));
                    responseStruct.setAnyStringVal(row.getStringValue("anyStringVal"));
                    response.addData(responseStruct);
                }
            } catch (Exception e) {
                System.out.println("An error occurred during getting IoT Data. " + e);
                responseObserver.onError(Status.INTERNAL
                        .withDescription(e.getMessage())
                        .withCause(e)
                        .asRuntimeException());
            }
            responseObserver.onNext(response.build());
            responseObserver.onCompleted();
        }

        @Override
        public void getIoTDataStream(ExampleOuterClass.GetIoTDataStreamReq request, StreamObserver<ExampleOuterClass.GetIoTDataStreamRes> responseObserver) {
            if(request.getDevice().getName().isEmpty()) {
                responseObserver.onError(Status.INTERNAL
                        .withDescription("Get IoT Data. Missing device name")
                        .asRuntimeException());
            }
            System.out.println("Get Iot Data Stream for device name: " + request.getDevice().getName());
            try {
                InfoTable infoTable = dataProcessor.exampleService(request.getDevice().getName());
                for(int i = 0; i < infoTable.getRowCount(); i++) {
                    ExampleOuterClass.GetIoTDataStreamRes.Builder response = ExampleOuterClass.GetIoTDataStreamRes.newBuilder();
                    ValueCollection row = infoTable.getRow(i);
                    ExampleOuterClass.IoTData.Builder responseStruct = ExampleOuterClass.IoTData.newBuilder();
                    responseStruct.setAnyBoolVal((Boolean) row.getValue("anyBoolVal"));
                    responseStruct.setAnyInt32Val((Integer) row.getValue("anyInt32Val"));
                    responseStruct.setAnyInt64Val((Long) row.getValue("anyInt64Val"));
                    responseStruct.setAnyStringVal(row.getStringValue("anyStringVal"));
                    response.setData(responseStruct);
                    responseObserver.onNext(response.build());
                }
            } catch (Exception e) {
                System.out.println("An error occurred during getting IoT Data. " + e);
                responseObserver.onError(Status.INTERNAL
                        .withDescription(e.getMessage())
                        .withCause(e)
                        .asRuntimeException());
            }
            responseObserver.onCompleted();
        }

        public StreamObserver<ExampleOuterClass.GetIoTDataStreamReq> getIoTDataBiDiStream(StreamObserver<ExampleOuterClass.GetIoTDataStreamRes> responseObserver) {
            return new StreamObserver<ExampleOuterClass.GetIoTDataStreamReq>() {
                @Override
                public void onNext(ExampleOuterClass.GetIoTDataStreamReq value) {
                    System.out.println("Got request for device: " + value.getDevice().getName());
                    ExampleOuterClass.GetIoTDataStreamRes.Builder response = ExampleOuterClass.GetIoTDataStreamRes.newBuilder();
                    ExampleOuterClass.IoTData.Builder responseStruct = ExampleOuterClass.IoTData.newBuilder();
                    responseStruct.setAnyStringVal(value.getDevice().getName() + " uuid: " + UUID.randomUUID().toString().replace("-", ""));
                    responseStruct.setAnyBoolVal(Math.random() < 0.5);
                    responseStruct.setAnyInt32Val(ThreadLocalRandom.current().nextInt(1, 5000 + 1));
                    responseStruct.setAnyInt64Val(ThreadLocalRandom.current().nextLong(100));
                    response.setData(responseStruct);
                    responseObserver.onNext(response.build());
                }

                @Override
                public void onError(Throwable t) {
                    System.out.println("on error " + t);
                    t.printStackTrace();
                }

                @Override
                public void onCompleted() {
                    System.out.println("on completed");
                    responseObserver.onCompleted();
                }
            };
        }
    }
}
