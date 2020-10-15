package main

import (
	"fmt"
	"main/model"
)

/*
给定一个非空二叉树，返回其最大路径和。

本题中，路径被定义为一条从树中任意节点出发，达到任意节点的序列。该路径至少包含一个节点，且不一定经过根节点。

示例 1:

输入: [1,2,3]

       1
      / \
     2   3

输出: 6
示例 2:

输入: [-10,9,20,null,null,15,7]

   -10
   / \
  9  20
    /  \
   15   7

输出: 42
*/

/*
解题思路：
简化一棵树：根，左子树，右子树
发明一个词：“根的单边路径”：从根节点出发的路径（此路径中每个节点深度都不同）
那么一个树的“最大路径和”可能的情况有这么几种：

只有根
完全在左子树中
完全在右子树中
根 + 以 根.left 为根的最大单边路径和
根 + 以 根.right 为根的最大单边路径和
根 + 以 根.left 为根的最大单边路径和 + 以 根.right 为根的最大单边路径和
如上，一棵树，只要知道如下4值，即可求得其最大路径和：

左子树最大路径和
右子树最大路径和
左子树的单边最大路径和
右子树的单边最大路径和
以上4个值可以递归计算，递归截止于叶子节点，4个值都是节点本身的值
*/

func maxPathSum(root *model.TreeNode) int {
	max, _ := maxPathSumImpl(root)
	return max
}

// 返回值意义：最大路径和，根的单边最大路径和
func maxPathSumImpl(root *model.TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}

	if root.Left == nil && root.Right == nil {
		return root.Val, root.Val
	}

	if root.Left != nil && root.Right != nil {
		maxL, maxSideL := maxPathSumImpl(root.Left)
		maxR, maxSideR := maxPathSumImpl(root.Right)
		maxSide := maxInt(maxSideL+root.Val, maxSideR+root.Val, root.Val)
		return maxInt(maxL, maxR, root.Val+maxSideL+maxSideR, maxSide), maxSide
	} else if root.Left != nil {
		maxL, maxSideL := maxPathSumImpl(root.Left)
		return maxInt(maxL, root.Val+maxSideL, root.Val), maxInt(maxSideL+root.Val, root.Val)
	}
	maxR, maxSideR := maxPathSumImpl(root.Right)
	return maxInt(maxR, root.Val+maxSideR, root.Val), maxInt(maxSideR+root.Val, root.Val)
}

func maxInt(a, b int, cs ...int) int {
	if a < b {
		a = b
	}
	for _, each := range cs {
		if a < each {
			a = each
		}
	}
	return a
}

func main() {
	root := &model.TreeNode{Val: -10}
	root.Left = &model.TreeNode{Val: 9}
	root.Right = &model.TreeNode{
		Val:   20,
		Left:  &model.TreeNode{Val: 15},
		Right: &model.TreeNode{Val: 7}}

	fmt.Println(maxPathSum(root))
}
