package btree

import (
    "fmt"
    "encoding/binary"
)

func Split(tree *BTree, page Page, pagePos uint64) {
    fmt.Printf("BTREE split %v\n", page.ToString())
    showPage(page)

    half := page.Records / 2
    pos := 0
    for r := 0; r < int(half); r++ {
        keyLen := int(page.Buf[pos])
        pos = pos + 1 + keyLen + 8
    }

    newPage := Page{}
    newPage.Records = page.Records - half
    newPage.Offset = page.Offset - uint16(pos)
    newPos := tree.Pages
    tree.Pages++

    page.Records = half
    page.Offset = uint16(pos)

    var parent Page
    var parentPos uint64
    if len(tree.StackPage) > 1 {
        parentItemOffset := tree.StackItemOffset[len(tree.StackItemOffset)-1]
        parent = tree.StackPage[len(tree.StackPage)-2]
        parentPos = tree.StackPagePos[len(tree.StackPagePos)-2]
        if parentItemOffset == -1 {
            parent.Next = newPos
        } else {
            parentKeyLen := int(parent.Buf[parentItemOffset])
            binary.LittleEndian.PutUint64(parent.Buf[parentItemOffset+1+parentKeyLen:], newPos)
        }
    } else {
        parent = Page{
            Type: NODE,
            Next: newPos,
        }
        parent.Buf = make([]byte, BUF_SIZE)
        parentPos = newPos + 1
        tree.Pages++
        tree.Root = parentPos
    }

    keyLen := int(page.Buf[pos])
    key := string(page.Buf[pos+1:pos+1+keyLen])
    Find(tree, key, parent)
    PutItemIntoPage(tree, key, int64(pagePos), &parent, parentPos)

    newPage.Buf = make([]byte, BUF_SIZE)
    copy(newPage.Buf[0:], page.Buf[pos:])
    for i := pos; i < int(BUF_SIZE); i++ {
        page.Buf[i] = 0
    }

    fmt.Printf("BTREE split actual %v\n", page.ToString())
    showPage(page)
    fmt.Printf("BTREE split new %v\n", newPage.ToString())
    showPage(newPage)
    fmt.Printf("BTREE split parent %v\n", parent.ToString())
    showPage(parent)

    WritePage(*tree, pagePos, page)
    WritePage(*tree, newPos, newPage)

    tree.F.Seek(0, 0)
    hdr := make([]byte, HDR_SIZE)
    hdr[0] = byte(page.Type)
    binary.LittleEndian.PutUint64(hdr[0:], tree.Pages)
    binary.LittleEndian.PutUint64(hdr[8:], tree.Root)
    tree.F.Write(hdr)
}

func showPage(page Page) {
    pos := 0
    for r := 0; r < int(page.Records); r++ {
        keyLen := int(page.Buf[pos])
        buf := make([]byte, keyLen)
        copy(buf, page.Buf[pos+1:pos+1+keyLen])
        recordKey := string(buf)
        address := binary.LittleEndian.Uint64(page.Buf[pos+1+keyLen:pos+1+keyLen+8])
        fmt.Printf("\t%v %v %v\n", r, recordKey, address)
        pos = pos + 1 + keyLen + 8
    }
}
