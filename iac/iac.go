package main

import (
    "bufio"
    "fmt"
    "strings"
    "os"
    "encoding/json"
)

type DynamodbAttribute struct {
    Name string `json:"name"`
    Type string `json:"type"`
}

type Inputs struct {
    TableName string `json:"table_name"`
    HashKey string `json:"hash_key"`
    HashKeyType string `json:"hash_key_type"`
    RangeKey string `json:"range_key"`
    RangeKeyType string `json:"range_key_type"`
    DynamodbAttributes []DynamodbAttribute `json:"dynamodb_attributes"`
}

func main() {
    if len(os.Args) != 2 || os.Args[1] != "apply" {
        fmt.Println("iac apply")
    }
    apply()
}

func apply() {
    f, err := os.Open("terragrunt.hcl")
    if err != nil {
        panic(err)
    }
    scanner := bufio.NewScanner(f)
    n := -1
    inputsJson := ""
    for scanner.Scan() {
        text := strings.TrimSpace(scanner.Text())
        if text == "" || text[0:1] == "#" {
            continue
        }
        if strings.HasPrefix(text, "inputs") && strings.HasSuffix(text, "{") {
            inputsJson = "{\n"
            n = 1
            continue
        }
        if n == -1 {
            continue
        }
        if strings.HasSuffix(text, "{") {
            n++
        }
        if strings.HasSuffix(text, "}") {
            n--
        }
        parts := strings.Split(text, "=")
        comma := ""
        if !strings.HasSuffix(text, ",") && !strings.HasSuffix(text, "{") && !strings.HasSuffix(text, "[") {
            comma = ","
        }
        if len(parts) == 2 {
            inputsJson = inputsJson + "\"" + strings.TrimSpace(parts[0]) + "\": " + parts[1] + comma + "\n"
        } else {
            inputsJson = inputsJson + text + comma + "\n"
        }
        if n == 0 {
            break
        }
    }
    f.Close()
    inputsJson = strings.ReplaceAll(inputsJson, ",\n}", "\n}")
    inputsJson = strings.ReplaceAll(inputsJson, ",\n]", "\n]")
    inputsJson = inputsJson[0:len(inputsJson)-2]

    inputs := Inputs{}
    json.Unmarshal([]byte(inputsJson), &inputs)
    fmt.Println(inputs)
}
