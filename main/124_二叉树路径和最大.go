package main

import (
	"fmt"
	"main/model"
	"math"
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

叶节点 9、15、7 的最大贡献值分别为 9、15、7。
得到叶节点的最大贡献值之后，再计算非叶节点的最大贡献值。节点 20 的最大贡献值等于20+max(15,7)=35，
节点-10的最大贡献值等于−10+max(9,35)=25。
*/

func maxPathSum(root *model.TreeNode) int {
	max := math.MinInt32
	var dfs func(root *model.TreeNode) int

	f := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	dfs = func(root *model.TreeNode) int {
		if root == nil {
			return 0
		}

		//当前节点的左孩子节点的最大子路径和（注意舍弃掉和小于0的路径
		leftSum := f(0, dfs(root.Left))
		rightSum := f(0, dfs(root.Right))
		//更新最大路径和（最大左路径和+最大右路径和+当前节点值）
		max = f(leftSum+rightSum+root.Val, max)
		//返回从当前节点出发的最大左/右路径和
		return f(leftSum, rightSum) + root.Val
	}

	dfs(root)
	return max
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
