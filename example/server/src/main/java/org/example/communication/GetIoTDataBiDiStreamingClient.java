package org.example.communication;

import example.protobuf.ExampleGrpc;
import example.protobuf.ExampleOuterClass;
import io.grpc.ManagedChannel;
import io.grpc.Status;
import io.grpc.stub.StreamObserver;

public class GetIoTDataBiDiStreamingClient implements StreamObserver<ExampleOuterClass.GetIoTDataStreamRes>{

    private ExampleGrpc.ExampleStub asyncStub;
    private StreamObserver<ExampleOuterClass.GetIoTDataStreamReq> requestObserver;

    GetIoTDataBiDiStreamingClient(ManagedChannel channel) {
        asyncStub = ExampleGrpc.newStub(channel);
        requestObserver = asyncStub.getIoTDataBiDiStream(this);
    }

    @Override
    public void onNext(ExampleOuterClass.GetIoTDataStreamRes value) {
        System.out.println("Got response: " + value);
    }

    @Override
    public void onError(Throwable t) {
        Status status = Status.fromThrowable(t);
        System.out.println(status);
        requestObserver = asyncStub.getIoTDataBiDiStream(this);
    }

    @Override
    public void onCompleted() {
        requestObserver = asyncStub.getIoTDataBiDiStream(this);
    }

    public void pushData(ExampleOuterClass.GetIoTDataStreamReq request) {
        synchronized (requestObserver) {
            requestObserver.onNext(request);
        }
    }
}
