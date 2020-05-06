package main

import (
	"fmt"
	"main/model"
)

/*
给定一个不含重复元素的整数数组。一个以此数组构建的最大二叉树定义如下：
二叉树的根是数组中的最大元素。
左子树是通过数组中最大值左边部分构造出的最大二叉树。
右子树是通过数组中最大值右边部分构造出的最大二叉树。
通过给定的数组构建最大二叉树，并且输出这个树的根节点。
示例 ：
输入：[3,2,1,6,0,5]
输出：返回下面这棵树的根节点：

      6
    /   \
   3     5
    \    /
     2  0
       \
        1

提示：
给定的数组的大小在 [1, 1000] 之间。
*/

func constructMaximumBinaryTree(nums []int) *model.TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	}

	return constructMaximumBinaryTreeHelp(nums, 0)
}

func constructMaximumBinaryTreeHelp(nums []int, start int) *model.TreeNode {
	if start < 0 {
		return nil
	}

	n := len(nums)
	if n == 0 {
		return nil
	}

	max := nums[0]
	loc := 0
	for i := 1; i < n; i++ {
		if nums[i] > max {
			max = nums[i]
			loc = i
		}
	}

	root := &model.TreeNode{Val: max}
	root.Left = constructMaximumBinaryTreeHelp(nums[:loc], start)
	root.Right = constructMaximumBinaryTreeHelp(nums[loc+1:], loc+1)
	return root
}

//先序遍历递归
func PreOrder(current *model.TreeNode) {
	if current != nil {
		fmt.Printf(" %v", current.Val)
		PreOrder(current.Left)
		PreOrder(current.Right)
	}
}

func InOrder(current *model.TreeNode) {
	if current != nil {
		InOrder(current.Left)
		fmt.Printf(" %v", current.Val)
		InOrder(current.Right)
	}
}

//后序遍历递归
func PostOrder(current *model.TreeNode) {
	if current != nil {
		PostOrder(current.Left)
		PostOrder(current.Right)
		fmt.Printf(" %v", current.Val)
	}
}

func main() {
	x := []int{3, 2, 1, 6, 0, 5}
	tree := constructMaximumBinaryTree(x)
	PreOrder(tree)
	fmt.Println("=====")
	InOrder(tree)
	fmt.Println("=====")
	PostOrder(tree)
}
