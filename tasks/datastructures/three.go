package datastructures

import (
	"bytes"
	"sort"
	"strconv"
)

// TNode ...
type TNode struct {
	Data  int
	Left  *TNode
	Right *TNode
}

// Tree ...
type Tree struct {
	Root *TNode
	Size int
}

// LevelOrderBinaryTree ...
func LevelOrderBinaryTree(arr []int) *Tree {
	tree := &Tree{}
	sort.Slice(arr, func(i int, j int) bool { return arr[i] < arr[j] })
	tree.Root = levelOrderBinaryTree(arr, 0, len(arr))
	return tree
}

func levelOrderBinaryTree(arr []int, start int, size int) *TNode {
	curr := &TNode{arr[start], nil, nil}
	left := 2*start + 1
	right := 2*start + 2
	if left < size {
		curr.Left = levelOrderBinaryTree(arr, left, size)
	}
	if right < size {
		curr.Right = levelOrderBinaryTree(arr, right, size)
	}
	return curr
}

//BinarySearchTree ...
func BinarySearchTree(arr []int) *Tree {
	tree := &Tree{}
	sort.Slice(arr, func(i int, j int) bool { return arr[i] < arr[j] })
	tree.Root = binarySearchTree(arr, 0, len(arr)-1)
	return tree
}

func binarySearchTree(arr []int, start int, end int) *TNode {
	if start > end {
		return nil
	}
	mid := (start + end) / 2
	n := &TNode{}
	n.Data = arr[mid]
	n.Left = binarySearchTree(arr, start, mid-1)
	n.Right = binarySearchTree(arr, mid+1, end)
	return n
}

//Add ...
func (t *Tree) Add(x int) {
	t.Root = add(t.Root, x)
}

func add(tn *TNode, x int) *TNode {
	if tn == nil {
		n := &TNode{}
		n.Data = x
		return n
	}
	if x < tn.Data {
		tn.Left = add(tn.Left, x)
	} else {
		tn.Right = add(tn.Right, x)
	}
	return tn
}

// IsEmpty ...
func (t *Tree) IsEmpty() bool {
	return t.Root == nil
}

//Free ...
func (t *Tree) Free() {
	t.Root = nil
}

// GetDepth ...
func (t *Tree) GetDepth() int {
	return cntDepth(t.Root)
}

func cntDepth(tn *TNode) int {
	if tn == nil {
		return 0
	}
	lDepth := cntDepth(tn.Left)
	rDepth := cntDepth(tn.Right)
	if lDepth > rDepth {
		return lDepth + 1
	}
	return rDepth + 1
}

// ToString ...
func (t *Tree) ToString(order string) string {
	var b bytes.Buffer
	cnt := 0
	b.WriteString("[")
	switch order {
	case "preOrder":
		preOrder(&b, &cnt, t.Root)
	case "inOrder":
		inOrder(&b, &cnt, t.Root)
	case "postOrder":
		postOrder(&b, &cnt, t.Root)
	case "levelOrder":
		levelOrder(&b, &cnt, t.Root)
	}
	b.WriteString("]")
	return b.String()
}

func preOrder(b *bytes.Buffer, cnt *int, tn *TNode) {
	if tn == nil {
		return
	}
	if *cnt > 0 {
		b.WriteString(" ")
	}
	b.WriteString(strconv.Itoa(tn.Data))
	*cnt++
	preOrder(b, cnt, tn.Left)
	preOrder(b, cnt, tn.Right)
}

func inOrder(b *bytes.Buffer, cnt *int, tn *TNode) {
	if tn == nil {
		return
	}
	inOrder(b, cnt, tn.Left)
	if *cnt > 0 {
		b.WriteString(" ")
	}
	b.WriteString(strconv.Itoa(tn.Data))
	*cnt++
	inOrder(b, cnt, tn.Right)
}

func postOrder(b *bytes.Buffer, cnt *int, tn *TNode) {
	if tn == nil {
		return
	}
	postOrder(b, cnt, tn.Left)
	postOrder(b, cnt, tn.Right)
	if *cnt > 0 {
		b.WriteString(" ")
	}
	b.WriteString(strconv.Itoa(tn.Data))
	*cnt++
}

func levelOrder(b *bytes.Buffer, cnt *int, tn *TNode) {
	q := &Queue{}
	if tn != nil {
		q.Append(tn)
	}
	for q.IsEmpty() == false {
		temp, _ := q.Get()
		n := temp.(*TNode)
		if *cnt > 0 {
			b.WriteString(" ")
		}
		*cnt++
		b.WriteString(strconv.Itoa(n.Data))
		if n.Left != nil {
			q.Append(n.Left)
		}
		if n.Right != nil {
			q.Append(n.Right)
		}
	}
}
