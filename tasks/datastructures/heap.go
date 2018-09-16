package datastructures

import (
	"fmt"
)

//Heap ...
type Heap struct {
	size  int
	arr   []int
	isMin bool
}

//NewHeap ...
func NewHeap(arrInput []int, isMin bool) *Heap {
	size := len(arrInput)
	arr := []int{0} //first element will not be used it is needed for indexing
	arr = append(arr, arrInput...)
	h := &Heap{size: size, arr: arr, isMin: isMin}
	for i := (h.size / 2); i > 0; i-- {
		h.proclateDown(i)
	}
	return h
}

//NewHeap2 ...
func NewHeap2(isMin bool) *Heap {
	arr := []int{0}
	h := &Heap{size: 0, arr: arr, isMin: isMin}
	return h
}

func (h *Heap) comp(i, j int) bool {
	if h.isMin == true {
		return h.arr[i] > h.arr[j]
	}
	return h.arr[i] < h.arr[j]
}

func (h *Heap) swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}

//Add ...
func (h *Heap) Add(x int) {
	h.size++
	h.arr = append(h.arr, x)
	h.proclateUp(h.size)
}

//Remove ...
func (h *Heap) Remove() (int, bool) {
	if h.IsEmpty() {
		fmt.Println("HeapEmptyError")
		return 0, false
	}
	val := h.arr[1]
	h.arr[1] = h.arr[h.size]
	h.size--
	h.proclateDown(1)
	h.arr = h.arr[0 : h.size+1]
	return val, true
}

func (h *Heap) proclateUp(child int) {
	parent := child / 2
	if parent == 0 {
		return
	}
	if h.comp(parent, child) {
		h.swap(child, parent)
		h.proclateUp(parent)
	}
}

func (h *Heap) proclateDown(parentIdx int) {
	lChildIdx := 2 * parentIdx
	rChildIdx := lChildIdx + 1
	idxOfSmallerChild := -1
	if lChildIdx <= h.size {
		idxOfSmallerChild = lChildIdx
	}
	// if we build a MinHeap:
	// if rChildIdx < h.size and h.arr[rCrildIdx] < h.arr[lCrildIdx]
	if rChildIdx <= h.size && h.comp(lChildIdx, rChildIdx) {
		idxOfSmallerChild = rChildIdx
	}
	// if we build a MinHeap:
	// if there is a Child and h.arr[parentIdx] > h.arr[idxOfSmallerChild]
	if idxOfSmallerChild != -1 && h.comp(parentIdx, idxOfSmallerChild) {
		h.swap(parentIdx, idxOfSmallerChild)
		h.proclateDown(idxOfSmallerChild)
	}
}

//IsEmpty ...
func (h *Heap) IsEmpty() bool {
	return h.size == 0
}

//Size ...
func (h *Heap) Size() int {
	return h.size
}

//Peek ...
func (h *Heap) Peek() (int, bool) {
	if h.IsEmpty() {
		fmt.Println("Heap empty Error.")
		return 0, false
	}
	return h.arr[1], true
}
