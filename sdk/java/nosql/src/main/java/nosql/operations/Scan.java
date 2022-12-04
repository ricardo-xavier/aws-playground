package nosql.operations;

import com.fasterxml.jackson.databind.ObjectMapper;
import nosql.model.Filter;
import nosql.model.ScanRequest;
import nosql.model.ScanResponse;
import software.amazon.awssdk.enhanced.dynamodb.TableSchema;
import software.amazon.awssdk.enhanced.dynamodb.model.PageIterable;
import software.amazon.awssdk.enhanced.dynamodb.model.ScanEnhancedRequest;

import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.util.ArrayList;
import java.util.List;

public class Scan<T> {
    public PageIterable<T> scan(String name, TableSchema<T> schema, ScanEnhancedRequest request) {
        try {
            ScanRequest scanRequest = new ScanRequest();
            scanRequest.setTable(name);

            if (request.getFilterExpression() != null) {
                List<Filter> filter = Filter.parse(request.getFilterExpression().getExpression(), request.getFilterExpression().getValues());
                scanRequest.setFilter(filter);
            }

            ObjectMapper mapper = new ObjectMapper();
            String json = mapper.writeValueAsString(scanRequest);

            HttpClient httpClient = HttpClient.newBuilder().build();
            HttpRequest httpRequest = HttpRequest.newBuilder()
                    .POST(HttpRequest.BodyPublishers.ofString(json))
                    .uri(URI.create("http://localhost:9000/nosql/scan")) //TODO configurar
                    .setHeader("Content-Type", "application/json")
                    .build();

            HttpResponse<String> httpResponse = httpClient.send(httpRequest, HttpResponse.BodyHandlers.ofString());

            ScanResponse response = mapper.readValue(httpResponse.body(), ScanResponse.class);

            List<T> items = new ArrayList<>();
            for (String record : response.getItems()) {
                StringBuilder jsonObject = new StringBuilder("{\n");
                String[] values = record.split("\\|");
                boolean first = true;
                for (int i = 0; i < response.getFields().size(); i++) {
                    if (i < values.length) {
                        String alias = schema.getAttributeNameMap().get(response.getFields().get(i));
                        if (alias == null) {
                            alias = response.getFields().get(i);
                        }
                        String tp = schema.getAttributeTypeMap().get(alias);
                        if (tp == null) {
                            continue;
                        }
                        if (first) {
                            first = false;
                        } else {
                            jsonObject.append(",\n");
                        }
                        jsonObject.append(quoted(alias)).append(":");
                        if (tp.equals("String")) {
                            jsonObject.append(quoted(values[i]));
                        } else {
                            jsonObject.append(values[i]);
                        }
                    }
                }
                jsonObject.append("\n}");
                T item = (T) mapper.readValue(jsonObject.toString(), schema.getInstance().getClass());
                items.add(item);
            }

            return new PageIterable<>(items);

        } catch (Exception e) {
            e.printStackTrace();
            return null;
        }
    }

    private String quoted(String s) {
        return "\"" + s + "\"";
    }

}
