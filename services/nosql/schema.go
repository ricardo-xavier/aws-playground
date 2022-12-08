package main

import (
    "bufio"
    "os"
    "strings"
    "nosql/model"
)

func LoadSchema(tableName string) []model.Attribute {
    var attrs []model.Attribute
    f, err := os.Open(tableName + ".sch")
    if err != nil {
        panic(err)
    }
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        item := scanner.Text()
        values := strings.Split(item, "|")
        attr := model.Attribute {
            Name: values[0],
            Type: model.StringToAttributeType(values[1]),
        }
        attrs = append(attrs, attr)
    }
    f.Close()
    return attrs
}
