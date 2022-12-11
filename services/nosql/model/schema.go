package model

import (
    "os"
    "encoding/json"
    "io/ioutil"
)

type Schema struct {
    Attributes []Attribute `json"attributes"`
    Indexes []Index `json"indexes"`
}

type Attribute struct {
    Name string `json"name"`
    Type string `json"type"`
}

type Index struct {
    Name string `json"name"`
    Hash string `json"hash"`
    Range string `json"range"`
}

func (s Schema) Write(fileName string) {
    f, err := os.Create(fileName)
    if err != nil {
        panic(err)
    }
    j, err := json.MarshalIndent(s, "", "  ")
    if err != nil {
        panic(err)
    }
    f.Write(j)
    f.Close()
}

func ReadSchema(fileName string) Schema {
    j, err := ioutil.ReadFile(fileName)
    if err != nil {
        panic(err)
    }
    s := Schema{}
    json.Unmarshal([]byte(j), &s)
    return s
}

func (s Schema) GetFieldNames() []string {
    var names []string
    for _, attr := range s.Attributes {
        names = append(names, attr.Name)
    }
    return names
}
