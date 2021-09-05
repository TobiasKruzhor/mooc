package tree

import (
	"fmt"
)

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

func (node TreeNode) Print() {
	fmt.Print(node.Value, " ")
}

func (node *TreeNode) SetValue(Value int) {
	if node == nil {
		fmt.Println("Setting Value to nil node. Ignored.")
		return
	}
	node.Value = Value
}


func CreateNode(Value int) *TreeNode {
	return &TreeNode{Value: Value}
}

func main() {
	var root TreeNode
	root = TreeNode{Value: 3}
	root.Left = &TreeNode{}
	root.Right = &TreeNode{5, nil, nil}
	root.Right.Left = new(TreeNode)
	root.Left.Right = CreateNode(2)

	root.Right.Left.SetValue(4)
	/**
	root.Right.Left.Print()
	fmt.Println()

	root.Print()
	root.SetValue(100)

	var pRoot *TreeNode
	// pRoot := &root
	pRoot.SetValue(200)
	pRoot.Print()
	pRoot.SetValue(300)
	pRoot.Print()
	**/
	root.Traverse()
}
