package model

type CreateRequest struct {
    TableName string `json:"table_name"`
    HashKey string `json:"hash_key"`
    HashKeyType string `json:"hash_key_type"`
    RangeKey string `json:"range_key"`
    RangeKeyType string `json:"range_key_type"`
    DynamodbAttributes []Attribute `json:"dynamodb_attributes"`
}
