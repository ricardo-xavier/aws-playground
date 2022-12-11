package main

import (
    "bufio"
    "bytes"
    "fmt"
    "strings"
    "os"
    "net/http"
)

func main() {
    if len(os.Args) != 2 || os.Args[1] != "apply" {
        fmt.Println("iac apply")
    }
    apply()
}

func apply() {
    f, err := os.Open("terragrunt.hcl")
    if err != nil {
        panic(err)
    }
    scanner := bufio.NewScanner(f)
    n := -1
    inputsJson := ""
    for scanner.Scan() {
        text := strings.TrimSpace(scanner.Text())
        if text == "" || text[0:1] == "#" {
            continue
        }
        if strings.HasPrefix(text, "inputs") && strings.HasSuffix(text, "{") {
            inputsJson = "{\n"
            n = 1
            continue
        }
        if n == -1 {
            continue
        }
        if strings.HasSuffix(text, "{") {
            n++
        }
        if strings.HasSuffix(text, "}") {
            n--
        }
        parts := strings.Split(text, "=")
        comma := ""
        if !strings.HasSuffix(text, ",") && !strings.HasSuffix(text, "{") && !strings.HasSuffix(text, "[") {
            comma = ","
        }
        if len(parts) == 2 {
            inputsJson = inputsJson + "\"" + strings.TrimSpace(parts[0]) + "\": " + parts[1] + comma + "\n"
        } else {
            inputsJson = inputsJson + text + comma + "\n"
        }
        if n == 0 {
            break
        }
    }
    f.Close()


    inputsJson = strings.ReplaceAll(inputsJson, ",\n}", "\n}")
    inputsJson = strings.ReplaceAll(inputsJson, ",\n]", "\n]")
    inputsJson = inputsJson[0:len(inputsJson)-2]

    url := os.Getenv("URL_NOSQL")
    if url == "" {
        panic("URL_NOSQL undefined")
    }
    response, err := http.Post(url + "create", "application/json", bytes.NewBuffer([]byte(inputsJson)))
    if err != nil {
        panic(err)
    }
    fmt.Println(response)
}
