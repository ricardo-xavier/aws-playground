package model

type AttributeType string

const (
    S AttributeType = "S"
    N               = "N"
)

var (
    typeMap = map[string]AttributeType {
        "S": S,
        "N": N,
    }
)

type Attribute struct {
    Name string
    Type AttributeType
}

func StringToAttributeType(s string) AttributeType {
    t, _ := typeMap[s]
    return t
}
