package main

import (
	"fmt"

	"imooc.com/TobiasKruzhor/sectionFour/tree"
)

type myTreeNode struct {
	*tree.Node // Embedding
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.Node == nil {
		return
	}

	left := myTreeNode{myNode.Left}
	left.postOrder()
	right := myTreeNode{myNode.Right}
	right.postOrder()
	myNode.Print()
}

func (myNode *myTreeNode) Traverse() {
	fmt.Println("this method is shadowed")
}

func main() {
	root := myTreeNode{&tree.Node{Value: 3}}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{Value: 5}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)

	fmt.Println("In-order traversal: ")
	fmt.Print("root.traversal(): ")
	root.Traverse()
	fmt.Print("root.Node.traversal(): ")
	root.Node.Traverse()

	fmt.Println()
	fmt.Print("My own post-order traversal: ")
	root.postOrder()
	fmt.Println()

	// var baseRoot *tree.Node
	// baseRoot := &root	//报错

}
