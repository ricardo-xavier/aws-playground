package software.amazon.awssdk.enhanced.dynamodb;

public class DynamoDbEnhancedClient {
    public DynamoDbEnhancedClient(DynamoDbEnhancedClientBuilder builder) {

    }

    public static DynamoDbEnhancedClientBuilder builder() {
        return new DynamoDbEnhancedClientBuilder();
    }

    public <T> DynamoDbTable<T> table(String name, TableSchema<T> schema) {
        return new DynamoDbTable<>();
    }
}
