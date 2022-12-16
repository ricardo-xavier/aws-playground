package main

import (
    "os"
    "nosql/model"
)



func Create(request model.CreateRequest) {
    dat(request)
    sch(request)
    BTreeCreate(request.TableName)
}

func dat(request model.CreateRequest) {
    f, err := os.Create(request.TableName + ".dat")
    if err != nil {
        panic(err)
    }
    f.Close()
}

func sch(request model.CreateRequest) {
    schema := model.Schema {
        Attributes: request.DynamodbAttributes,
    }
    if request.HashKey != "" {
        index := model.Index {
            Name: request.TableName,
            Hash: request.HashKey,
            Range: request.RangeKey,
        }
        schema.Indexes = append(schema.Indexes, index)
    }
    schema.Write(request.TableName + ".sch")
}
