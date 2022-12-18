package main

import (
    "fmt"
    "os"
    "encoding/binary"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Println("show_idx <idx>")
    }

    f, err := os.Open(os.Args[1])
    if err != nil {
        panic(err)
    }

    buf := make([]byte, 16)
    f.Read(buf)
    pages := int(binary.LittleEndian.Uint64(buf[0:]))
    root := int(binary.LittleEndian.Uint64(buf[8:]))
    fmt.Printf("HEADER pages=%v root=%v\n\n", pages, root)

    for i := 0; i < pages; i++ {
        f.Seek(int64(16 + i * 4096), 0)
        buf = make([]byte, 6)
        f.Read(buf)
        tp := buf[0]
        records := binary.LittleEndian.Uint16(buf[2:])
        offset := binary.LittleEndian.Uint16(buf[4:])
        fmt.Printf("PAGE %d: tp=%v records=%v offset=%v\n", i, tp, records, offset)
        for r := 0; r < int(records); r++ {
            buf = make([]byte, 1)
            f.Read(buf)
            keyLen := int(buf[0])
            buf = make([]byte, keyLen)
            f.Read(buf)
            key := string(buf)
            buf = make([]byte, 8)
            f.Read(buf)
            data := int(binary.LittleEndian.Uint64(buf))
            fmt.Printf("\t%d %d [%s] %v\n", r, keyLen, key, data)
        }
        fmt.Println()
    }
}
