package main

import (
	"dsa/sll"
	"log"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	n := sll.LinkedList{}

	n.AddBeginning(1)
	n.AddBeforeValue(1, 4)
	n.AddBeforeValue(1, 3)
	n.AddBeforeValue(1, 2)
	n.AddBeforeValue(5, 13)
	n.Display()

	n.DeleteValue(7)
	n.AddAfterValue(7, 12)
	log.Println()
	n.Display()
}
