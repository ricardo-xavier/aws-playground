package btree

import (
    "fmt"
    "encoding/binary"
)

func findPage(tree *BTree, key string) Page {
    page := ReadPage(*tree, tree.Root)
    tree.StackPage = nil
    tree.StackPagePos = nil
    tree.StackItemOffset = nil
    tree.StackPage = append(tree.StackPage, page)
    tree.StackPagePos = append(tree.StackPagePos, tree.Root)
    for page.Type == NODE {
        pos := 0
        next := page.Next
        found := false
        for r := 0; r < int(page.Records); r++ {
            keyLen := int(page.Buf[pos])
            buf := make([]byte, keyLen)
            copy(buf, page.Buf[pos+1:pos+1+keyLen])
            recordKey := string(buf)
            if recordKey > key {
                tree.StackItemOffset = append(tree.StackItemOffset, pos)
                next = binary.LittleEndian.Uint64(page.Buf[pos+1+keyLen:pos+1+keyLen+8])
                page = ReadPage(*tree, next)
                tree.StackPage = append(tree.StackPage, page)
                tree.StackPagePos = append(tree.StackPagePos, next)
                found = true
                break
            }
            pos = pos + 1 + keyLen + 8
        }
        if !found {
            page = ReadPage(*tree, next)
            tree.StackPage = append(tree.StackPage, page)
            tree.StackPagePos = append(tree.StackPagePos, next)
            tree.StackItemOffset = append(tree.StackItemOffset, -1)
        }
        fmt.Printf("BTREE find leaf: %v\n", page.ToString())
    }
    return page
}
