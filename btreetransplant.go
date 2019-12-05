package piscine

//BTreeTransplant is a function good ..
func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if node.Parent == nil {
		root = rplc
		return root
	}
	if node.Parent.Left == node {
		node.Parent.Left = rplc
		return root
	}
	if node.Parent.Right == node {
		node.Parent.Right = rplc
		return root
	}
	return root
}
