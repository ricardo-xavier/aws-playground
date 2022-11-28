package main

import (
    "encoding/json"
    "net/http"
    "nosql/model"
)

func main() {
    port := "9000"
    http.HandleFunc("/dynamodb/", Handler)
    http.ListenAndServe(":"+port, nil)
}

func Handler(res http.ResponseWriter, req *http.Request) {
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
