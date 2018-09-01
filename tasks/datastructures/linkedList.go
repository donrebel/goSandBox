package datastructures

import (
	"bytes"
	"fmt"
)

//LinkedList ...
type LinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

//Node ...
type Node struct {
	Data int
	Next *Node
}

//ToString ...
func (n *Node) ToString() string {
	var b bytes.Buffer
	ln := n
	b.WriteString("[")
	cnt := 0
	for ln != nil {
		if cnt > 0 {
			b.WriteString(", ")
		}
		b.WriteString(fmt.Sprint(ln.Data))
		ln = ln.Next
		cnt++
	}
	b.WriteString("]")
	return b.String()
}

//Push ...
func (l *LinkedList) Push(d int) {
	l.Head = &Node{d, l.Head}
	l.Size++
}

//Append ...
func (l *LinkedList) Append(d int) {
	t := &Node{d, nil}
	cn := l.Head
	if cn == nil {
		l.Head = t
		l.Size++
		return
	}
	for cn.Next != nil {
		cn = cn.Next
	}
	cn.Next = t
	l.Size++
}

//ToString ...
func (l *LinkedList) ToString() string {
	var b bytes.Buffer
	ptr := l.Head
	b.WriteString("[")
	cnt := 0
	for i := 0; i < l.Size; i++ {
		if cnt > 0 {
			b.WriteString(", ")
		}
		b.WriteString(fmt.Sprint(ptr.Data))
		cnt++
		ptr = ptr.Next
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
