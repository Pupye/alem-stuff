package piscine

//BTreeDeleteNode ...
func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if node.Left == nil && node.Right == nil {
		if node.Parent == nil {
			root = nil
			return root
		}
		BTreeTransplant(root, node, nil)
		return root
	}

	if node.Left != nil && node.Right != nil {
		inOrderSucc := BTreeMin(node.Right)
		node.Data = inOrderSucc.Data
		if inOrderSucc.Left != nil {
			BTreeTransplant(root, inOrderSucc, inOrderSucc.Left)
			return root
		}
		BTreeTransplant(root, inOrderSucc, nil)
		return root
	}

	if node.Left != nil {
		BTreeTransplant(root, node, node.Left)
		return root
	}

	if node.Right != nil {
		BTreeTransplant(root, node, node.Right)
		return root
	}
	return root
}
