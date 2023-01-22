package awsplayground.demo.nosql;

import software.amazon.awssdk.enhanced.dynamodb.DynamoDbEnhancedClient;
import software.amazon.awssdk.enhanced.dynamodb.DynamoDbTable;
import software.amazon.awssdk.enhanced.dynamodb.Key;
import software.amazon.awssdk.enhanced.dynamodb.TableSchema;
import software.amazon.awssdk.enhanced.dynamodb.model.ScanEnhancedRequest;
import software.amazon.awssdk.regions.Region;
import software.amazon.awssdk.services.dynamodb.DynamoDbClient;

public class ScanAndGet {
    public static void main(String[] args) throws Exception {
        DynamoDbClient ddb = DynamoDbClient.builder()
                .region(Region.SA_EAST_1)
                .build();

        DynamoDbEnhancedClient enhancedClient = DynamoDbEnhancedClient.builder()
                .dynamoDbClient(ddb)
                .build();

        ScanEnhancedRequest request = ScanEnhancedRequest.builder().build();

        DynamoDbTable<PrimaryKey> pkTable = enhancedClient.table("PAULOBET", TableSchema.fromBean(PrimaryKey.class));
        for (PrimaryKey pk : pkTable.scan(request).items()) {
            //System.out.printf("%s : %s%n", pk.getId(), pk.getSort());
            switch (pk.getSort()) {
                case "USER":
                    getUser(enhancedClient, pk);
                    break;
            }
        }
    }

    private static void getUser(DynamoDbEnhancedClient enhancedClient, PrimaryKey pk) {
        Key key = Key.builder().partitionValue(pk.getId()).sortValue(pk.getSort()).build();
        DynamoDbTable<User> userTable = enhancedClient.table("PAULOBET", TableSchema.fromBean(User.class));
        User user = userTable.getItem(key);
        System.out.printf("USER %s : %s%n", user.getId(), user.getName());
    }
}
