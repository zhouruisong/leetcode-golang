package main

import (
	"main/model"
	"fmt"
)

/*
翻转一棵二叉树。

示例：

输入：

     4
   /   \
  2     7
 / \   / \
1   3 6   9
输出：

     4
   /   \
  7     2
 / \   / \
9   6 3   1

*/
//递归
func invertTree(root *model.TreeNode) *model.TreeNode {
	if root == nil {
		return nil
	}

	left := root.Left
	right := root.Right

	root.Left = right
	root.Right = left

	invertTree(root.Left)
	invertTree(root.Right)

	return root
}

func PreOrder(current *model.TreeNode) {
	if current != nil {
		fmt.Printf(" %v", current.Val)
		PreOrder(current.Left)
		PreOrder(current.Right)
	}
}

func main() {
	root := &model.TreeNode{
		Val: 3,
	}
	tree := model.InitBinaryTree2(root)
	PreOrder(tree)
	invertTree(tree)
	fmt.Println("\n======")
	PreOrder(tree)

}
