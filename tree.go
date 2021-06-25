package main

import "fmt"

type Tree struct {
	LeftNode  *Tree
	Value     int
	RightNode *Tree
}

func (tree *Tree) insert(m int) {
	if tree != nil {
		if tree.LeftNode == nil {
			tree.LeftNode = &Tree{nil, m, nil}
		} else {
			if tree.RightNode == nil {
				tree.RightNode = &Tree{nil, m, nil}
			} else {
				if tree.LeftNode != nil {
					tree.LeftNode.insert(m)
				} else {
					tree.RightNode.insert(m)
				}
			}

		}

	} else {
		tree = &Tree{nil, m, nil}
	}
}

func printTree(tree *Tree) {
	if tree != nil {
		fmt.Println("Value", (*tree).Value)
		fmt.Printf("Tree left node")
		printTree(tree.LeftNode)
		fmt.Printf("Tree right node")
		printTree(tree.RightNode)

	} else {
		fmt.Println("NIl")
	}
}
func main() {
var tree *Tree = &Tree{nil,1,nil}
tree.insert(3)
tree.insert(3)
tree.insert(3)
tree.insert(3)
tree.LeftNode.insert(2)
print(tree)
}
