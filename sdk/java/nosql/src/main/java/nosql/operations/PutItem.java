package nosql.operations;

import com.fasterxml.jackson.databind.ObjectMapper;
import nosql.model.PutItemRequest;
import software.amazon.awssdk.enhanced.dynamodb.TableSchema;
import software.amazon.awssdk.services.dynamodb.model.AttributeValue;

import java.lang.reflect.Method;
import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.util.HashMap;
import java.util.Map;

public class PutItem<T> {
    public <T> void put(String name, TableSchema<T> schema, T item) throws Exception {
        Map<String, AttributeValue> items = new HashMap<>();
        for (Method method : schema.getInstance().getClass().getDeclaredMethods()) {
            if (method.getName().startsWith("get")
                    && Character.isUpperCase(method.getName().charAt(3))
                    && method.getParameterCount() == 0) {
                String fieldName = Character.toLowerCase(method.getName().charAt(3)) + method.getName().substring(4);
                String fieldType = schema.getAttributeTypeMap().get(fieldName);
                if (fieldType != null) {
                    String alias = fieldName;
                    for (String key : schema.getAttributeNameMap().keySet()) {
                        String value = schema.getAttributeNameMap().get(key);
                        if (fieldName.equals(value)) {
                            alias = key;
                            break;
                        }
                    }
                    if (fieldType.equals("String")) {
                        String fieldValue = (String) method.invoke(item);
                        items.put(alias, AttributeValue.builder().s(fieldValue).build());
                    } else {
                        Integer fieldValue = (Integer) method.invoke(item);
                        items.put(alias, AttributeValue.builder().n(fieldValue).build());
                    }
                }
            }
        }

        PutItemRequest putItemRequest = new PutItemRequest();
        putItemRequest.setTable(name);
        putItemRequest.setItems(items);
        ObjectMapper mapper = new ObjectMapper();
        String json = mapper.writeValueAsString(putItemRequest);

        String url = System.getProperty("URL_NOSQL");
        if (url == null) {
            throw new Exception("URL_NOSQL undefined");
        }

        HttpClient httpClient = HttpClient.newBuilder().build();
        HttpRequest httpRequest = HttpRequest.newBuilder()
                .POST(HttpRequest.BodyPublishers.ofString(json))
                .uri(URI.create(url + "put-item"))
                .setHeader("Content-Type", "application/json")
                .build();
        System.out.println(json);
        HttpResponse<String> httpResponse = httpClient.send(httpRequest, HttpResponse.BodyHandlers.ofString());
        System.out.println(httpResponse.statusCode());
    }
}
