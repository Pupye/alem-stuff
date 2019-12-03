package piscine

//BTreeMax is a function
func BTreeMax(root *TreeNode) *TreeNode {
	current := root

	for current.Right != nil {
		current = current.Right
	}

	return current
}
