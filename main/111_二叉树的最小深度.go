package main

import (
	"fmt"
	"main/model"
	"math"
)

/*
给定一个二叉树，找出其最小深度。

最小深度是从根节点到最近叶子节点的最短路径上的节点数量。

说明: 叶子节点是指没有子节点的节点。

示例:

给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回它的最小深度  2.
*/

func minDepth(root *model.TreeNode) int {
	if root == nil {
		return 0
	}
	if (root.Left == nil) && (root.Right == nil) {
		return 1
	}
	min_Depth := math.MaxInt16
	if root.Left != nil {
		min_Depth = int(math.Min(float64(minDepth(root.Left)), float64(min_Depth)))
	}
	if root.Right != nil {
		min_Depth = int(math.Min(float64(minDepth(root.Right)), float64(min_Depth)))
	}
	return min_Depth + 1
}

func main() {
	root := &model.TreeNode{
		Val: 3,
	}
	tree := model.InitBinaryTree2(root)
	fmt.Println(minDepth(tree))
}
