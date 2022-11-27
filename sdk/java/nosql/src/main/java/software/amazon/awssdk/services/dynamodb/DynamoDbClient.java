package software.amazon.awssdk.services.dynamodb;

public class DynamoDbClient {
    public DynamoDbClient(DynamoDbClientBuilder builder) {
    }

    public static DynamoDbClientBuilder builder() {
        return new DynamoDbClientBuilder();
    }
}
