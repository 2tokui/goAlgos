package algo

import "fmt"

type LinkedList struct {
	Head *NodeS
	Tail *NodeS
}

type DoubleLinkedList struct {
	Head *NodeD
	Tail *NodeD
}

// singly linked list
type NodeS struct {
	Next *NodeS
	Data *any
}

// singly linked list
type NodeD struct {
	Next     *NodeD
	Previous *NodeD
	Data     *any
}

func NewLinkedList() *LinkedList {
	return &LinkedList{nil, nil}
}

func (ll *LinkedList) Pop(element any) {
	node := ll.Head
	if node == nil {
		return
	}

	// check head and tail first
	if *ll.Head.Data == element {
		ll.Head = ll.Head.Next
	}

	// check what is in the middle
	var lastGuy *NodeS
	for {
		if node == nil {
			return
		}
		if *node.Data == element {
			lastGuy.Next = node.Next
			node = nil
			return
		}

		lastGuy = node
		node = node.Next
	}
}

func (ll *LinkedList) Push(data any) {
	node := &NodeS{}
	node.Data = &data
	node.Next = nil

	if ll.Head == nil {
		ll.Head = node
		ll.Tail = node
	} else {
		ll.Tail.Next = node
		ll.Tail = node
	}
}

func (ll *LinkedList) PrintEverything() {
	node := ll.Head
	if node == nil {
		return
	}

	var data any

	for {
		if node == nil {
			return
		}

		data = *node.Data
		fmt.Println(data)

		node = node.Next
	}
}

func NewDoubleLinkedList() *DoubleLinkedList {
	return &DoubleLinkedList{}
}
