package trees

import (
	"fmt"
)

type BinaryTree[T any] struct {
	Root *TreeNode[T]
	Cmp  func(T, T) int
}

type TreeNode[T any] struct {
	Data  T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

func (tree *BinaryTree[T]) Find(data T) **TreeNode[T] {
	pt := &tree.Root

	for *pt != nil {
		switch cmp := tree.Cmp(data, (*pt).Data); {
		case cmp > 0:
			pt = &(*pt).Right
		case cmp < 0:
			pt = &(*pt).Left
		default:
			return pt
		}
	}

	return pt
}

func (tree *BinaryTree[T]) Insert(data T) bool {
	node := tree.Find(data)

	if *node != nil {
		return false
	}

	*node = &TreeNode[T]{
		Data: data,
	}

	return true
}

func (tree *BinaryTree[T]) Print() {
	if tree.Root == nil {
		fmt.Printf("Tree is empty")
	}

	tree.Root.printSubTree(0)
}

func (node *TreeNode[T]) printSubTree(space int) {

	if node == nil {
		return
	}

	// Increase distance between levels
	space += 10

	if node.Right != nil {
		node.Right.printSubTree(space)
	}

	fmt.Print("\n")

	for i := 10; i < space; i++ {
		fmt.Print(" ")
	}

	fmt.Printf("%v\n", node.Data)

	if node.Left != nil {
		node.Left.printSubTree(space)
	}
}
