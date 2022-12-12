package main

import (
    "fmt"
    "os"
    "encoding/json"
    "net/http"
    "nosql/model"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Println("mycloud_nosql <port>")
    }
    if os.Args[1] == "-v" {
        fmt.Println("nosql 0.1.0")
        return
    }
    port := os.Args[1]
    fmt.Printf("nosql listening on port %s...\n", port)
    http.HandleFunc("/nosql/scan", ScanHandler)
    http.HandleFunc("/nosql/create", CreateHandler)
    http.HandleFunc("/nosql/put-item", PutItemHandler)
    http.ListenAndServe(":"+port, nil)
}

func ScanHandler(res http.ResponseWriter, req *http.Request) {
    var scanRequest model.ScanRequest
    err := json.NewDecoder(req.Body).Decode(&scanRequest)
    if err != nil {
        panic(err)
    }
    fmt.Printf("nosql scan %s...\n", scanRequest.Table)
    scanResponse := Scan(scanRequest.Table, scanRequest.Filter)
    fmt.Printf("nosql scan %s %d\n", scanRequest.Table, len(scanResponse.Items))
    response, err := json.Marshal(scanResponse)
    if err != nil {
        panic(err)
    }
    res.Write([]byte(response))
}

func CreateHandler(res http.ResponseWriter, req *http.Request) {
    var createRequest model.CreateRequest
    err := json.NewDecoder(req.Body).Decode(&createRequest)
    if err != nil {
        panic(err)
    }
    fmt.Printf("nosql create %v...\n", createRequest.TableName)
    Create(createRequest)
}

func PutItemHandler(res http.ResponseWriter, req *http.Request) {
    var putItemRequest model.PutItemRequest
    err := json.NewDecoder(req.Body).Decode(&putItemRequest)
    if err != nil {
        panic(err)
    }
    fmt.Printf("nosql put-item %v [%v]...\n", putItemRequest.Table, putItemRequest.Items)
    PutItem(putItemRequest)
}
