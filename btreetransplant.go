package student

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	current := node
	if node.Parent == nil {
		root = rplc
	} else if node == node.Parent.Right {
		current.Parent.Right = rplc
	} else if node == node.Parent.Left {
		current.Parent.Left = rplc
	}
	current.Parent = node.Parent

	return root
}
