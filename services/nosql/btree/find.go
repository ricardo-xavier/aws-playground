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
        pos := 0
        for r := 0; r < int(page.Records); r++ {
            keyLen := int(page.Buf[pos])
            pos++
            buf := make([]byte, keyLen)
            copy(buf, page.Buf[pos:pos+keyLen])
            pos = pos + keyLen
            recordKey := string(buf)
            fmt.Printf("%d [%s]\n", r, recordKey)
            pos = pos + 8
            if recordKey >= key {
                panic("TODO recordKey >= key")
            }
        }
        tree.ItemOffset = uint16(pos)
        fmt.Printf("BTREE find %v: found = %v offset = %v\n", key, tree.Found, tree.ItemOffset)
    } else {
        panic("TODO root=NODE")
    }
}
