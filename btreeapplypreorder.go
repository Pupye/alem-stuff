package piscine

//BTreeApplyPreorder ...
func BTreeApplyPreorder(root *TreeNode, f func(...interface{}) (int, error)) {
	f(root.Data)

	if root.Left != nil {
		BTreeApplyInorder(root.Left, f)
	}

	if root.Right != nil {
		BTreeApplyInorder(root.Right, f)
	}

}
