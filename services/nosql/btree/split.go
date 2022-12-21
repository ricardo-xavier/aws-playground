package btree

import (
    "fmt"
    "encoding/binary"
)

func Split(tree *BTree, page Page, pagePos uint64) {
    if len(tree.Stack) > 1 {
        panic("TODO split depth > 1")
    }
    fmt.Printf("BTREE split %v\n", page.ToString())
    half := page.Records / 2
    pos := 0
    for r := 0; r < int(half); r++ {
        keyLen := int(page.Buf[pos])
        pos = pos + 1 + keyLen + 8
    }
    fmt.Printf("BTREE split %v %v\n", half, pos)
    newPage := Page{}
    newPage.Records = page.Records - half
    newPage.Offset = page.Offset - uint16(pos)
    page.Records = half
    page.Offset = uint16(pos)

    parent := Page{
        Type: NODE,
        Records: 1,
        Next: tree.Pages,
    }
    parent.Buf = make([]byte, BUF_SIZE)
    keyLen := int(page.Buf[pos])
    parent.Buf[0] = page.Buf[pos]
    copy(parent.Buf[1:], page.Buf[pos+1:pos+1+keyLen])
    binary.LittleEndian.PutUint64(parent.Buf[1+keyLen:], pagePos)

    fmt.Printf("BTREE split %v\n", page.ToString())
    fmt.Printf("BTREE split %v\n", newPage.ToString())
    fmt.Printf("BTREE split %v\n", parent.ToString())

    newPage.Buf = make([]byte, BUF_SIZE)
    copy(newPage.Buf[0:], page.Buf[pos:])
    for i := pos; i < int(BUF_SIZE); i++ {
        page.Buf[i] = 0
    }

    //TODO atualizar a raiz
    WritePage(*tree, pagePos, page)
    WritePage(*tree, tree.Pages, newPage)
    tree.Pages++
    WritePage(*tree, tree.Pages, parent)
    tree.Pages++

    tree.F.Seek(0, 0)
    hdr := make([]byte, HDR_SIZE)
    hdr[0] = byte(page.Type)
    binary.LittleEndian.PutUint64(hdr[0:], tree.Pages)
    binary.LittleEndian.PutUint64(hdr[8:], tree.Root)
    tree.F.Write(hdr)
}
