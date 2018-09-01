package datastructures

import (
	"bytes"
	"fmt"
)

//Stack ...
type Stack struct {
	Head *SNode
	size int
}

//SNode ...
type SNode struct {
	Data interface{}
	Next *SNode
}

//Fill ...
func (s *Stack) Fill(arr []int) {
	for _, x := range arr {
		s.Push(x)
	}
}

//Push ...
func (s *Stack) Push(x interface{}) {
	s.Head = &SNode{x, s.Head}
	s.size++
}

//Pop ...
func (s *Stack) Pop() interface{} {
	res := s.Head.Data
	s.Head = s.Head.Next
	s.size--
	return res
}

//Peek ...
func (s *Stack) Peek() interface{} {
	return s.Head.Data
}

//IsEmpty ...
func (s *Stack) IsEmpty() bool {
	return s.Head == nil
}

//GetSize ...
func (s *Stack) GetSize() int {
	return s.size
}

//ToString ...
func (s *Stack) ToString() string {
	if s.IsEmpty() {
		return "[]"
	}
	var b bytes.Buffer
	cnt := 0
	n := s.Head
	b.WriteString("[")
	for n != nil {
		if cnt > 0 {
			b.WriteString(" ")
		}
		b.WriteString(fmt.Sprint(n.Data))
		n = n.Next
		cnt++
	}
	b.WriteString("]")
	return b.String()
}
