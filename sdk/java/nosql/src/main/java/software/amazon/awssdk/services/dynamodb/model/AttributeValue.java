package software.amazon.awssdk.services.dynamodb.model;

public class AttributeValue {
    private final AttributeType attributeType;
    private final Object value;

    public AttributeValue(AttributeValueBuilder builder) {
        this.attributeType = builder.getAttributeType();
        this.value = builder.getValue();
    }

    public static AttributeValueBuilder builder() {
        return new AttributeValueBuilder();
    }

    public AttributeType getAttributeType() {
        return this.attributeType;
    }

    public Object getValue() {
        return this.value;
    }
}
