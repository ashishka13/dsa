package sll

import (
	"fmt"
	"log"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (ll *LinkedList) AddLast(value int) {
	newNode := Node{data: value}

	if ll.head == nil {
		ll.head = &newNode
		return
	}

	temp := ll.head

	for temp.next != nil {
		temp = temp.next
	}

	temp.next = &newNode
}

func (ll *LinkedList) AddBeginning(value int) {
	newNode := Node{data: value}

	if ll.head == nil {
		ll.head = &newNode
		return
	}
	newNode.next = ll.head
	ll.head = &newNode
}

func (ll *LinkedList) DeleteValue(value int) {
	if ll.head == nil {
		return
	}
	if ll.head.data == value {
		tempHead := ll.head.next
		ll.head = nil
		ll.head = tempHead
		return
	}

	temp := ll.head
	previous := ll.head

	for temp.next != nil || temp.next == nil {
		if temp.data == value {
			previous.next = temp.next
			temp = nil
			return
		}

		if temp.next == nil {
			log.Println(value, "delete value not found")
			return
		}

		previous = temp
		temp = temp.next
	}
}

func (ll *LinkedList) AddAfterValue(afterValue, newValue int) {
	newNode := Node{data: newValue}

	if ll.head == nil {
		ll.head = &newNode
		return
	}

	temp := ll.head

	for temp.next != nil || temp.next == nil {
		if temp.data == afterValue {
			nextValue := temp.next
			newNode.next = nextValue
			temp.next = &newNode
			return
		}

		if temp.next == nil {
			log.Println(afterValue, "after value not found")
			return
		}

		temp = temp.next
	}
}

func (ll *LinkedList) AddBeforeValue(beforeValue, newValue int) {
	newNode := Node{data: newValue}

	if ll.head == nil || ll.head.data == beforeValue {
		newNode.next = ll.head
		ll.head = &newNode
		return
	}

	temp := ll.head
	oldValue := ll.head

	for temp.next != nil || temp.next == nil {
		if temp.data == beforeValue {
			newNode.next = temp
			oldValue.next = &newNode
			return
		}
		if temp.next == nil {
			log.Println(beforeValue, "before value not found")
			return
		}

		oldValue = temp
		temp = temp.next
	}
}

func (ll LinkedList) Display() {
	temp := ll.head

	for temp.next != nil {
		fmt.Printf("|%d|-", temp.data)
		temp = temp.next
	}
	fmt.Printf("|%d|\n", temp.data)
}
