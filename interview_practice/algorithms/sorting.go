package algorithms

import (
	ds "interview_practice/datastructures"
)

func BinarySearch(s []int, x int) bool {
	minIdx := 0
	maxIdx := len(s)
	var midIdx int
	for minIdx < maxIdx {
		midIdx = (minIdx + maxIdx) / 2
		// midIdx = minIdx + (maxIdx-minIdx)/2
		if s[midIdx] == x {
			return true
		}
		if s[midIdx] < x {
			minIdx = midIdx
		} else {
			maxIdx = midIdx
		}
	}
	return false
}

// HeapSort ...
func HeapSort(a []int) {
	h := ds.NewHeap(a, true)
	for i := 0; i < len(a); i++ {
		a[i], _ = h.Remove()
	}
}
