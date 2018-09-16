package concepts

import (
	"fmt"
	ds "interview_practice/tasks/datastructures"
)

// Smartphone ...
type Smartphone struct {
	name string
}

// Geek ...
type Geek struct {
	smartphone *Smartphone
}

func replaceByNG(s **Smartphone) {
	*s = &Smartphone{"Galaxy Nexus"}
}

func replaceByIPhone(s *Smartphone) {
	s = &Smartphone{"IPhone 4S"}
}

// ReverseList1 is an exanple of incorrect usage of pointers. This function wont reverse the list.
func ReverseList1(h *ds.Node) {
	var prev, next, curr *ds.Node
	curr = h
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	h = prev
}

// ReverseList2 ...
func ReverseList2(h **ds.Node) {
	var prev, next, curr *ds.Node
	curr = *h
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	*h = prev
}

// ReverseList3 ...
func ReverseList3(l *ds.LinkedList) {
	var prev, next, curr *ds.Node
	curr = l.Head
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	l.Head = prev
}

// Test ...
func Test() {
	geek := Geek{&Smartphone{"Nexus S"}}
	println(geek.smartphone.name)

	replaceByIPhone(geek.smartphone)
	println(geek.smartphone.name)

	replaceByNG(&geek.smartphone)
	println(geek.smartphone.name)

	fmt.Println()

	l := &ds.LinkedList{}
	l.Push(1)
	l.Push(2)
	l.Push(3)
	fmt.Println(l.Head.ToString())
	ReverseList2(&l.Head)
	fmt.Println(l.Head.ToString())
	ReverseList3(l)
	fmt.Println(l.Head.ToString())
	ReverseList1(l.Head)
	fmt.Println(l.Head.ToString())
}
