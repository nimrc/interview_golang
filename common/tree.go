package common

/*
 * 二叉树 Binary Tree
 */

type TreeNode struct {
	Value  string
	Parent *TreeNode
	Left   *TreeNode
	Right  *TreeNode
}

func NewBinaryTree(values []string) *TreeNode {
	return createSubTree(nil, values, 0)
}

func createSubTree(parent *TreeNode, values []string, idx int) *TreeNode {
	if idx >= len(values) || values[idx] == "" {
		return nil
	}

	node := &TreeNode{Value: values[idx], Parent: parent}

	node.Left = createSubTree(node, values, idx*2+1)
	node.Right = createSubTree(node, values, idx*2+2)

	return node
}
