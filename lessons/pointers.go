package main

import (
	"fmt"
)

type Node struct {
	val  int
	next *Node
}

func main() {
	var newNode Node = Node{5, nil}
	var head *Node = &newNode
	fmt.Println(head.val)
	fmt.Println(head.next)
}
