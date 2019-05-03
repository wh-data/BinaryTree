//binary tree node data
package BinaryTree

//binary tree node data structure
type BinaryTreeNodeExampleData struct {
	Id       int
	Pid      int
	LchildId int
	RchildId int
	ExtInfo  []byte
}

func NewBinaryTreeNodeData(id, pid, lid, rid int, extinfo []byte) *BinaryTreeNodeExampleData {
	return &BinaryTreeNodeExampleData{
		Id:       id,
		Pid:      pid,
		LchildId: lid,
		RchildId: rid,
		ExtInfo:  extinfo,
	}
}

func (data *BinaryTreeNodeExampleData) GetId() int {
	if data != nil {
		return data.Id
	}
	return 0
}
func (data *BinaryTreeNodeExampleData) GetPid() int {
	if data != nil {
		return data.Pid
	}
	return 0
}

func (data *BinaryTreeNodeExampleData) GetLchildId() int {
	if data != nil {
		return data.LchildId
	}
	return 0
}

func (data *BinaryTreeNodeExampleData) GetRchildId() int {
	if data != nil {
		return data.RchildId
	}
	return 0
}
