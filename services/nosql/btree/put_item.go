package btree

import (
    "fmt"
    "encoding/binary"
)

func PutItem(tree *BTree, key string, offset int64) {
    fmt.Printf("BTREE put %v %v\n", key, offset)
    Find(tree, key)
    if tree.Found {
        panic("DUPLICATE KEY: " + key)
    }
    page := tree.Stack[len(tree.Stack)-1]
    itemSize := uint16(1 + len(key) + 8)
    if (page.Offset + itemSize) > PAGE_SIZE {
        panic("page full")
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
    WritePage(*tree, tree.Root, page) //TODO pos
}

func push(buf []byte, start uint16, end uint16, size uint16) {
    tmp := make([]byte, end-start+1)
    copy(tmp, buf[start:end])
    copy(buf[start+size:], tmp)
}
