package main

import (
	"fmt"
	"main/model"
)

/*
将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树。

本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。

示例:

给定有序数组: [-10,-3,0,5,9],

一个可能的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：

      0
     / \
   -3   9
   /   /
 -10  5

*/

// 辅助函数 输入是原整数数组以及每次构建子树时需要用到的子数组的开始和结束下标，输出是构建出来的二叉搜索树
func arrayToBST(nums []int, start, end int) *model.TreeNode {
	if start > end { // 如果开始下标大于结束下标
		return nil // 返回空指针
	}
	// 否则计算开始下标和结束下标的中间值mid
	mid := (start + end) >> 1
	// 并用下标mid对应的数字构建一个树节点root
	root := &model.TreeNode{Val: nums[mid]}
	// 然后用mid左边的子数组构建root的左子树,这里更新结束下标为mid-1即可
	root.Left = arrayToBST(nums, start, mid-1)
	// 同理用mid右边的子数组构建root的右子树
	root.Right = arrayToBST(nums, mid+1, end)
	return root // 最后返回root即可
}

// 递归方法 Time: O(n), Space: O(log(n))
func sortedArrayToBST(nums []int) *model.TreeNode {
	if nums == nil { // 如果数组为空
		return nil // 返回空指针
	}
	return arrayToBST(nums, 0, len(nums)-1)
}

//先序遍历
func PreOrder(current *model.TreeNode) {
	if current != nil {
		fmt.Printf(" %v", current.Val)
		PreOrder(current.Left)
		PreOrder(current.Right)
	}
}

func main() {
	//root := &model.TreeNode{
	//	Val: 0,
	//}
	//tree := model.InitBinaryTree3(root)

	root := sortedArrayToBST([]int{-10, -3, 0, 5, 9})
	PreOrder(root)
}
