package org.example.data;

import com.thingworx.relationships.RelationshipTypes;
import com.thingworx.types.InfoTable;
import com.thingworx.types.collections.ValueCollection;
import org.example.twx.AgentClient;

public class ThingworxDataInterface implements DataInterface{
    private final AgentClient client;

    public ThingworxDataInterface(AgentClient client) {
        this.client = client;
    }

    public InfoTable processServiceRequest(String thingName, String serviceName, ValueCollection query) throws Exception {
        return client.invokeService(RelationshipTypes.ThingworxEntityTypes.Things, thingName, serviceName, query, 35000);
    }
}
