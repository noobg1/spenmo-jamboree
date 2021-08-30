package bst

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

func findElem(root *TreeNode, value int) bool {
	// check for leaf node pointer
	if root == nil {
		return false
	}

	if value == root.value {
		return true
	} else if value < root.value {
		return findElem(root.left, value)
	} else {
		return findElem(root.right, value)
	}
}
