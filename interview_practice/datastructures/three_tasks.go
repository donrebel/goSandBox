package datastructures

import "fmt"

// NthPreOrder - returns value in nth position while PreOrder traversing of the tree
func NthPreOrder(tn *TNode, index int) (int, bool) {
	cnt := 0
	found := false
	res := 0
	nthPreOrder(tn, index, &cnt, &found, &res)
	return res, found
}

func nthPreOrder(tn *TNode, index int, cnt *int, found *bool, res *int) {
	if tn == nil {
		return
	}
	if *cnt == index {
		*res = tn.Data
		*found = true
	}
	(*cnt)++
	nthPreOrder(tn.Left, index, cnt, found, res)
	nthPreOrder(tn.Right, index, cnt, found, res)
}

// CopyTree ...
func CopyTree(t *Tree) *Tree {
	tres := &Tree{}
	tres.Root = copyTree(t.Root)
	return tres
}

func copyTree(tn *TNode) *TNode {
	if tn != nil {
		n := &TNode{}
		n.Data = tn.Data
		n.Left = copyTree(tn.Left)
		n.Right = copyTree(tn.Right)
		return n
	}
	return nil
}

//GetNumberNodes ...
func GetNumberNodes(t *Tree) int {
	return countNodes(t.Root)
}

func countNodes(tn *TNode) int {
	if tn == nil {
		return 0
	}
	return 1 + countNodes(tn.Left) + countNodes(tn.Right)
}

//GetNumberOfLeafNodes ...
func GetNumberOfLeafNodes(t *Tree) int {
	return countLeafs(t.Root)
}

func countLeafs(tn *TNode) int {
	if tn == nil {
		return 0
	}
	if tn.Left == nil && tn.Right == nil {
		return 1
	}
	return countLeafs(tn.Left) + countLeafs(tn.Right)
}

//IsIdentical ...
func IsIdentical(t1 *Tree, t2 *Tree) bool {
	return isIdentical(t1.Root, t2.Root)
}

func isIdentical(t1 *TNode, t2 *TNode) bool {
	switch {
	case t1 == nil && t2 == nil:
		return true
	case t1 == nil || t2 == nil:
		return false
	default:
		return t1.Data == t2.Data && isIdentical(t1.Left, t2.Left) && isIdentical(t1.Right, t2.Right)
	}
}

//PrintAllPath ...
func PrintAllPath(t *Tree) {
	stk := &Stack{}
	printAllPath(t.Root, stk)
}

func printAllPath(tn *TNode, stk *Stack) {
	if tn == nil {
		return
	}
	stk.Push(tn.Data)
	if tn.Left == nil && tn.Right == nil {
		fmt.Println(stk.ToString())
		stk.Pop()
		return
	}
	printAllPath(tn.Left, stk)
	printAllPath(tn.Right, stk)
	stk.Pop()
}

//GetLeastCommonAncestor ...
func GetLeastCommonAncestor(t *Tree, first int, second int) (int, bool) {
	if checkIfValuesExistInTree(t, first, second) == false {
		return 0, false
	}

	res := lcaUtil2(t.Root, first, second)
	if res != nil {
		return res.Data, true
	}
	fmt.Println("NotFoundError")
	return 0, false
}

func lcaUtil(curr *TNode, first int, second int) *TNode {
	var left, right *TNode
	if curr == nil {
		return nil
	}
	if curr.Data == first || curr.Data == second {
		return curr
	}

	left = lcaUtil(curr.Left, first, second)
	right = lcaUtil(curr.Right, first, second)

	if left != nil && right != nil {
		return curr
	} else if left != nil {
		return left
	} else {
		return right
	}
}

func lcaUtil2(t *TNode, first int, second int) *TNode {
	if t == nil {
		return nil
	}
	if t.Data == first || t.Data == second {
		return t
	}
	isFirstOnLeft := covers(t.Left, first)
	isSecondOnLeft := covers(t.Left, second)
	// if first and second are on different sides - return t
	if isFirstOnLeft != isSecondOnLeft {
		return t
	}
	var childSide *TNode
	if isFirstOnLeft {
		childSide = t.Left
	} else {
		childSide = t.Right
	}
	return lcaUtil2(childSide, first, second)
}

//BSTGetLeastCommonAncestor ...
func BSTGetLeastCommonAncestor(t *Tree, a int, b int) (int, bool) {
	return bstLeastCommonAncestor(t.Root, a, b)
}

func bstLeastCommonAncestor(tn *TNode, a int, b int) (int, bool) {
	if tn == nil {
		fmt.Println("Not Found")
		return 0, false
	}
	if tn.Data > a && tn.Data > b {
		return bstLeastCommonAncestor(tn.Left, a, b)
	}
	if tn.Data < a && tn.Data < b {
		return bstLeastCommonAncestor(tn.Right, a, b)
	}
	return tn.Data, true
}

func checkIfValuesExistInTree(t *Tree, ds ...int) bool {
	res := true
	for _, d := range ds {
		if covers(t.Root, d) == false {
			fmt.Println("Value ", d, " does not exist in the tree.")
			res = false
		}
	}
	return res
}

// find value in a subtree
func covers(root *TNode, d int) bool {
	if root == nil {
		return false
	} else if root.Data == d {
		return true
	}
	return covers(root.Left, d) || covers(root.Right, d)
}

//GetExtraNode returns Max or Min node of the tree depends on comp function parameter
func GetExtraNode(t *Tree, comp func(int, int) bool) int {
	extr := t.Root.Data
	getExtraNode(t.Root, &extr, comp)
	return extr
}

func getExtraNode(n *TNode, extr *int, comp func(int, int) bool) {
	if n == nil {
		return
	}
	if comp(n.Data, *extr) {
		*extr = n.Data
	}
	getExtraNode(n.Left, extr, comp)
	getExtraNode(n.Right, extr, comp)
}

//Search ...
func Search(t *Tree, x int) bool {
	return covers(t.Root, x)
}

//GetMaxDepth ...
func GetMaxDepth(t *Tree) int {
	return getMaxDepth(t.Root)
}

func getMaxDepth(tn *TNode) int {
	if tn == nil {
		return 0
	}
	lMaxDepth := 1 + getMaxDepth(tn.Left)
	rMaxDepth := 1 + getMaxDepth(tn.Right)
	if lMaxDepth > rMaxDepth {
		return lMaxDepth
	}
	return rMaxDepth
}

//GetNumOfFullNodes ...
func GetNumOfFullNodes(t *Tree) int {
	return cntNumOfFullNodes(t.Root)
}

func cntNumOfFullNodes(tn *TNode) int {
	if tn == nil {
		return 0
	}
	cnt := cntNumOfFullNodes(tn.Left) + cntNumOfFullNodes(tn.Right)
	if tn.Left != nil && tn.Right != nil {
		cnt++
	}
	return cnt
}

//GetTreeDiameter ...
func GetTreeDiameter(t *Tree) int {
	return getTreeDiameter(t.Root)
}

func getTreeDiameter(tn *TNode) int {
	if tn == nil {
		return 0
	}
	currentNodeDiameter := 1 + getMaxDepth(tn.Left) + getMaxDepth(tn.Right)
	leftNodeDiameter := getTreeDiameter(tn.Left)
	rightNodeDiameter := getTreeDiameter(tn.Right)
	if leftNodeDiameter > currentNodeDiameter {
		return leftNodeDiameter
	} else if rightNodeDiameter > currentNodeDiameter {
		return rightNodeDiameter
	}
	return currentNodeDiameter
}

//SumOfAllNodes ...
func SumOfAllNodes(t *Tree) int {
	return sumOfAllNodes(t.Root)
}

func sumOfAllNodes(tn *TNode) int {
	if tn == nil {
		return 0
	}
	return tn.Data + sumOfAllNodes(tn.Left) + sumOfAllNodes(tn.Right)
}
