package software.amazon.awssdk.enhanced.dynamodb.model;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import software.amazon.awssdk.enhanced.dynamodb.TableSchema;

import java.util.ArrayList;
import java.util.List;

public class PageIterable<T> {
    private List<T> items;

    public PageIterable(TableSchema<T> schema, List<String> fields, List<String> records) throws JsonProcessingException {
        ObjectMapper mapper = new ObjectMapper();
        items = new ArrayList<>();
        for (String record : records) {
            StringBuilder jsonObject = new StringBuilder("{\n");
            String[] values = record.split("\\|");
            boolean first = true;
            for (int i = 0; i < fields.size(); i++) {
                if (i < values.length) {
                    String alias = schema.getAttributeNameMap().get(fields.get(i));
                    if (alias == null) {
                        alias = fields.get(i);
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
    }

    private String quoted(String s) {
        return "\"" + s + "\"";
    }

    public List<T> items() {
        return items;
    }
}
