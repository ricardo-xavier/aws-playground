package nosql.model;

import software.amazon.awssdk.services.dynamodb.model.AttributeValue;

import java.util.Map;

public class PutItemRequest {
    private String table;
    private Map<String, AttributeValue> items;

    public String getTable() {
        return table;
    }

    public void setTable(String table) {
        this.table = table;
    }

    public Map<String, AttributeValue> getItems() {
        return items;
    }

    public void setItems(Map<String, AttributeValue> items) {
        this.items = items;
    }
}
