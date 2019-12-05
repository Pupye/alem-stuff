package piscine

//BTreeTransplant is a function good ..
func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if node.Parent == nil { // node here is root
		node = rplc
		if node != nil {
			node.Parent = nil
		}

		return root
	}

	if node.Parent.Left == node { //node is left child of its parent
		node.Parent.Left = rplc
		if rplc != nil {
			rplc.Parent = node.Parent
		}
		return root
	}

	if node.Parent.Right == node {
		node.Parent.Right = rplc
		if rplc != nil {
			rplc.Parent = node.Parent
		}
		return root
	}

	return root
}
