package lists

import (
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
	Prev  *Node
}

type LinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

func (list LinkedList) Print() {
	node := list.Head

	for {
		fmt.Print(node.Value)
		node = node.Next

		if node != nil {
			fmt.Print(" -- ")
		} else {
			break
		}
	}
}

func (list LinkedList) GetNode(index int) *Node {
	node := list.Head

	for i := 1; i <= index; i++ {
		node = node.Next
	}

	return node
}

func (list *LinkedList) Append(Value int) {
	newNode := Node{Value: Value, Next: nil, Prev: list.Tail}
	list.Tail.Next = &newNode
	list.Tail = &newNode
	list.Length++
}

func (list *LinkedList) Prepend(Value int) {
	newNode := Node{Value: Value, Next: list.Head, Prev: nil}
	list.Head.Prev = &newNode
	list.Head = &newNode
	list.Length++
}

func (list *LinkedList) Insert(Value int, index int) {
	node := list.GetNode(index)
	newNode := Node{Value: Value, Next: node, Prev: node.Prev}
	node.Prev.Next = &newNode
	node.Prev = &newNode
	list.Length++
}

func (list *LinkedList) Remove(index int) {
	node := list.GetNode(index)
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	list.Length--
}

func (list *LinkedList) Push(Value int) {
	list.Prepend(Value)
}

func (list *LinkedList) Pop() *Node {
	node := list.Head
	node.Next.Prev = nil
	list.Head = node.Next
	list.Length--
	return node
}

func (list *LinkedList) Peek() *Node {
	return list.Head
}

func (list *LinkedList) Enqueue(Value int) {
	list.Append(Value)
}

func (list *LinkedList) Dequeue() *Node {
	return list.Pop()
}
