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
