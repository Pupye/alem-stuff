package piscine

//TreeNode ...
type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

//BTreeInsertData ....
func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root.Data == data {
		//dothing
		return root
	}

	if root.Data > data {
		if root.Left != nil {
			BTreeInsertData(root.Left, data)
		} else {
			n := &TreeNode{Data: data, Parent: root}

			root.Left = n
			return root
		}

	} else if root.Right != nil {
		BTreeInsertData(root.Right, data)
	} else {
		n := &TreeNode{Data: data, Parent: root}
		root.Right = n
		return root
	}
	return root
}
