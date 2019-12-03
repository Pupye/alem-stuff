package piscine

//BTreeLevelCount ...
func BTreeLevelCount(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return 1
	}

	leftLevel := 0
	if root.Left != nil {
		leftLevel = BTreeLevelCount(root.Left)
	}
	rightLevel := 0
	if root.Right != nil {
		rightLevel = BTreeLevelCount(root.Right)
	}

	if leftLevel > rightLevel {
		return leftLevel + 1
	}
	return rightLevel + 1
}
