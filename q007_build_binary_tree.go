package main

import (
	"fmt"
	. "github.com/fyibmsd/interview/common"
)

/*
 * Q: 通过先序遍历和中序遍历重建二叉树
 *
 * 思路：递归左右子树
 */
func buildBinaryTree(preOrder, inOrder []int) *TreeNode {
	if len(preOrder) == 0 || len(inOrder) == 0 {
		return nil
	}

	rootValue := preOrder[0]
	rootPos := 0

	for inOrder[rootPos] != rootValue {
		rootPos++
	}

	root := &TreeNode{
		Value: preOrder[0],
		Left:  buildBinaryTree(preOrder[1:rootPos+1], inOrder[:rootPos]),
		Right: buildBinaryTree(preOrder[rootPos+1:], inOrder[rootPos+1:]),
	}

	return root
}

func main() {
	preOrder := []int{1, 2, 4, 7, 3, 5, 6, 8}
	inOrder := []int{4, 7, 2, 1, 5, 3, 8, 6}

	tree := buildBinaryTree(preOrder, inOrder)

	fmt.Println(tree)
}
