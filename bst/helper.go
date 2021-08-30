package bst

func insertToBST(root *TreeNode, value int) *TreeNode {
	if root == nil {
		// create new node if it's nil
		return &TreeNode{value: value, left: nil, right: nil}
	}

	if (TreeNode{}) == *root {
		// case when the node is root (structs cant be assigned to nil)
		*root = TreeNode{value: value, left: nil, right: nil}
		return root
	}

	if value < root.value {
		root.left = insertToBST(root.left, value)
	} else {
		root.right = insertToBST(root.right, value)
	}

	return root
}

func createBST(treeElements []int) TreeNode {
	root := TreeNode{}

	for _, value := range treeElements {
		insertToBST(&root, value)
	}
	return root
}
