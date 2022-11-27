package software.amazon.awssdk.enhanced.dynamodb;

import software.amazon.awssdk.services.dynamodb.model.AttributeValue;

import java.util.Map;

public class Expression {
    private final String expression;
    private final Map<String, AttributeValue> values;

    public Expression(ExpressionBuilder builder) {
        this.expression = builder.getExpression();
        this.values = builder.getValues();
    }

    public static ExpressionBuilder builder() {
        return new ExpressionBuilder();
    }

    public String getExpression() {
        return expression;
    }

    public Map<String, AttributeValue> getValues() {
        return values;
    }
}
