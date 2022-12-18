package btree

import (
    "fmt"
)


func Find(tree *BTree, key string) {
    page  := ReadPage(*tree, tree.Root)
    tree.Stack = nil
    tree.Stack = append(tree.Stack, page)
    if page.Type == LEAF {
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
    } else {
        panic("TODO root=NODE")
    }
}
