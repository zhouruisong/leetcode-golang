package main

import (
	"container/list"
	"main/model"
)

/*
给定一个二叉搜索树，编写一个函数 kthSmallest 来查找其中第 k 个最小的元素。

说明：
你可以假设 k 总是有效的，1 ≤ k ≤ 二叉搜索树元素个数。

示例 1:

输入: root = [3,1,4,null,2], k = 1
   3
  / \
 1   4
  \
   2
输出: 1
示例 2:

输入: root = [5,3,6,2,4,null,null,1], k = 3
       5
      / \
     3   6
    / \
   2   4
  /
 1
输出: 3
进阶：
如果二叉搜索树经常被修改（插入/删除操作）并且你需要频繁地查找第 k 小的值，你将如何优化 kthSmallest 函数？
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *model.TreeNode, k int) int {
	//中序遍历二叉树后,得到的结果是升序排列，只需要中序遍历k次即可
	if root == nil {
		return 0
	}

	nodelist := list.New()

	res := 0
	n := 0
	//循环k次
	for n < k && (root != nil || nodelist.Len() != 0) {
		for root != nil {
			nodelist.PushFront(root)
			root = root.Left
		}

		if nodelist.Len() != 0 {
			root = nodelist.Remove(nodelist.Front()).(*model.TreeNode)
			res = root.Val
			root = root.Right
		}
		n++
	}

	return res
}
