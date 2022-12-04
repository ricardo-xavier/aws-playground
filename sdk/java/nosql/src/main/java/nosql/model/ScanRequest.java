package nosql.model;

import nosql.model.Filter;

import java.util.List;

public class ScanRequest {
    private String table;
    private List<Filter> filter;

    public String getTable() {
        return table;
    }

    public void setTable(String table) {
        this.table = table;
    }

    public List<Filter> getFilter() {
        return filter;
    }

    public void setFilter(List<Filter> filter) {
        this.filter = filter;
    }
}
