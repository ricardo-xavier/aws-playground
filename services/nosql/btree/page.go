package btree

import (
    "fmt"
    "encoding/binary"
)

func ReadPage(tree BTree, pos uint64) Page {
    tree.F.Seek(int64(HDR_SIZE) + int64(pos) * int64(PAGE_SIZE), 0)
    hdr := make([]byte, PAGE_HDR_SIZE)
    tree.F.Read(hdr)
    page := Page{}
    page.Type = PageType(hdr[0])
    page.Records = binary.LittleEndian.Uint16(hdr[2:])
    page.Offset = binary.LittleEndian.Uint16(hdr[4:])
    page.Buf = make([]byte, BUF_SIZE)
    tree.F.Read(page.Buf)
    fmt.Printf("BTREE read page %v: %v\n", pos, page.ToString())
    return page
}

func WritePage(tree BTree, pos uint64, page Page) {
    tree.F.Seek(int64(HDR_SIZE) + int64(pos) * int64(PAGE_SIZE), 0)
    hdr := make([]byte, PAGE_HDR_SIZE)
    hdr[0] = byte(page.Type)
    binary.LittleEndian.PutUint16(hdr[2:], page.Records)
    binary.LittleEndian.PutUint16(hdr[4:], page.Offset)
    tree.F.Write(hdr)
    tree.F.Write(page.Buf)
}
