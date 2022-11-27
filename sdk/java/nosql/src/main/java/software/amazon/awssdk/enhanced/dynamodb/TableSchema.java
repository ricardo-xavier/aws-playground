package software.amazon.awssdk.enhanced.dynamodb;

public class TableSchema<T> {
    public static <T> TableSchema<T> fromBean(Class<T> beanClass) {
        return new TableSchema<>();
    }
}