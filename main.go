package main

import (
	"fmt"
	ds "interview_practice/datastructures"
)

func main() {
	arr := []int{6, 4, 5, 2, 1, 8, 2, 9, 7, 10, 3}
	fmt.Println(arr)
	t := ds.BinarySearchTree(arr)
	// t := ds.LevelOrderBinaryTree(arr)
	fmt.Println("preOrder: ", t.ToString("preOrder"))
	fmt.Println("inOrder: ", t.ToString("inOrder"))
	fmt.Println("postOrder: ", t.ToString("postOrder"))
	fmt.Println("levelOrder: ", t.ToString("levelOrder"))
	fmt.Println()

	fmt.Println(ds.SumOfAllNodes(t))
}
