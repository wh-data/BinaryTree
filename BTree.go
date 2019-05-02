/*author : wenhai*/
//binary tree
package BinaryTree

import (
	"container/list"
	"fmt"
)

//binary tree structure
type BinaryTree struct {
	Root   *BinaryTreeNode
	Height int
	Size   int
}

//make a new binary tree
func NewBinaryTree(root *BinaryTreeNode) *BinaryTree {
	return &BinaryTree{
		Root:   root,
		Height: root.GetHeight(),
		Size:   root.GetSize(),
	}
}

//get tree height
//calculate the highest level of the tree
func (BTree *BinaryTree) GetTreeHeight() (int, error) {
	if BTree == nil {
		return 0, fmt.Errorf("Tree is nil")
	}
	return BTree.Height, nil
}

//get tree size
//get the amount of nodes including all children node and it self
func (BTree *BinaryTree) GetTreeSize() (int, error) {
	if BTree == nil {
		return 0, fmt.Errorf("Tree is nil")
	}
	return BTree.Size, nil
}

//get the whole tree
func (BTree *BinaryTree) GetTree(order int) *list.List {
	switch order {
	case 0:

	}
	return nil
}

//preorder前序遍历
func (BTree *BinaryTree) PreOrder() *list.List {
	travaler := list.New() //&list.List{}
	preOrder(BTree.Root, travaler)
	return travaler
}

func preOrder(node *BinaryTreeNode, l *list.List) {
	if node == nil {
		return
	}
	l.PushBack(node)
	preOrder(node.GetLchild(), l)
	preOrder(node.GetRchild(), l)
}

//in-order中序遍历

//sur-order后序遍历
