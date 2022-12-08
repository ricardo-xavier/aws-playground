package model

type Filter struct {
    Attr string `json:"attr"`
    Op string `json:"op"`
    Type string `json:"type"`
    Value string `json:"value"`
}

type ScanRequest struct {
    Table string `json:"table"`
    Filter []Filter `json:"filter"`
}
