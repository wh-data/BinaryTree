//binary tree node

package BinaryTree

import (
	"fmt"
	"math"
)

//binary tree node structure
type BinaryTreeNode struct {
	Data   interface{}
	Parent *BinaryTreeNode
	Lchild *BinaryTreeNode
	Rchild *BinaryTreeNode
	Height int
	Size   int
}

//make a new binary tree node
func NewBinaryTreeNode(data interface{}) *BinaryTreeNode {
	return &BinaryTreeNode{
		Data:   data,
		Height: 0,
		Size:   1,
	}
}

//get node level
func (BTnode *BinaryTreeNode) GetNodeLevel() int {
	level := 0
	if !BTnode.HasParent() {
		return level
	}
	level++
	return level + BTnode.Parent.GetNodeLevel()
}

//does node has data
func (BTnode *BinaryTreeNode) HasData() bool {
	if BTnode.Data != nil {
		return true
	}
	return false
}

//dose node has parent
func (BTnode *BinaryTreeNode) HasParent() bool {
	return BTnode.Parent != nil
}

//dose node has child
func (BTnode *BinaryTreeNode) HasChild() bool {
	return BTnode.Lchild != nil || BTnode.Rchild != nil
}

//does node has left child
func (BTnode *BinaryTreeNode) HasLchild() bool {
	return BTnode.Lchild != nil
}

//does node has right child
func (BTnode *BinaryTreeNode) HasRchild() bool {
	return BTnode.Rchild != nil
}

//is left child
func (BTnode *BinaryTreeNode) IsLchild() bool {
	return BTnode.Parent != nil && BTnode.Parent.Lchild == BTnode
}

//is right child
func (BTnode *BinaryTreeNode) IsRchild() bool {
	return BTnode.Parent != nil && BTnode.Parent.Rchild == BTnode
}

//is leaf node
func (BTnode *BinaryTreeNode) IsLeaf() bool {
	return BTnode.Lchild == nil && BTnode.Rchild == nil
}

//get node data
func (BTnode *BinaryTreeNode) GetData() interface{} {
	if BTnode != nil && BTnode.Data != nil {
		return BTnode.Data
	}
	return nil
}

//get node parent
func (BTnode *BinaryTreeNode) GetParent() *BinaryTreeNode {
	if BTnode != nil && BTnode.HasParent() {
		return BTnode.Parent
	}
	return nil
}

//get node left child
func (BTnode *BinaryTreeNode) GetLchild() *BinaryTreeNode {
	if BTnode != nil && BTnode.HasLchild() {
		return BTnode.Lchild
	}
	return nil
}

//get node right child
func (BTnode *BinaryTreeNode) GetRchild() *BinaryTreeNode {
	if BTnode != nil && BTnode.HasRchild() {
		return BTnode.Rchild
	}
	return nil
}

//get node height
func (BTnode *BinaryTreeNode) GetHeight() int {
	if BTnode != nil {
		return BTnode.Height
	}
	return 0
}

//get node Size
func (BTnode *BinaryTreeNode) GetSize() int {
	if BTnode != nil {
		return BTnode.Size
	}
	return 0
}

//set node Height and regression parent height
//the leaf node's height is 0
func (BTnode *BinaryTreeNode) SetHeight() {
	height := int(math.Max(float64(BTnode.GetLchild().GetHeight()+1), float64(BTnode.GetRchild().GetHeight()+1)))
	if BTnode.Height == height {
		//no need regression
		return
	}
	BTnode.Height = height
	if BTnode.Parent == nil {
		//the end of regression
		return
	}
	BTnode.GetParent().SetHeight()
}

//set node size and regression parent size
func (BTnode *BinaryTreeNode) SetSize() {
	size := 1 + BTnode.GetLchild().GetSize() + BTnode.GetRchild().GetSize()
	if BTnode.Size == size {
		//no need regression
		return
	}
	BTnode.Size = size
	if BTnode.Parent == nil {
		//the end of regression
		return
	}
	BTnode.GetParent().SetSize()
}

//set node data
func (BTnode *BinaryTreeNode) SetData(data interface{}) error {
	if BTnode == nil {
		return fmt.Errorf("node is nil")
	}
	BTnode.Data = data
	return nil
}

//set node left child
func (BTnode *BinaryTreeNode) SetLchild(child *BinaryTreeNode) (*BinaryTreeNode, error) {
	if BTnode == nil {
		return nil, fmt.Errorf("node is nil")
	}
	if child == nil {
		return nil, fmt.Errorf("child(to be set) is nil")
	} else {
		child.CutoffFromParent()
	}
	lchild := BTnode.GetLchild()
	if lchild != nil {
		lchild.Parent = nil //dont use cutoff because that function will regression size and height
	}
	BTnode.Lchild = child
	child.Parent = BTnode
	BTnode.SetHeight()
	BTnode.SetSize()
	return lchild, nil
}

//set node right child
func (BTnode *BinaryTreeNode) SetRchild(child *BinaryTreeNode) (*BinaryTreeNode, error) {
	if BTnode == nil {
		return nil, fmt.Errorf("node is nil")
	}
	if child == nil {
		return nil, fmt.Errorf("child(to be set) is nil")
	} else {
		child.CutoffFromParent()
	}
	rchild := BTnode.GetRchild()
	if rchild != nil {
		rchild.Parent = nil //dont use cutoff because that function will regression size and height
	}
	BTnode.Rchild = child
	child.Parent = BTnode
	BTnode.SetHeight()
	BTnode.SetSize()
	return rchild, nil
}

//cutoff node from parent
func (BTnode *BinaryTreeNode) CutoffFromParent() error {
	if BTnode == nil {
		return fmt.Errorf("node is nil")
	}
	if !BTnode.HasParent() {
		return fmt.Errorf("parent is nil")
	}
	parent := BTnode.Parent //BTnode.Parent has to assign to a new pointer before it is assigned to nil
	if BTnode.IsLchild() {
		parent.Lchild = nil
	}
	if BTnode.IsRchild() {
		parent.Rchild = nil
	}
	BTnode.Parent = nil
	parent.SetHeight()
	parent.SetSize()
	return nil

}
