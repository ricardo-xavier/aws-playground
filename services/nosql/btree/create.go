package btree

import (
    "os"
    "encoding/binary"
)

func Create(tableName string) {
    f, err := os.Create(tableName + ".idx")
    if err != nil {
        panic(err)
    }
    buf := make([]byte, HDR_SIZE + PAGE_SIZE)
    binary.LittleEndian.PutUint64(buf[0:], 1) // pages
    binary.LittleEndian.PutUint64(buf[8:], 0) // root
    f.Write(buf)
    f.Close()
}
