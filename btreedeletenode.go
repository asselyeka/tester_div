package student

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if node == nil {
		return root
	}
	if node.Data < root.Data {
		root.Left = BTreeDeleteNode(root.Left, node)
	} else if node.Data > root.Data {
		root.Right = BTreeDeleteNode(root.Right, node)
	} else {
		if root.Left == nil {
			current := root.Right
			root = nil
			return current
		} else if root.Right == nil {
			current := root.Left
			root = nil
			return current
		}
		current := BTreeMin(root.Right)
		root.Data = current.Data
		root.Right = BTreeDeleteNode(root.Right, current)
	}
	return root
}

func BTreeMin(root *TreeNode) *TreeNode {
	if root == nil || root.Left == nil {
		return root
	}
	return BTreeMin(root.Left)
}
func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil || root.Data == elem {
		return root
	}
	if elem < root.Data {
		return BTreeSearchItem(root.Left, elem)
	}
	return BTreeSearchItem(root.Right, elem)
}
func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	//BTreeApplyInorder(BTreeApplyInorder(root.Right), f)
	BTreeApplyInorder(root.Left, f)
	f(root.Data)
	BTreeApplyInorder(root.Right, f)
}
