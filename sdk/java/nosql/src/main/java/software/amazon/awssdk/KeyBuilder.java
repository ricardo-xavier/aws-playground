package software.amazon.awssdk;

import software.amazon.awssdk.enhanced.dynamodb.Key;
import software.amazon.awssdk.services.dynamodb.model.AttributeValue;

public class KeyBuilder {
    private AttributeValue partitionValue;
    private AttributeValue sortValue;

    public Key build() {
        return new Key(this);
    }

    public KeyBuilder partitionValue(String value) {
        this.partitionValue = AttributeValue.builder().s(value).build();
        return this;
    }

    public KeyBuilder sortValue(String value) {
        this.sortValue = AttributeValue.builder().s(value).build();
        return this;
    }

    public AttributeValue getPartitionValue() {
        return partitionValue;
    }

    public AttributeValue getSortValue() {
        return sortValue;
    }
}
