package main

import (

	// alg "interview_practice/algorithms"
	c "interview_practice/concurrency"
	ds "interview_practice/datastructures"
)

func main() {
	arr := []int{1, 9, 6, 7, 8, -1, 2, 4, 5, 3}
	// fmt.Println(arr)
	hp := ds.NewHeap(nil, true)
	// hp := ds.NewHeap(arr, true)
	n := len(arr)
	for i := 0; i < n; i++ {
		hp.Add(arr[i])
	}
	// for i := 0; i < n; i++ {
	// 	fmt.Println("pop value :: ", ds.Slice(hp.Remove())[0])
	// }
	c.ExecPattern()
}
