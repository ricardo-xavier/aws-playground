package main

import (
    "encoding/json"
    "net/http"
    "nosql/model"
)

func main() {
    port := "9000" //TODO configurar porta
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
