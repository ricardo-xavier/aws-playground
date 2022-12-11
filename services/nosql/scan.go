package main

import (
    "bufio"
    "os"
    "strings"
    "nosql/model"
)

func Scan(tableName string, filter []model.Filter) model.ScanResponse {
    var response model.ScanResponse

    schema := model.ReadSchema(tableName + ".sch")
    response.Fields = schema.GetFieldNames()

    f, err := os.Open(tableName + ".dat")
    if err != nil {
        panic(err)
    }
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        item := scanner.Text()
        if len(filter) > 0 {
            values := strings.Split(item, "|")
            if !valid(filter, response.Fields, values) {
                continue
            }
        }
        response.Items = append(response.Items, item)
    }
    f.Close()
    return response
}

func valid(filter []model.Filter, fields []string, values []string) bool {
    for _, condition := range filter {
        for i, _ := range fields {
            if fields[i] == condition.Attr {
                switch condition.Op {
                    case "=":
                        if values[i] != condition.Value {
                            return false
                        }
                        break
                }
            }
        }
    }
    return true
}
