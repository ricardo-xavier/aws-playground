package software.amazon.awssdk.enhanced.dynamodb;

import software.amazon.awssdk.enhanced.dynamodb.model.PageIterable;
import software.amazon.awssdk.enhanced.dynamodb.model.ScanEnhancedRequest;

public class DynamoDbTable<T> {
    public PageIterable<T> scan(ScanEnhancedRequest request) {
        return new PageIterable<>();
    }
}