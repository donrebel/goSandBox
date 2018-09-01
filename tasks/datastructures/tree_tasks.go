package datastructures

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
