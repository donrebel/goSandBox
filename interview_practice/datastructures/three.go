package datastructures

import (
	"bytes"
	"fmt"
	"math"
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

//DeleteNode ...
func (t *Tree) DeleteNode(x int) {
	t.Root = deleteNode(t.Root, x)
}

func deleteNode(tn *TNode, x int) *TNode {
	var temp *TNode
	if tn != nil {
		if tn.Data == x {
			if tn.Left == nil && tn.Right == nil {
				return nil
			}
			if tn.Left == nil {
				temp = tn.Right
				return temp
			}
			if tn.Right == nil {
				temp = tn.Left
				return temp
			}
			maxNode := getMaxNode(tn.Left)
			maxValue := maxNode.Data
			tn.Data = maxValue
			tn.Left = deleteNode(tn.Left, maxValue)
		} else {
			if tn.Data > x {
				tn.Left = deleteNode(tn.Left, x)
			} else {
				tn.Right = deleteNode(tn.Right, x)
			}
		}
	}
	return tn
}

//Find ...
func (t *Tree) Find(x int) bool {
	return find(t.Root, x)
}

func find(tn *TNode, x int) bool {
	if tn == nil {
		return false
	}
	if tn.Data == x {
		return true
	} else if x < tn.Data {
		return find(tn.Left, x)
	} else {
		return find(tn.Right, x)
	}
}

//BSTCheck ...
func (t *Tree) BSTCheck() bool {
	return bstCheck(t.Root)
}

func bstCheck(tn *TNode) bool {
	if tn == nil {
		return true
	}
	if tn.Left != nil {
		return tn.Left.Data < tn.Data && bstCheck(tn.Left)
	}
	if tn.Right != nil {
		return tn.Right.Data >= tn.Data && bstCheck(tn.Right)
	}
	return true
}

//GetMin ...
func (t *Tree) GetMin() int {
	return getMin(t.Root)
}

func getMin(tn *TNode) int {
	if tn.Left == nil {
		return tn.Data
	}
	return getMin(tn.Left)
}

//GetMax ...
func (t *Tree) GetMax() int {
	return getMax(t.Root)
}

func getMax(tn *TNode) int {
	if tn.Right == nil {
		return tn.Data
	}
	return getMax(tn.Right)
}

func getMaxNode(tn *TNode) *TNode {
	if tn.Right == nil {
		return tn
	}
	return getMaxNode(tn.Right)
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

//BSTTrinOutsideDataRange ...
func (t *Tree) BSTTrinOutsideDataRange(min int, max int) {
	t.Root = trimOutsideDataRange(t.Root, min, max)
}

func trimOutsideDataRange(tn *TNode, min int, max int) *TNode {
	if tn == nil {
		return nil
	}
	tn.Left = trimOutsideDataRange(tn.Left, min, max)
	tn.Right = trimOutsideDataRange(tn.Right, min, max)
	if tn.Data < min {
		return tn.Right
	}
	if tn.Data > max {
		return tn.Left
	}
	return tn
}

//BSTPrintNodesInRange ...
func (t *Tree) BSTPrintNodesInRange(min int, max int) {
	printNodesInRange(t.Root, min, max)
}

func printNodesInRange(tn *TNode, min int, max int) {
	if tn == nil {
		return
	}
	printNodesInRange(tn.Left, min, max)
	if tn.Data >= min && tn.Data <= max {
		fmt.Printf("%v ", tn.Data)
	}
	printNodesInRange(tn.Right, min, max)
}

//BSTFloor ...
func (t *Tree) BSTFloor(x int) int {
	tn := t.Root
	floor := math.MinInt32
	for tn != nil {
		if tn.Data == x {
			floor = tn.Data
			break
		} else if tn.Data > x {
			tn = tn.Left
		} else {
			floor = tn.Data
			tn = tn.Right
		}
	}
	return floor
}

//BSTCeil ...
func (t *Tree) BSTCeil(x int) int {
	tn := t.Root
	ceil := math.MaxInt32
	for tn != nil {
		if tn.Data == x {
			ceil = tn.Data
		} else if tn.Data < x {
			tn = tn.Right
		} else {
			ceil = tn.Data
			tn = tn.Left
		}
	}
	return ceil
}
