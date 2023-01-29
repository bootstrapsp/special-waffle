package org.example.application;

import com.thingworx.communications.client.ClientConfigurator;
import com.thingworx.types.InfoTable;
import com.thingworx.types.collections.ValueCollection;
import org.example.communication.RemoteServices;
import org.example.data.ThingworxDataInterface;
import org.example.grpc.GrpcServer;
import org.example.twx.AgentClient;

import java.io.IOException;

public class Application {

    private static String remoteThingName = "SpecialWaffle_RemoteThing";
    private static ClientConfigurator config = new ClientConfigurator();
    private static AgentClient client;
    private static ThingworxDataInterface dataInterface;
    private static GrpcServer server;

    public static void main(String[] args) {
        System.out.println("Application Started");

        try {
            init();
            start();
        } catch (Exception e) {
            System.out.println("An exception occurred during execution." + e);
        }
    }

    private static void init() throws Exception {
        config.setUri("ws://url-with-port/Thingworx/WS");
        config.setAppKey("aaplicationKey");
        config.tunnelsEnabled(true);
        config.ignoreSSLErrors(true);
        config.setTimeoutCountReconnectionThreshold(15);
        client = new AgentClient(config);
    }

    private static void CreateRemoteThing() {
        try {
            ValueCollection query = new ValueCollection();
            InfoTable infoTable = dataInterface.processServiceRequest("SpecialWaffle_Helper", "CreateRemoteThing", query);
        } catch (Exception e) {
            System.out.println("An exception occurred during creating remote thing " + remoteThingName + " " + e);
            throw new RuntimeException();
        }
    }

    private static void start() throws Exception {
        RemoteServices thing = new RemoteServices(remoteThingName, "A basic virtual thing for special-waffle", client, "localhost:50052");
        client.bindThing(thing);

        client.start();

        dataInterface = new ThingworxDataInterface(client);
        server = new GrpcServer(dataInterface);

        new Thread(() -> {
            try {
                server.startGrpcServer();
                server.blockUntilShutdown();
            } catch (InterruptedException e){
                System.out.println("InterruptedException " + e);
            } catch (IOException e) {
                System.out.println("IOException " + e);
            }
        }).start();

        if(client.waitForConnection(10000)) {
            System.out.println("The client is now connected.");

            CreateRemoteThing();

            while (true) {
                Thread.sleep(15000);
                // Every 15 seconds we tell the thing to process a scan request. This is
                // an opportunity for the thing to query a data source, update property
                // values, and push new property values to the server.
                //
                // This loop demonstrates how to iterate over multiple VirtualThings
                // that have bound to a client. In this simple example the things
                // collection only contains one VirtualThing.
//                for (VirtualThing vt : client.getThings().values()) {
//                    vt.processScanRequest();
//                }
            }
        } else {
            System.out.println("Client did not connect within 10 seconds. Exiting");
        }
    }
}
