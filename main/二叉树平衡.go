package main

import (
	"fmt"
	"main/model"
	"math"
)

/*
给定一个二叉树，判断它是否是高度平衡的二叉树。

本题中，一棵高度平衡二叉树定义为：

一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1。

示例 1:

给定二叉树 [3,9,20,null,null,15,7]

    3
   / \
  9  20
    /  \
   15   7
返回 true 。

示例 2:

给定二叉树 [1,2,2,3,3,null,null,4,4]

       1
      / \
     2   2
    / \
   3   3
  / \
 4   4
返回 false 。
*/

func isBalanced(root *model.TreeNode) bool {
	//从上到下遍历节点，计算每个节点的左右高度，求高度差
	if root == nil {
		return true
	}

	right := maxDepth(root.Right)
	left := maxDepth(root.Left)
	nodeBalanced := false
	if math.Abs(float64(left-right)) <= float64(1) {
		nodeBalanced = true
	}

	return nodeBalanced && isBalanced(root.Left) && isBalanced(root.Right)
}

func maxDepth(root *model.TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	heigth := left
	if left < right {
		heigth = right
	}

	return heigth + 1
}

func main() {
	root := &model.TreeNode{
		Val: 5,
	}
	tree := model.InitBinaryTree(root)
	fmt.Println(isBalanced(tree))
}
