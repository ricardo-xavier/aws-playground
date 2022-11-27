package software.amazon.awssdk.enhanced.dynamodb;

import software.amazon.awssdk.services.dynamodb.DynamoDbClient;

public class DynamoDbEnhancedClientBuilder {
    public DynamoDbEnhancedClient build() {
        return new DynamoDbEnhancedClient(this);
    }

    public DynamoDbEnhancedClientBuilder dynamoDbClient(DynamoDbClient client) {
        return this;
    }
}
