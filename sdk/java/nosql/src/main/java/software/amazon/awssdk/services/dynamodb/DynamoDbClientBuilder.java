package software.amazon.awssdk.services.dynamodb;

import software.amazon.awssdk.regions.Region;

public class DynamoDbClientBuilder {
    public DynamoDbClientBuilder region(Region region) {
        return this;
    }

    public DynamoDbClient build() {
        return new DynamoDbClient(this);
    }
}
