package sll

import "log"

func TrySll() {

	l1 := new(LinkedList)
	l1.Add(1)
	l1.Add(2)
	l1.Add(3)
	l1.View()
}

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

func (ll *LinkedList) Add(value int) {
	newNode := new(Node)
	newNode.Value = value
	newNode.Next = nil

	if ll.Head == nil {
		ll.Head = newNode
		return
	}

	iterator := ll.Head
	for ; iterator.Next != nil; iterator = iterator.Next {
	}
	iterator.Next = newNode
}

func (ll LinkedList) View() {
	if ll.Head == nil {
		log.Println("linked list empty")
		return
	}

	for iterator := ll.Head; iterator != nil; iterator = iterator.Next {
		log.Println("value ", iterator.Value)
	}
}
