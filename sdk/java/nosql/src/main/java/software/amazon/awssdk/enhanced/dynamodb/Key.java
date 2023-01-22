package software.amazon.awssdk.enhanced.dynamodb;

import software.amazon.awssdk.KeyBuilder;
import software.amazon.awssdk.services.dynamodb.model.AttributeValue;

public class Key {
    private final AttributeValue partitionValue;
    private final AttributeValue sortValue;

    public Key(KeyBuilder builder) {
        this.partitionValue = builder.getPartitionValue();
        this.sortValue = builder.getSortValue();
    }

    public static KeyBuilder builder() {
        return new KeyBuilder();
    }
}
