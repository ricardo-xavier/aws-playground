package main

import (
    "os"
    "encoding/binary"
    "nosql/model"
)



func Create(request model.CreateRequest) {
    dat(request)
    sch(request)
    idx(request)
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

func idx(request model.CreateRequest) {
    f, err := os.Create(request.TableName + ".idx")
    if err != nil {
        panic(err)
    }
    buf := make([]byte, 8 + 4096)
    b := make([]byte, 4)
    binary.LittleEndian.PutUint32(b, 1)
    copy(buf[0:], b)
    binary.LittleEndian.PutUint32(b, 0)
    copy(buf[4:], b)
    f.Write(buf)
    f.Close()
}
