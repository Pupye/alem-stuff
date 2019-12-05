package piscine

//BTreeDeleteNode ...
func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if node.Left == nil && node.Right == nil {
		if node.Parent == nil {
			return nil
		}
		root = BTreeTransplant(root, node, nil)
		return root
	}

	if node.Left != nil && node.Right != nil {
		inOrderSucc := BTreeMin(node.Right)
		node.Data = inOrderSucc.Data
		if inOrderSucc.Right != nil {
			root = BTreeTransplant(root, inOrderSucc, inOrderSucc.Right)
			return root
		}

		root = BTreeTransplant(root, inOrderSucc, nil)
		return root
	}

	if node.Left != nil {
		root = BTreeTransplant(root, node, node.Left)
		return root
	}

	if node.Right != nil {
		root = BTreeTransplant(root, node, node.Right)
		return root
	}
	return root
}
