package main

import (
	"fmt"
	"google_go/tree"
)

type myTreeNode struct {
	node *tree.Noe
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := myTreeNode{myNode.node.Left}
	left.postOrder()
	right := myTreeNode{myNode.node.Right}
	right.postOrder()
	myNode.node.Print()
}

func main() {
	var root tree.Noe
	root = tree.Noe{Value: 3}
	root.Left = &tree.Noe{}
	root.Right = &tree.Noe{5, nil,nil}
	root.Right.Left = new(tree.Noe)
	root.Left.Right = tree.CreateNode(2)

	root.Traverse()

	fmt.Println("\n")
	myroot := myTreeNode{&root}
	myroot.postOrder()
	fmt.Println()
}
