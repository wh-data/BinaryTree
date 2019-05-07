package BinaryTree

import (
	"container/list"
	"fmt"
	"testing"

	pb "github.com/wh-data/BinaryTree/exampleData"
)

//tesing cmd exmaple:
///usr/local/bin/go test -timeout 30s DataStructure/BinaryTree -run TestTree -v

//this is for test only
func NewBinaryTreeNodeData(id, pid, lid, rid *int32, extinfo []byte) *pb.BinaryTreeNodeExampleData {
	return &pb.BinaryTreeNodeExampleData{
		Id:       id,
		Pid:      pid,
		LchildId: lid,
		RchildId: rid,
		ExtInfo:  extinfo,
	}
}

//Construct a tree from an array of exmaple data
func BuildBinaryTreeWithExampleData(data []*pb.BinaryTreeNodeExampleData) (tree *BinaryTreeNode) {
	root := &BinaryTreeNode{}
	leftset := false
	queue := list.New()
	tempmap := make(map[int32]*BinaryTreeNode)
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
		d, ok := e.Value.(*pb.BinaryTreeNodeExampleData)
		if !ok {
			fmt.Println("data type wrong ")
			queue.PushBack(e)
			count++
			continue
		}
		if parent, exist := tempmap[d.GetPid()]; exist {
			parentData, ok := parent.GetData().(*pb.BinaryTreeNodeExampleData)
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
