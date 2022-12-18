package btree

import (
    "fmt"
    "os"
    "encoding/binary"
)

func Open(tableName string) BTree {
    f, err := os.OpenFile(tableName + ".idx", os.O_RDWR, 0644)
    if err != nil {
        panic(err)
    }
    buf := make([]byte, HDR_SIZE)
    f.Read(buf)
    pages := binary.LittleEndian.Uint64(buf[0:])
    root := binary.LittleEndian.Uint64(buf[8:])
    fmt.Printf("BTREE open %s: pages = %v root = %v\n", tableName, pages, root)
    tree := BTree {
        F: f,
        Pages: pages,
        Root: root,
    }
    return tree
}
