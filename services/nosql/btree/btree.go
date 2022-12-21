package btree

import (
    "os"
)

const (
    HDR_SIZE uint16 = 16 // 8 + 8
    BUF_SIZE uint16 = 4082 // 4096 - 14
    PAGE_HDR_SIZE uint16 = 14 // 1 + 1 + 2 + 2 + 8
    PAGE_SIZE uint16 = PAGE_HDR_SIZE + BUF_SIZE
    LEAF = 0
    NODE = 1
)

type PageType byte

type BTree struct {
    F *os.File
    Pages uint64
    Root uint64
    Stack []Page
    StackPos []uint64
    Found bool
    ItemOffset uint16
    Push bool
}

type Page struct {
    Type PageType
    Filler byte
    Records uint16
    Offset uint16
    Next uint64
    Buf []byte
}
