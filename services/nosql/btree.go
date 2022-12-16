package main

import (
    "fmt"
    "os"
    "encoding/binary"
)

const (
    HDR_SIZE  int = 8
    PAGE_SIZE int = 4096
    LEAF = 0
    NODE = 1
)

type PageType byte

type BTree struct {
    f *os.File
    pages uint32
    root uint32
}

type Page struct {
    tp PageType
    records uint32
    offset uint32
}

func BTreeCreate(tableName string) {
    f, err := os.Create(tableName + ".idx")
    if err != nil {
        panic(err)
    }
    buf := make([]byte, HDR_SIZE + PAGE_SIZE)
    b := make([]byte, 4)
    binary.LittleEndian.PutUint32(b, 1) // num pages
    copy(buf[0:], b)
    binary.LittleEndian.PutUint32(b, 0) // root
    copy(buf[4:], b)
    f.Write(buf)
    f.Close()
}

func BTreeOpen(tableName string) BTree {
    f, err := os.OpenFile(tableName + ".idx", os.O_RDWR, 0644)
    if err != nil {
        panic(err)
    }
    buf := make([]byte, HDR_SIZE)
    f.Read(buf)
    pages := binary.LittleEndian.Uint32(buf[0:])
    root := binary.LittleEndian.Uint32(buf[4:])
    fmt.Printf("pages = %v root = %v\n", pages, root)
    btree := BTree {
        f: f,
        pages: pages,
        root: root,
    }
    return btree
}

func (btree BTree) ReadPage(pos uint32) Page {
    btree.f.Seek(int64(HDR_SIZE) + int64(pos) * int64(PAGE_SIZE), 0)
    buf := make([]byte, HDR_SIZE)
    btree.f.Read(buf)
    records := binary.LittleEndian.Uint32(buf[0:])
    page := Page {
        records: records,
    }
    return page
}

func (btree BTree) PutItem(key string, offset int64) {
    fmt.Printf("put [%v] %v\n", key, offset)
    page := btree.ReadPage(btree.root)
    fmt.Printf("records = %v\n", page.records)
}

func (btree BTree) Close() {
    btree.f.Close()
}
