package datastructures

import (
	"bytes"
	"fmt"
)

// QNode ...
type QNode struct {
	Data interface{}
	Next *QNode
}

// Queue ...
type Queue struct {
	Head, Tail *QNode
	size       int
}

// Append ...
func (q *Queue) Append(d interface{}) {
	n := &QNode{d, nil}
	if q.IsEmpty() {
		q.Head = n
		q.Tail = n
	}
	q.Tail.Next = n
	q.Tail = q.Tail.Next
	q.size++
}

// Get ...
func (q *Queue) Get() (interface{}, bool) {
	if q.IsEmpty() {
		return 0, false
	}
	res := q.Head.Data
	q.Head = q.Head.Next
	q.size--
	return res, true
}

// Peek ...
func (q *Queue) Peek() (interface{}, bool) {
	if q.IsEmpty() {
		return 0, false
	}
	return q.Head.Data, true
}

// IsEmpty ...
func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

//GetSize ...
func (q *Queue) GetSize() int {
	return q.size
}

// ToString ...
func (q *Queue) ToString() string {
	if q.IsEmpty() {
		return "[]"
	}
	b := bytes.Buffer{}
	b.WriteString("[")
	n := q.Head
	cnt := 0
	for n != nil {
		if cnt > 0 {
			b.WriteString(" ")
		}
		b.WriteString(queueNodeConvToString(n.Data))
		cnt++
		n = n.Next
	}
	b.WriteString("]")
	return b.String()
}

func queueNodeConvToString(x interface{}) string {
	return fmt.Sprint(x)
}

// Fill ...
func (q *Queue) Fill(arr []int) {
	for _, x := range arr {
		q.Append(x)
	}
}
