package trees

import (
	"datastructures/lists"
	"fmt"

	constraints "golang.org/x/exp/constraints"
)

type TreeNode2[T constraints.Ordered] struct {
	Data  T
	Left  *TreeNode2[T]
	Right *TreeNode2[T]
}

func (node *TreeNode2[T]) Find(data T) (*TreeNode2[T], *TreeNode2[T]) {
	return node.find(data, nil)
}

func (node *TreeNode2[T]) find(data T, parent *TreeNode2[T]) (*TreeNode2[T], *TreeNode2[T]) {

	if data < node.Data {
		return node.Left.find(data, node)
	} else if data > node.Data {
		return node.Right.find(data, node)
	} else {
		return node, parent
	}
}

func (tree *TreeNode2[T]) BFS(doAction func(node *TreeNode2[T])) {

	queue := lists.GenericQueue[*TreeNode2[T]]{}
	queue.Enqueue((tree))

	var node *TreeNode2[T]

	for !queue.IsEmpty() {
		node, _ = queue.Dequeue()

		if node != nil {
			queue.Enqueue(node.Left)
			queue.Enqueue(node.Right)
			doAction(node)
		}

	}
}

func (tree *TreeNode2[T]) DFS(doAction func(node *TreeNode2[T])) {

	stack := lists.GenericQueue[*TreeNode2[T]]{}
	stack.Enqueue((tree))

	var node *TreeNode2[T]

	for !stack.IsEmpty() {
		node, _ = stack.Pop()

		if node != nil {
			stack.Push(node.Right)
			stack.Push(node.Left)
			doAction(node)
		}

	}
}

func (node *TreeNode2[T]) Insert(data T) *TreeNode2[T] {

	if node == nil {
		node = &TreeNode2[T]{
			Data: data,
		}
	} else if data == node.Data {
		node = nil
	} else if data < node.Data {
		node.Left = node.Left.Insert(data)
	} else {
		node.Right = node.Right.Insert(data)
	}

	return node
}

func (tree *TreeNode2[T]) Remove(data T) bool {

	node, parent := tree.Find(data)

	if node == nil {
		return false
	}

	if parent == nil {
		tree = nil
	} else if parent.Left != nil && parent.Left.Data == data {
		parent.Left = node.Left
	} else if parent.Right != nil && parent.Right.Data == data {
		parent.Right = node.Right
	}

	return true

}

func (node *TreeNode2[T]) Insert2(data T) bool {

	if node == nil {
		return false
	}

	newNode := &TreeNode2[T]{
		Data: data,
	}

	next := node

	for next != nil {
		if data < next.Data {
			if next.Left == nil {
				next.Left = newNode
				break
			} else {
				next = next.Left
			}
		} else if data > next.Data {
			if next.Right == nil {
				next.Right = newNode
				break
			} else {
				next = next.Right
			}
		} else {
			return false
		}
	}

	return true
}

func (node *TreeNode2[T]) Print(space int) {

	if node == nil {
		return
	}

	// Increase distance between levels
	space += 10

	if node.Right != nil {
		node.Right.Print(space)
	}

	fmt.Print("\n")

	for i := 10; i < space; i++ {
		fmt.Print(" ")
	}

	fmt.Printf("%v\n", node.Data)

	if node.Left != nil {
		node.Left.Print(space)
	}
}

func (node *TreeNode2[T]) IsValidBST() bool {

	if node == nil {
		return true
	}

	if node.Right != nil && node.Right.Data < node.Data || node.Left != nil && node.Left.Data > node.Data {
		return false
	}

	return node.Right.IsValidBST() && node.Left.IsValidBST()
}
