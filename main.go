package main

import (
	"fmt"
	ds "interview_practice/tasks/datastructures"
)

func main() {
	l := &ds.LinkedList{}
	input := []int{6, 3, 2, 1, 5, 4, 9, 8, 7, 10}
	for x := range input {
		l.Append(x)
	}
	fmt.Println(l.ToString())
	ds.ReverseList2(&l.Head)
	fmt.Println(l.ToString())
}
