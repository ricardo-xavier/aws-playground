package main

import (
    "bytes"
    "fmt"
    "os"
    "strings"
    "cli/model"
    "encoding/json"
    "net/http"
)

type AttributeValue struct {
    S string //TODO outros tipos
}

func nosql_main(args []string) {
    if len(args) < 1 {
        nosql_usage()
    }
    switch args[0] {
        case "scan":
            scan(args[1:])
            break
        default:
            nosql_usage()
            break
    }
}

func scan(args []string) {
    table := ""
    expression := ""
    values := ""
    i := 0
    for i < len(args) {
        arg := args[i]
        switch arg {
            case "--table-name":
                i++
                table = args[i]
                break
            case "--filter-expression":
                i++
                expression = args[i]
                break
            case "--expression-attribute-values":
                i++
                values = args[i]
                break
            default:
                nosql_usage()
                break
        }
        i++
    }
    if table == "" || (expression != "" && values == "") || (expression == "" && values != "") {
        nosql_usage()
    }

    scanRequest := model.ScanRequest {
        Table: table,
    }
    if expression != "" {
        parts := strings.Split(expression, " ")
        if len(parts) != 3 || parts[1] != "=" || parts[2][0:1] != ":" { //TODO outros operadores
            panic("invalid expression " + expression)
        }
        attr := parts[0]
        op := parts[1]
        prm := parts[2]
        attrAv := make(map[string]AttributeValue)
        json.Unmarshal([]byte(values), &attrAv)
        av := attrAv[prm]
        if av.S == "" {
            panic("prm not found " + prm)
        }
        filter := model.Filter {
            Attr: attr,
            Op: op,
            Type: "S",
            Value: av.S,
        }
        scanRequest.Filter = append(scanRequest.Filter, filter)
    }
    request, err := json.Marshal(scanRequest)
    if err != nil {
        panic(err)
    }

    url := os.Getenv("URL_NOSQL")
    if url == "" {
        panic("URL_NOSQL undefined")
    }
    response, err := http.Post(url + "scan", "application/json", bytes.NewBuffer(request))
    if err != nil {
        panic(err)
    }

    scanResponse := model.ScanResponse{}
    json.NewDecoder(response.Body).Decode(&scanResponse)
    fmt.Println(scanResponse)
    //TODO formatar a resposta
}

func nosql_usage() {
    panic("scan --table-name <name> [--filter-expression <expression> -expression-attribute-values <values>]")
}
