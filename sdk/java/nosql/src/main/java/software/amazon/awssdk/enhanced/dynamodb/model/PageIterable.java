package software.amazon.awssdk.enhanced.dynamodb.model;

import java.util.ArrayList;
import java.util.List;

public class PageIterable<T> {
    private List<T> items;

    public PageIterable(List<T> items) {
        this.items = items;
    }

    public List<T> items() {
        return items;
    }
}
