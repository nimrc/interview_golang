package main

import (
	"fmt"
	. "github.com/fyibmsd/interview/common"
)

/*
 * Q: 指定节点求二叉树中序遍历的下一个节点
 */
func nextNode(node *TreeNode) string {
	if node == nil {
		return ""
	}

	next := node

	// 节点有右子树的情况
	if node.Right != nil {
		right := node.Right
		for right.Left != nil {
			right = right.Left
		}

		next = right
	} else if node.Parent != nil {
		// 无右子树、有父节点
		curr := node
		parent := node.Parent
		for parent != nil && curr == parent.Right {
			curr = parent
			parent = parent.Parent
		}

		next = curr.Parent
	}

	if next == nil {
		return ""
	}

	return next.Value
}

func main() {
	values := []string{"a", "b", "c", "d", "e", "f", "g", "", "", "h", "i"}
	tree := NewBinaryTree(values)

	// b 的下一个节点，Expect: h
	fmt.Println(nextNode(tree.Left))

	// a 的下一个节点，Expect: f
	fmt.Println(nextNode(tree))

	// d 的下一个节点，Expect: b
	fmt.Println(nextNode(tree.Left.Left))

	// f 的下一个节点，Expect: c
	fmt.Println(nextNode(tree.Right.Left))

	// i 的下一个节点，Expect: c
	fmt.Println(nextNode(tree.Left.Right.Right))

	// g 的下一个节点，Expect: ""
	fmt.Println(nextNode(tree.Right.Right))
}
