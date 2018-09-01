package datastructures

import (
	"fmt"
)

// RemoveDuplicates ...
func RemoveDuplicates(ll *Node) {
	n := ll
	m := make(map[int]bool)
	m[n.Data] = true
	for n.Next != nil {
		if _, found := m[n.Next.Data]; found {
			n.Next = n.Next.Next
		} else {
			m[n.Next.Data] = true
			n = n.Next
		}
	}
}

// ReverseList ...
func ReverseList(l *LinkedList, ln *Node) {
	var prev, next, curr *Node
	prev = nil
	curr = ln
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	fmt.Println(prev.ToString())
	l.Head = prev
	ln = prev
}

// SortLinkedListSelectionSort ...
func SortLinkedListSelectionSort(ll *LinkedList) {
	if ll.Size == 1 {
		return
	}
	var max int
	var nmax, cn *Node
	for i := 0; i < ll.Size; i++ {
		max = ll.Head.Data
		nmax = ll.Head
		cn = ll.Head
		for j := 0; j < ll.Size-i-1; j++ {
			if cn.Data > max {
				max = cn.Data
				nmax = cn
			}
			cn = cn.Next
		}
		cn.Data, nmax.Data = nmax.Data, cn.Data
	}
}

// SortLinkedListBubbleSort ...
func SortLinkedListBubbleSort(ll *LinkedList) {
	if ll.Size < 2 {
		return
	}
	var cn *Node
	flip := true
	for i := 0; i < ll.Size && flip; i++ {
		cn = ll.Head
		flip = false
		for j := 0; j < ll.Size-i-1; j++ {
			if cn.Data > cn.Next.Data {
				flip = true
				cn.Data, cn.Next.Data = cn.Next.Data, cn.Data
			}
			cn = cn.Next
		}
	}
}

// MergeSort ...
func MergeSort(l *LinkedList) {
	h := mergeSortUtil(l.Head)
	l.Head = h
}

func mergeSortUtil(h *Node) *Node {
	if h == nil || h.Next == nil {
		return h
	}
	middle := getMiddle(h)
	nextToMiddle := middle.Next
	middle.Next = nil
	left := mergeSortUtil(h)
	right := mergeSortUtil(nextToMiddle)
	sortedList := sortedMerge(left, right)
	return sortedList
}

func getMiddle(h *Node) *Node {
	if h == nil {
		return h
	}
	fastPtr := h.Next
	slowPtr := h
	for fastPtr != nil {
		fastPtr = fastPtr.Next
		if fastPtr != nil {
			slowPtr = slowPtr.Next
			fastPtr = fastPtr.Next
		}
	}
	return slowPtr
}

func sortedMerge(a *Node, b *Node) *Node {
	var result *Node
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	if a.Data <= b.Data {
		result = a
		result.Next = sortedMerge(a.Next, b)
	} else {
		result = b
		result.Next = sortedMerge(b.Next, a)
	}
	return result
}
