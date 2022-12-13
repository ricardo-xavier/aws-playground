package main

import (
    "os"
    "nosql/model"
)

func PutItem(request model.PutItemRequest) {
    schema := model.ReadSchema(request.Table + ".sch")
    dat, err := os.OpenFile(request.Table + ".dat", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        panic(err)
    }

    var values [256]string
    save := false
    for name, attr := range request.Items {
        idx := -1
        for i, field := range schema.Attributes {
            if field.Name == name {
                idx = i
                break
            }
        }
        if idx == -1 {
            newAttr := model.Attribute {
                Name: name,
                Type: attr.AttributeType,
            }
            idx = len(schema.Attributes)
            schema.Attributes = append(schema.Attributes, newAttr)
            save = true
        }
        values[idx] = attr.Value
    }
    if save {
        schema.Write(request.Table + ".sch")
    }
    for i := 0; i < len(schema.Attributes); i++ {
        if i > 0 {
            dat.WriteString("|")
        }
        dat.WriteString(values[i])
    }
    dat.WriteString("\n")
    dat.Close()
}
