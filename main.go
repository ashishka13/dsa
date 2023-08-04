package main

import (
	"dsa/dll"
	"log"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	// n := sll.LinkedList{}

	// n.AddBeginning(1)
	// n.AddBeforeValue(1, 4)
	// n.AddBeforeValue(1, 3)
	// n.AddBeforeValue(1, 2)
	// n.AddBeforeValue(5, 13)
	// n.Display()

	// n.DeleteValue(7)
	// n.AddAfterValue(7, 12)
	// log.Println()
	// n.Display()

	n2 := dll.DoublyLinkedList{}
	n2.AddAfter(1)
	n2.AddAfter(2)
	n2.AddAfter(3)
	// n2.AddAfter(4)
	// n2.AddAfter(5)

	// n2.AddBefore(1)
	// n2.AddBefore(2)
	// n2.AddBefore(3)

	n2.AddBeforeValue(2, 12)

	n2.Display()
}
