package model

type ScanResponse struct {
    Fields []string `json:"fields"`
    Items []string `json:"items"`
}
