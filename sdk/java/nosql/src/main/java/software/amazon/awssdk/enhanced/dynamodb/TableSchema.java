package software.amazon.awssdk.enhanced.dynamodb;

import software.amazon.awssdk.enhanced.dynamodb.mapper.annotations.DynamoDbAttribute;

import java.lang.annotation.Annotation;
import java.lang.reflect.Method;
import java.util.HashMap;
import java.util.Map;

public class TableSchema<T> {
    private Map<String, String> attributeMap;

    public static <T> TableSchema<T> fromBean(Class<T> beanClass) {
        TableSchema<T> schema = new TableSchema<>();
        schema.setAttributeMap(new HashMap<>());
        for (Method declaredMethod : beanClass.getDeclaredMethods()) {
            for (Annotation declaredAnnotation : declaredMethod.getDeclaredAnnotations()) {
                if (declaredAnnotation.annotationType() == DynamoDbAttribute.class) {
                    DynamoDbAttribute annotation = (DynamoDbAttribute) declaredAnnotation;
                    schema.attributeMap.put(declaredMethod.getName(), annotation.value());
                }
            }
        }
        return  schema;
    }

    public Map<String, String> getAttributeMap() {
        return attributeMap;
    }

    public void setAttributeMap(Map<String, String> attributeMap) {
        this.attributeMap = attributeMap;
    }
}