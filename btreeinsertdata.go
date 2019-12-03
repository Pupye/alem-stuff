package piscine

//TreeNode ...
type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

//BTreeInsertData ....
func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root.Data > data {
		if root.Left != nil {
			return BTreeInsertData(root.Left, data)
		}
		n := &TreeNode{Data: data, Parent: root}

		root.Left = n
		return n

	}

	if root.Right != nil {
		return BTreeInsertData(root.Right, data)
	}
	n := &TreeNode{Data: data, Parent: root}
	root.Right = n
	return n

}
