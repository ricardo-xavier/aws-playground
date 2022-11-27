package software.amazon.awssdk.enhanced.dynamodb;

import software.amazon.awssdk.services.dynamodb.model.AttributeValue;

import java.util.Map;

public class ExpressionBuilder {
    private String expression;
    private Map<String, AttributeValue> values;

    public Expression build() {
        return new Expression(this);
    }

    public ExpressionBuilder expression(String expression) {
        this.expression = expression;
        return this;
    }

    public ExpressionBuilder expressionValues(Map<String, AttributeValue> values) {
        this.values = values;
        return this;
    }

    public String getExpression() {
        return expression;
    }

    public Map<String, AttributeValue> getValues() {
        return values;
    }
}
