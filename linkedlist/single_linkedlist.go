package main

import "fmt"

// single linkedklist
type Node struct {
	property int
	nextNode *Node
}
type LinkedList struct {
	headNode *Node
}

func (l *LinkedList) AddToHead(property int) {
	var node = Node{}
	node.property = property
	if node.nextNode != nil {
		node.nextNode = l.headNode
	}
	l.headNode = &node
}

func (ll LinkedList) iterateList() {
	var node *Node
	for node = ll.headNode; node != nil; node = node.nextNode {
		fmt.Println(node.property)
	}
}

func main() {
	ll := LinkedList{}
	ll.AddToHead(10)
	ll.AddToHead(20)
	ll.AddToHead(30)
	ll.iterateList()
	print(ll.headNode.property)
}

func (ll LinkedList) lastNode() *Node {
	var node *Node
	var lastNode *Node
	for node = ll.headNode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			lastNode = node
		}
	}
	return lastNode

}

func (ll LinkedList) AddToEnd(property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil
	var lastNode *Node
	lastNode = ll.lastNode()
	if lastNode != nil {
		lastNode.nextNode = node
	}
}

func (ll LinkedList) NodeWithValue(property int) *Node {
	var node *Node

	var nodeWith *Node

	for node = ll.headNode; node != nil; node = node.nextNode {
		if node.property == property {
			nodeWith = node
			break
		}

	}
	return nodeWith
}

func (ll *LinkedList) addAfter(property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil
	var nodeWith *Node
	nodeWith = ll.NodeWithValue(property)
	if node != nil {
		node.nextNode = nodeWith.nextNode
		nodeWith.nextNode = node
	}
}




