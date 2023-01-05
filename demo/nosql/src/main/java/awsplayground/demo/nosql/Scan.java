package awsplayground.demo.nosql;

import software.amazon.awssdk.enhanced.dynamodb.DynamoDbEnhancedClient;
import software.amazon.awssdk.enhanced.dynamodb.DynamoDbTable;
import software.amazon.awssdk.enhanced.dynamodb.Expression;
import software.amazon.awssdk.enhanced.dynamodb.TableSchema;
import software.amazon.awssdk.enhanced.dynamodb.model.ScanEnhancedRequest;
import software.amazon.awssdk.regions.Region;
import software.amazon.awssdk.services.dynamodb.DynamoDbClient;
import software.amazon.awssdk.services.dynamodb.model.AttributeValue;

import java.util.HashMap;
import java.util.Map;

public class Scan {
    public static void main(String[] args) throws Exception {
        DynamoDbClient ddb = DynamoDbClient.builder()
                .region(Region.SA_EAST_1)
                .build();

        DynamoDbEnhancedClient enhancedClient = DynamoDbEnhancedClient.builder()
                .dynamoDbClient(ddb)
                .build();

        AttributeValue sort = AttributeValue.builder()
                .s("USER")
                .build();

        Map<String, AttributeValue> expressionValues = new HashMap<>();
        expressionValues.put(":sort", sort);

        Expression expression = Expression.builder()
                .expression("sort = :sort")
                .expressionValues(expressionValues)
                .build();

        ScanEnhancedRequest request = ScanEnhancedRequest.builder()
                .filterExpression(expression)
                .build();

        DynamoDbTable<User> userTable = enhancedClient.table("PAULOBET", TableSchema.fromBean(User.class));
        for (User user : userTable.scan(request).items()) {
            System.out.printf("%s : %s%n", user.getId(), user.getName());
        }
    }
}
