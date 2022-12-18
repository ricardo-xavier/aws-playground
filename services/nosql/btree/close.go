package btree

func Close(tree BTree) {
    tree.F.Close()
}
