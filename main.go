package main

import (
	listTypes "datastructures/lists"
	treeTypes "datastructures/trees"
	"fmt"
)

func main() {
	fmt.Println(("Hello world!"))

	testLists()
	testTrees()

}

func testLists() {
	rootNode := listTypes.Node{Value: 10, Next: nil, Prev: nil}

	list := listTypes.LinkedList{Head: &rootNode, Tail: &rootNode, Length: 1}

	list.Print()
	fmt.Println("\n------------")

	list.Append(21)
	list.Prepend(8)
	list.Prepend(11)
	list.Prepend(12)
	list.Insert(99, 2)
	list.Remove(4)

	list.Print()

	fmt.Println("\n------------")

	node := list.Pop()

	fmt.Println("pop", node.Value)

	list.Push(999)
	node = list.Peek()

	fmt.Println("peek", node.Value)

	list.Print()
	fmt.Println("\n------------")

	node = list.Dequeue()

	fmt.Println("dequeue", node.Value)

	list.Enqueue(888)

	list.Print()

	queue := listTypes.GenericQueue[string]{}

	for i := 0; i < 5; i++ {
		queue.Enqueue("hej")
	}

	queue.Print()

	for i := 0; i < 5; i++ {
		str, err := queue.Dequeue()
		if err != nil {
			fmt.Println((err.Error()))
		} else {
			fmt.Println(str)
		}

	}

	queue.Print()
}

func testTrees() {
	// tree := treeTypes.BinaryTree[int]{
	// 	Root: &treeTypes.TreeNode[int]{
	// 		Data: 1,
	// 		Left: &treeTypes.TreeNode[int]{
	// 			Data: 2,
	// 			Left: &treeTypes.TreeNode[int]{
	// 				Data:  4,
	// 				Left:  nil,
	// 				Right: nil,
	// 			},
	// 			Right: &treeTypes.TreeNode[int]{
	// 				Data:  5,
	// 				Left:  nil,
	// 				Right: nil,
	// 			},
	// 		},
	// 		Right: &treeTypes.TreeNode[int]{
	// 			Data: 3,
	// 			Left: &treeTypes.TreeNode[int]{
	// 				Data:  6,
	// 				Left:  nil,
	// 				Right: nil,
	// 			},
	// 			Right: &treeTypes.TreeNode[int]{
	// 				Data:  7,
	// 				Left:  nil,
	// 				Right: nil,
	// 			},
	// 		},
	// 	},

	// 	Cmp: func(a, b int) int {
	// 		return a - b
	// 	},
	// }

	// tree := treeTypes.BinaryTree[int]{
	// 	Root: &treeTypes.TreeNode[int]{
	// 		Data:  10,
	// 		Left:  nil,
	// 		Right: nil,
	// 	},

	// 	Cmp: func(a, b int) int {
	// 		return a - b
	// 	},
	// }

	tree := treeTypes.TreeNode2[int]{
		Data: 10,
	}

	tree.Insert2(13)
	tree.Insert2(5)
	tree.Insert2(7)
	tree.Insert2(6)
	tree.Insert2(12)
	tree.Insert2(15)

	var child, parent = tree.Find(15)

	fmt.Printf("Find node: %v - %v \n", *child, *parent)

	fmt.Println("\n------------")

	tree.Print(0)

	fmt.Println("\n------------")

	tree.BFS(func(node *treeTypes.TreeNode2[int]) {
		fmt.Printf("-- %v ", node.Data)
	})

	fmt.Println("\n------------")

	tree.DFS(func(node *treeTypes.TreeNode2[int]) {
		fmt.Printf("-- %v ", node.Data)
	})

	fmt.Println("\n------------")

	validBST := tree.IsValidBST()

	fmt.Printf("valid BST: %v", validBST)

	fmt.Println("\n------------")

	node, _ := tree.Find(15)
	node.Data = 0

	validBST2 := tree.IsValidBST()

	tree.Print(0)

	fmt.Printf("valid BST: %v", validBST2)
}
