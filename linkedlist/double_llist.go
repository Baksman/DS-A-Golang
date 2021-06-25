package main

import "fmt"

type LinkedList struct {
	headNode *Node
}

type Node struct {
	property     int
	nextNode     *Node
	previousNode *Node
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

func (ll *LinkedList) nodeBetweenValues(firstPropert, secondProperty int) *Node {
	var node *Node
	var nodeWith *Node
	for node = ll.headNode; node != nil; node = node.nextNode {
		if node.previousNode != nil && node.nextNode != nil {
			if firstPropert == node.previousNode.property && secondProperty == node.nextNode.property {
				nodeWith = node
				break
			}
		}
	}
	return nodeWith
}

func (ll *LinkedList) addAfter(nodeProperty, property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil
	var nodeWith *Node
	for node = ll.headNode; node != nil; node = node.nextNode {
		if node.property == nodeProperty {
			nodeWith = node
			nodeWith.previousNode = node
			node.nextNode = nodeWith
			break

		}
	}

}

func (ll *LinkedList) AddToEnd(property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil
	var lastNode *Node
	lastNode = ll.lastNode()

	if lastNode != nil {
		lastNode.nextNode = node
		node.previousNode = lastNode
	}

}
func (l *LinkedList) AddToHead(property int) {
	var node = Node{}
	node.property = property
	if node.nextNode != nil {
		node.nextNode = l.headNode
	}
	l.headNode = &node
}

func main() {
	var linkedList LinkedList
	linkedList = LinkedList{}
	linkedList.AddToHead(1)
	linkedList.AddToHead(3)
	linkedList.AddToEnd(5)
	linkedList.addAfter(1, 7)
	fmt.Println(linkedList.headNode.property)
	var node *Node
	node = linkedList.nodeBetweenValues(1, 5)
	fmt.Println(node.property)
}
