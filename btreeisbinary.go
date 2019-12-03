package piscine

//BTreeIsBinary ..
func BTreeIsBinary(root *TreeNode) bool {
	if root.Left == nil && root.Right == nil {
		return true
	}

	if root.Left != nil && root.Right != nil {
		if root.Left.Data < root.Data && root.Data < root.Right.Data {
			return BTreeIsBinary(root.Left) && BTreeIsBinary(root.Right)
		}
		return false
	} else if root.Left != nil {
		if root.Left.Data < root.Data {
			return BTreeIsBinary(root.Left)
		}
		return false
	} else {
		if root.Data < root.Right.Data {
			return BTreeIsBinary(root.Right)
		}
		return false
	}
}
