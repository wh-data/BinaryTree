package BinaryTree

import (
	"testing"
)

//tesing cmd exmaple:
///usr/local/bin/go test -timeout 30s DataStructure/BinaryTree -run TestTree -v
func TestTree(t *testing.T) {
	root := NewBinaryTreeNode(0)
	lchild1 := NewBinaryTreeNode("y1")
	rchild1 := NewBinaryTreeNode("z1")
	lchild2 := NewBinaryTreeNode("y2")
	lchild1.SetLchild(lchild2)
	root.SetLchild(lchild1)
	root.SetRchild(rchild1)
	tree := NewBinaryTree(root)
	t.Log(tree.Root.GetHeight(), tree.Root.GetSize())
	list := tree.PreOrder()
	for data := list.Front(); data != nil; data = data.Next() {
		val, ok := data.Value.(*BinaryTreeNode)
		if ok {
			t.Log(val.Data)
		} else {
			t.Error("wrong")
		}
	}

}
