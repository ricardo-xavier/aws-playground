package nosql.model;

import software.amazon.awssdk.services.dynamodb.model.AttributeType;
import software.amazon.awssdk.services.dynamodb.model.AttributeValue;

import java.util.List;
import java.util.Map;

public class Filter {
    private String attr;
    private String op; //TODO enum
    private String type; //TODO enum
    private String value;

    public static List<Filter> parse(String expression, Map<String, AttributeValue> values) throws Exception {
        String[] tokens = expression.split("\\s+");
        if (tokens.length != 3 || !"=".equals(tokens[1]) || !tokens[2].startsWith(":")) {
            throw new Exception("invalid expression: " + expression);
        }
        Filter filter = new Filter();
        filter.setAttr(tokens[0]);
        filter.setOp("=");
        AttributeValue av = values.get(tokens[2]);
        if (av.getAttributeType() != AttributeType.S) {
            throw new Exception("invalid type: " + av);
        }
        filter.setType("S");
        filter.setValue((String) av.getValue());
        return List.of(filter);
    }

    public String getAttr() {
        return attr;
    }

    public void setAttr(String attr) {
        this.attr = attr;
    }

    public String getOp() {
        return op;
    }

    public void setOp(String op) {
        this.op = op;
    }

    public String getType() {
        return type;
    }

    public void setType(String type) {
        this.type = type;
    }

    public String getValue() {
        return value;
    }

    public void setValue(String value) {
        this.value = value;
    }
}
