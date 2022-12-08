package main

import (
    "fmt"
    "os"
)

func main() {
    if len(os.Args) > 1 && os.Args[1] == "-v" {
        fmt.Println("nosql 0.1.0")
        return
    }
    defer func() {
        if r := recover(); r != nil {
            fmt.Println(r)
        }
    }()
    if len(os.Args) < 2 {
        usage()
    }
    switch os.Args[1] {
        case "nosql":
            nosql_main(os.Args[2:])
            break
        default:
            usage()
            break
    }
}

func usage() {
    panic("mycloud <nosql> ...")
}
