package software.amazon.awssdk.enhanced.dynamodb.model;

import software.amazon.awssdk.enhanced.dynamodb.Expression;

public class ScanEnhancedRequest {
    private Expression filterExpression;

    public ScanEnhancedRequest(ScanEnhancedRequestBuilder builder) {
        this.filterExpression = builder.getFilterExpression();
    }

    public static ScanEnhancedRequestBuilder builder() {
        return new ScanEnhancedRequestBuilder();
    }

    public Expression getFilterExpression() {
        return filterExpression;
    }
}
