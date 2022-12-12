package model

type PutItemRequest struct {
    Table string `json:"table"`
    Items map[string]AttributeValue `json:"items"`
}

type AttributeValue struct {
    AttributeType string `json"attributeType"`
    Value string `json"value"`
}
