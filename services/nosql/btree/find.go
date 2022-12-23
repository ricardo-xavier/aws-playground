package btree

import (
    "fmt"
    "encoding/binary"
)


func Find(tree *BTree, key string) {
    page := findLeaf(tree, key)
    tree.Found = false
    tree.Push = false
    pos := 0
    for r := 0; r < int(page.Records); r++ {
        keyLen := int(page.Buf[pos])
        buf := make([]byte, keyLen)
        copy(buf, page.Buf[pos+1:pos+1+keyLen])
        recordKey := string(buf)
        if recordKey == key {
            tree.Found = true
            break
        }
        if recordKey > key {
            tree.Push = true
            break
        }
        pos = pos + 1 + keyLen + 8
    }
    tree.ItemOffset = uint16(pos)
    fmt.Printf("BTREE find %v: found=%v push=%v offset=%v\n", key, tree.Found, tree.Push, tree.ItemOffset)
}

func findLeaf(tree *BTree, key string) Page {
    page := ReadPage(*tree, tree.Root)
    tree.Stack = nil
    tree.StackPos = nil
    tree.Stack = append(tree.Stack, page)
    tree.StackPos = append(tree.StackPos, tree.Root)
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
                next = binary.LittleEndian.Uint64(page.Buf[pos+1+keyLen:pos+1+keyLen+8])
                page = ReadPage(*tree, next)
                tree.Stack = append(tree.Stack, page)
                tree.StackPos = append(tree.StackPos, next)
                found = true
                break
            }
            pos = pos + 1 + keyLen + 8
        }
        if !found {
            page = ReadPage(*tree, next)
            tree.Stack = append(tree.Stack, page)
            tree.StackPos = append(tree.StackPos, next)
        }
        fmt.Printf("BTREE find leaf: %v\n", page.ToString())
    }
    return page
}
