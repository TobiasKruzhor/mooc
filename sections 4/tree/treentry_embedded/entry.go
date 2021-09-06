package main

import (
	"fmt"

	"imooc.com/TobiasKruzhor/sectionFour/tree"
)

type myTreeNode struct {
	*tree.TreeNode // Embedding 内嵌
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.TreeNode == nil {
		return
	}
	left := myTreeNode{myNode.Left}
	right := myTreeNode{myNode.Right}
	left.postOrder()
	right.postOrder()
	myNode.Print()
}

func (myNode *myTreeNode) Traverse() {
	fmt.Println("this method is shadowed")
}

func main() {
	root := myTreeNode{&tree.TreeNode{Value: 3}}
	root.Left = &tree.TreeNode{}
	root.Right = &tree.TreeNode{Value: 5}
	root.Right.Left = new(tree.TreeNode)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)

	fmt.Print("In-oeder traversal:")
	
	root.TreeNode.Traverse()
	fmt.Println()
	// myRoot := myTreeNode{&root}
	fmt.Print("My own post-oeder traversal:")
	root.postOrder()
	fmt.Println()
}
