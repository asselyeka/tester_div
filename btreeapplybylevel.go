package student

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	levels := BTreeLevelCount(root)
	for i := 1; i <= levels; i++ {
		Levels(root, i, f)

	}
}
func Levels(root *TreeNode, level int, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	if level == 1 {
		f(root.Data)
	} else {
		Levels(root.Left, level-1, f)
		Levels(root.Right, level-1, f)
	}
}

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}
	counterL := BTreeLevelCount(root.Left)
	counterR := BTreeLevelCount(root.Right)
	if counterL > counterR {
		return counterL + 1
	}
	return counterR + 1
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}
	if data < root.Data {
		root.Left = BTreeInsertData(root.Left, data)
		root.Left.Parent = root
	} else {
		root.Right = BTreeInsertData(root.Right, data)
		root.Right.Parent = root
	}
	return root
}
