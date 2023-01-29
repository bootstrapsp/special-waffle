package org.example.data;

import com.thingworx.types.InfoTable;
import com.thingworx.types.collections.ValueCollection;
import com.thingworx.types.primitives.StringPrimitive;

public class DataProcessor {
    private final ThingworxDataInterface thingworxDataInterface;

    public DataProcessor(ThingworxDataInterface thingworxDataInterface) {
        this.thingworxDataInterface = thingworxDataInterface;
    }

    public InfoTable exampleService (String name) throws Exception {
        ValueCollection query = new ValueCollection();
        query.put("name", new StringPrimitive(name));
        try {
            return thingworxDataInterface.processServiceRequest("SpecialWaffle_Helper", "ExampleService", query);
        } catch (Exception e) {
            System.out.println("An exception occurred during getting example service data " + e);
            throw e;
        }
    }
}
