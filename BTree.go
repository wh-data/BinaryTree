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
	if root == nil {
		return &BinaryTree{}
	}
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

//Construct a tree from an array of exmaple data
func ConstructBinaryTreeWithExampleData(data []*BinaryTreeNodeExampleData) (tree *BinaryTreeNode) {
	root := &BinaryTreeNode{}
	leftset := false
	queue := list.New()
	tempmap := make(map[int]*BinaryTreeNode)
	for _, d := range data {
		e := queue.PushBack(d)
		if d.GetPid() == 0 {
			if !leftset {
				left := NewBinaryTreeNode(d)
				root.SetLchild(left)
				queue.Remove(e)
				tempmap[d.GetId()] = left
				leftset = true
			}
			right := NewBinaryTreeNode(d)
			root.SetRchild(right)
			queue.Remove(e)
			tempmap[d.GetId()] = right
		}

	}

	count := 0
	for {
		if queue.Len() == 0 || count == queue.Len() {
			break
		}
		e := queue.Front()
		queue.Remove(e)
		d, ok := e.Value.(*BinaryTreeNodeExampleData)
		if !ok {
			fmt.Println("data type wrong ")
			queue.PushBack(e)
			count++
			continue
		}
		if parent, exist := tempmap[d.GetPid()]; exist {
			parentData, ok := parent.GetData().(*BinaryTreeNodeExampleData)
			if !ok {
				fmt.Println("parent data type wrong ")
				queue.PushBack(e)
				count++
				continue
			}
			node := NewBinaryTreeNode(d)
			if parentData.GetLchildId() == d.GetId() {
				parent.SetLchild(node)
			} else {
				parent.SetRchild(node)
			}
			tempmap[d.GetId()] = node
			count = 0
		} else {
			queue.PushBack(e)
			count++
		}
	}
	return root
}
