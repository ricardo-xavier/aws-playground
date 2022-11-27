package software.amazon.awssdk.enhanced.dynamodb.model;

import java.util.ArrayList;
import java.util.List;

public class PageIterable<T> {
    public List<T> items() {
        return new ArrayList<>();
    }
}
