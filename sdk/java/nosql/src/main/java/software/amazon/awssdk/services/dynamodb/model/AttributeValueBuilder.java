package software.amazon.awssdk.services.dynamodb.model;

public class AttributeValueBuilder {
    private AttributeType attributeType;
    private Object value;

    public AttributeValue build() {
        return new AttributeValue(this);
    }

    public AttributeValueBuilder s(String s) {
        this.attributeType = AttributeType.S;
        this.value = s;
        return this;
    }

    public AttributeValueBuilder n(Integer n) {
        this.attributeType = AttributeType.N;
        this.value = n;
        return this;
    }

    public AttributeType getAttributeType() {
        return this.attributeType;
    }

    public Object getValue() {
        return this.value;
    }
}
