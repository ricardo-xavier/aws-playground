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
    http.HandleFunc("/nosql/scan", ScanHandler)
    http.ListenAndServe(":"+port, nil)
}

func ScanHandler(res http.ResponseWriter, req *http.Request) {
    var scanRequest model.ScanRequest
    err := json.NewDecoder(req.Body).Decode(&scanRequest)
    if err != nil {
        panic(err)
    }
    response := Scan(scanRequest.Table, scanRequest.Filter)
    jsonResponse, err := json.Marshal(response)
    if err != nil {
        panic(err)
    }
    res.Write([]byte(jsonResponse))
}
