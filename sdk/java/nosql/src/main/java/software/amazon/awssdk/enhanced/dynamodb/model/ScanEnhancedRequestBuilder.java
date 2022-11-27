package software.amazon.awssdk.enhanced.dynamodb.model;

import software.amazon.awssdk.enhanced.dynamodb.Expression;

public class ScanEnhancedRequestBuilder {
    private Expression filterExpression;

    public ScanEnhancedRequest build() {
        return new ScanEnhancedRequest(this);
    }

    public ScanEnhancedRequestBuilder filterExpression(Expression filterExpression) {
        this.filterExpression = filterExpression;
        return this;
    }

    public Expression getFilterExpression() {
        return filterExpression;
    }
}
