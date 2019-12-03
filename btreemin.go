package piscine

//BTreeMin ...
func BTreeMin(root *TreeNode) *TreeNode {
	current := root

	for current.Left != nil {
		current = current.Left
	}
	return current
}
