package btree

import (
    "fmt"
    "encoding/binary"
)

func PutItem(tree *BTree, key string, offset int64) {
    fmt.Printf("BTREE put %v %v\n", key, offset)
    page := findPage(tree, key)
    Find(tree, key, page)
    if tree.Found {
        panic("DUPLICATE KEY: " + key)
    }
    pagePos := tree.StackPagePos[len(tree.StackPagePos)-1]
    PutItemIntoPage(tree, key, offset, &page, pagePos)
}

func PutItemIntoPage(tree *BTree, key string, offset int64, page *Page, pagePos uint64) {
    itemSize := uint16(1 + len(key) + 8)
    if (page.Offset + itemSize) > BUF_SIZE {
        Split(tree, *page, pagePos)
        PutItem(tree, key, offset)
        return
    }
    if tree.Push {
        push(page.Buf, tree.ItemOffset, page.Offset, itemSize)
    }
    itemOffset := int(tree.ItemOffset)
    page.Records = page.Records + 1
    page.Offset = page.Offset + itemSize
    page.Buf[itemOffset] = byte(len(key))
    itemOffset++
    copy(page.Buf[itemOffset:], []byte(key))
    itemOffset = itemOffset + len(key)
    binary.LittleEndian.PutUint64(page.Buf[itemOffset:], uint64(offset))
    WritePage(*tree, pagePos, *page)
}

func push(buf []byte, start uint16, end uint16, size uint16) {
    tmp := make([]byte, end-start+1)
    fmt.Printf("BTREE push %v %v %v %v\n", len(buf), start, end, size)
    copy(tmp, buf[start:end])
    copy(buf[start+size:], tmp)
}
