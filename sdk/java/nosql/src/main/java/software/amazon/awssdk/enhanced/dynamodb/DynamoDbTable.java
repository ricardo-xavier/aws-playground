package software.amazon.awssdk.enhanced.dynamodb;

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
        System.err.println("SCAN " + name);
        System.err.println(schema.getAttributeMap());
        System.err.println(request.getFilterExpression());
        return new PageIterable<>();
    }
}