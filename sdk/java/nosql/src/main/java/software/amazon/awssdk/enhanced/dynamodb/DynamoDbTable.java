package software.amazon.awssdk.enhanced.dynamodb;

import nosql.operations.PutItem;
import nosql.operations.Scan;
import software.amazon.awssdk.enhanced.dynamodb.model.PageIterable;
import software.amazon.awssdk.enhanced.dynamodb.model.ScanEnhancedRequest;

public class DynamoDbTable<T> {
    private final String name;
    private final TableSchema<T> schema;

    public DynamoDbTable(String name, TableSchema<T> schema) {
        this.name = name;
        this.schema = schema;
    }

    public PageIterable<T> scan(ScanEnhancedRequest request) {
        return new Scan().scan(name, schema, request);
    }

    public void putItem(T item) {
        new PutItem<>().put(name, schema, item);
    }
}