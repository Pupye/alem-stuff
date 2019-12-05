package piscine

//BTreeApplyByLevel ...
func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	h := BTreeLevelCount(root)
	for index := 1; index <= h; index++ {
		helper(root, f, index)
	}
}

func helper(node *TreeNode, f func(...interface{}) (int, error), level int) {
	if node == nil {
		return
	}
	if level == 1 {
		f(node.Data)
	} else if level > 1 {
		helper(node.Left, f, level-1)
		helper(node.Right, f, level-1)
	}
}
