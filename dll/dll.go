package dll

import (
	"fmt"
	"log"
)

type Node struct {
	data     int
	previous *Node
	next     *Node
}

type DoublyLinkedList struct {
	head *Node
}

func (dll *DoublyLinkedList) AddAfter(value int) {
	newNode := &Node{data: value}

	if dll.head == nil {
		dll.head = newNode
		return
	}

	temp := dll.head
	for temp.next != nil {
		temp = temp.next
	}

	newNode.previous = temp
	temp.next = newNode
}

func (dll *DoublyLinkedList) AddBefore(value int) {
	newNode := &Node{data: value}

	if dll.head == nil {
		dll.head = newNode
		return
	}

	temp := dll.head

	for temp.previous != nil {
		temp = temp.previous
	}

	newNode.next = temp
	temp.previous = newNode
	dll.head = newNode
}

func (dll *DoublyLinkedList) AddAfterValue(afterValue, value int) {
	newNode := &Node{data: value}

	if dll.head.data == afterValue {
		newNode.previous = dll.head
		dll.head.next = newNode
		return
	}

	temp := dll.head

	for temp.next != nil || temp.next == nil {
		if temp.data == afterValue {
			log.Println("code was here")
			newNode.previous = temp
			newNode.next = temp.next
			temp.next = newNode
			return
		}
		if temp.next == nil {
			log.Println("value not found")
			return
		}
		temp = temp.next
	}
}

func (dll *DoublyLinkedList) AddBeforeValue(beforeValue, value int) {
	newNode := &Node{data: value}

	if dll.head.data == beforeValue {
		newNode.next = dll.head
		dll.head.previous = newNode
		dll.head = newNode
		return
	}

	temp := dll.head

	for temp.next != nil || temp.next == nil {
		if temp.data == beforeValue {
			newNode.next = temp
			newNode.previous = temp.previous
			temp.previous = newNode
			log.Println("code was here", newNode.next.data, newNode.previous.data, temp.previous.data)

			return
		}
		if temp.next == nil {
			log.Println("value not found")
			return
		}

		temp = temp.next
	}
}

func (dll DoublyLinkedList) Display() {
	temp := dll.head

	for temp.next != nil {
		fmt.Printf("-|%d|-", temp.data)
		temp = temp.next
	}

	fmt.Printf("-|%d|-", temp.data)
}
