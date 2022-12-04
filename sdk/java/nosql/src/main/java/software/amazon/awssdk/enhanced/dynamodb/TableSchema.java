package software.amazon.awssdk.enhanced.dynamodb;

import software.amazon.awssdk.enhanced.dynamodb.mapper.annotations.DynamoDbAttribute;

import java.lang.annotation.Annotation;
import java.lang.reflect.InvocationTargetException;
import java.lang.reflect.Method;
import java.util.HashMap;
import java.util.Map;

public class TableSchema<T> {
    private Map<String, String> attributeNameMap;
    private Map<String, String> attributeTypeMap;
    private T instance;

    public static <T> TableSchema<T> fromBean(Class<T> beanClass) {
        TableSchema<T> schema = new TableSchema<>();
        try {
            T instance = (T) beanClass.getConstructors()[0].newInstance();
            schema.setInstance(instance);
        } catch (InstantiationException | InvocationTargetException | IllegalAccessException e) {
            e.printStackTrace();
        }
        schema.setAttributeNameMap(new HashMap<>());
        schema.setAttributeTypeMap(new HashMap<>());
        for (Method declaredMethod : beanClass.getDeclaredMethods()) {
            if (declaredMethod.getName().startsWith("get")) {
                String name = Character.toLowerCase(declaredMethod.getName().charAt(3)) + declaredMethod.getName().substring(4);
                schema.getAttributeTypeMap().put(name, declaredMethod.getReturnType().getSimpleName());
                for (Annotation declaredAnnotation : declaredMethod.getDeclaredAnnotations()) {
                    if (declaredAnnotation.annotationType() == DynamoDbAttribute.class) {
                        DynamoDbAttribute annotation = (DynamoDbAttribute) declaredAnnotation;
                        schema.getAttributeNameMap().put(annotation.value(), name);
                    }
                }
            }
        }
        return  schema;
    }

    public Map<String, String> getAttributeNameMap() {
        return attributeNameMap;
    }

    public void setAttributeNameMap(Map<String, String> attributeNameMap) {
        this.attributeNameMap = attributeNameMap;
    }

    public Map<String, String> getAttributeTypeMap() {
        return attributeTypeMap;
    }

    public void setAttributeTypeMap(Map<String, String> attributeTypeMap) {
        this.attributeTypeMap = attributeTypeMap;
    }

    public T getInstance() {
        return this.instance;
    }

    public void setInstance(T instance) {
        this.instance = instance;
    }
}