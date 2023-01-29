package org.example.twx;

import com.thingworx.communications.client.ClientConfigurator;
import com.thingworx.communications.client.ConnectedThingClient;

public class AgentClient extends ConnectedThingClient {

    public AgentClient(ClientConfigurator config) throws Exception {
        super(config);
    }
}
