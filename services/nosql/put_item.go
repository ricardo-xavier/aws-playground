package main

import (
"fmt"
    "nosql/model"
)

func PutItem(request model.PutItemRequest) {
    for k, v := range request.Items {
        fmt.Printf("%v = [%v]\n", k, v)
    }
}
