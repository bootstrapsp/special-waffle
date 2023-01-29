package org.example.communication;

import com.thingworx.communications.client.ConnectedThingClient;
import com.thingworx.communications.client.things.VirtualThing;
import com.thingworx.metadata.annotations.ThingworxServiceDefinition;
import com.thingworx.metadata.annotations.ThingworxServiceParameter;
import com.thingworx.metadata.annotations.ThingworxServiceResult;
import example.protobuf.ExampleOuterClass;
import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;

import java.util.concurrent.TimeUnit;

public class RemoteServices extends VirtualThing {
    private static GetIoTDataBiDiStreamingClient getIoTDataBiDiStreamingClient;

    public RemoteServices(String name, String description, ConnectedThingClient client, String streamTargetAddress) throws Exception {
        super(name, description, client);
        this.initializeFromAnnotations();
        ManagedChannel channel = ManagedChannelBuilder.forTarget(streamTargetAddress)
                .usePlaintext()
                .keepAliveTime(300, TimeUnit.SECONDS)
                .keepAliveTimeout(43200, TimeUnit.SECONDS)
                .keepAliveWithoutCalls(true)
                .build();
        getIoTDataBiDiStreamingClient = new GetIoTDataBiDiStreamingClient(channel);
    }

    @ThingworxServiceDefinition(name = "PushData", description = "Push Data")
    @ThingworxServiceResult(name = "result", description = "", baseType = "NOTHING")
    public void PushData(
            @ThingworxServiceParameter(name = "name", description = "Device name", baseType = "STRING") String name
    ) throws Exception {
        new Thread(() -> {
            try {
                ExampleOuterClass.GetIoTDataStreamReq.Builder request = ExampleOuterClass.GetIoTDataStreamReq.newBuilder();
                ExampleOuterClass.IoTThing.Builder requestStruct = ExampleOuterClass.IoTThing.newBuilder();
                if(name != null && !name.isEmpty()){ requestStruct.setName(name); }
                request.setDevice(requestStruct);
                System.out.println("Push message for device: " + name);
                getIoTDataBiDiStreamingClient.pushData(request.build());

            } catch (Exception e) {
                System.out.println("getIoTDataBiDi stream failed: " + e.getMessage());
            }
        }).start();
    }
}
