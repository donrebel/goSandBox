package datastructures

import (
	"bytes"
	"strconv"
)

// LinkedList ...
type LinkedList struct {
	Head, Tail *Node
	Size       int
}

// Node ...
type Node struct {
	Data int
	Next *Node
}

// ToString ...
func (n *Node) ToString() string {
	var b bytes.Buffer
	if n == nil {
		return "[]"
	}
	b.WriteString("[")
	cnt := 0
	for n != nil {
		if cnt > 0 {
			b.WriteString(", ")
		}
		b.WriteString(strconv.Itoa(n.Data))
		n = n.Next
		cnt++
		if cnt == 50 {
			break
		}
	}
	b.WriteString("]")
	return b.String()
}

// Append ..
func (l *LinkedList) Append(d int) {
	n := &Node{d, nil}
	if l.IsEmpty() {
		l.Head = n
		l.Tail = n
		l.Size = 1
	} else {
		l.Tail.Next = n
		l.Tail = n
		l.Size++
	}
}

// Push ...
func (l *LinkedList) Push(d int) {
	n := &Node{d, nil}
	if l.IsEmpty() {
		l.Head = n
		l.Tail = n
		l.Size = 1
	} else {
		n.Next = l.Head
		l.Head = n
		l.Size++
	}
}

// IsEmpty ...
func (l *LinkedList) IsEmpty() bool {
	return l.Head == nil
}

// ToString ...
func (l *LinkedList) ToString() string {
	var b bytes.Buffer

	if l.IsEmpty() {
		return "[]"
	}
	n := l.Head
	cnt := 0
	b.WriteString("[")
	for n != nil {
		if cnt > 0 {
			b.WriteString(", ")
		}
		b.WriteString(strconv.Itoa(n.Data))
		n = n.Next
		cnt++
	}
	b.WriteString("]")
	return b.String()
}

// SortedInsert ...
func (l *LinkedList) SortedInsert(d int) {
	t := &Node{d, nil}
	if l.Head == nil || l.Head.Data > d {
		t.Next = l.Head
		l.Head = t
		l.Size++
	}
	cn := l.Head
	for cn.Next != nil && cn.Data < d {
		cn = cn.Next
	}
	t.Next = cn.Next
	cn.Next = t
	l.Size++
}
