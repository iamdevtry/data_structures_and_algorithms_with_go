package main

import "fmt"

// Node class
type Node struct {
	property int
	nextNode *Node
}

// LinkedList class
type LinkedList struct {
	headNode *Node
}

func (linkedList *LinkedList) AddToHead(property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil

	if linkedList.headNode != nil {
		node.nextNode = linkedList.headNode
	}

	linkedList.headNode = node
}

// IterateList method iterates over LinkedList
func (linkedList *LinkedList) IterateList() {
	var node *Node
	for node = linkedList.headNode; node != nil; node = node.nextNode {
		fmt.Println(node.property)
	}
}

// LastNode method returns the last Node
func (linkedList *LinkedList) LastNode() *Node {
	var node *Node
	var lastNode *Node
	for node = linkedList.headNode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			lastNode = node
		}
	}
	return lastNode
}

// AddToEnd method adds the node with property to the end
func (linkedList *LinkedList) AddToEnd(property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil

	var lastNode *Node
	lastNode = linkedList.LastNode()
	if lastNode != nil {
		lastNode.nextNode = node
	}
}

// NodeWithValue method returns Node given parameter property
func (linkedList *LinkedList) NodeWithValue(property int) *Node {
	var node *Node
	var nodeWith *Node
	for node = linkedList.headNode; node != nil; node = node.nextNode {
		if node.property == property {
			nodeWith = node
			break
		}
	}

	return nodeWith
}

// AddAfter method  adds a node with nodeProperty after node with property
func (linkedList *LinkedList) AddAfter(nodeProperty int, property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil

	var nodeWith *Node

	nodeWith = linkedList.NodeWithValue(nodeProperty)
	if nodeWith != nil {
		node.nextNode = nodeWith.nextNode
		nodeWith.nextNode = node
	}
}

func main() {
	var linkedList LinkedList
	linkedList = LinkedList{}
	linkedList.AddToHead(1)
	linkedList.AddToHead(3)
	linkedList.AddToEnd(5)
	linkedList.AddAfter(1, 7)
	linkedList.IterateList()

}
