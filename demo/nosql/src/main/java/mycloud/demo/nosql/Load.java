package mycloud.demo.nosql;

import software.amazon.awssdk.enhanced.dynamodb.DynamoDbEnhancedClient;
import software.amazon.awssdk.enhanced.dynamodb.DynamoDbTable;
import software.amazon.awssdk.enhanced.dynamodb.TableSchema;
import software.amazon.awssdk.regions.Region;
import software.amazon.awssdk.services.dynamodb.DynamoDbClient;

public class Load {
    public static void main(String[] args) {
        DynamoDbClient ddb = DynamoDbClient.builder()
                .region(Region.SA_EAST_1)
                .build();

        DynamoDbEnhancedClient enhancedClient = DynamoDbEnhancedClient.builder()
                .dynamoDbClient(ddb)
                .build();

        DynamoDbTable<User> userTable = enhancedClient.table("MYCLOUD", TableSchema.fromBean(User.class));
        User user = new User();
        user.setId("id");
        user.setName("name");
        user.setPassword("password");
        userTable.putItem(user);
    }
}
