package main

import (
	"fmt"
	"main/model"
)

/*
 * 根据一棵树的前序遍历与中序遍历构造二叉树。
 *
 * 注意:
 * 你可以假设树中没有重复的元素。
 *
 * 例如，给出
 *
 * 前序遍历 preorder = [3,9,20,15,7]
 * 中序遍历 inorder = [9,3,15,20,7]
 *
 * 返回如下的二叉树：
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *model.TreeNode {
	/*
		思路：
		由先序序列第一个pre[0]在中序序列中找到根节点位置gen
		以gen为中心遍历
		0~gen左子树
		子中序序列：0~gen-1，放入vin_left[]
		子先序序列：1~gen放入pre_left[]，+1可以看图，因为头部有根节点
		gen+1~vinlen为右子树
		子中序序列：gen ~ vinlen-1放入vin_right[]
		子先序序列：gen ~ vinlen-1放入pre_right[]
		由先序序列pre[0]创建根节点
		连接左子树，按照左子树子序列递归（pre_left[]和vin_left[]）
		连接右子树，按照右子树子序列递归（pre_right[]和vin_right[]）
		返回根节点*/
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	//查找根节点顺序
	var root int
	for k, v := range inorder {
		if v == preorder[0] {
			root = k
			break
		}
	}

	return &model.TreeNode{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:root+1], inorder[0:root]),
		Right: buildTree(preorder[root+1:], inorder[root+1:]),
	}
}

//中序遍历
func InOrder(current *model.TreeNode) {
	if current != nil {
		InOrder(current.Left)
		fmt.Printf(" %v", current.Val)
		InOrder(current.Right)
	}
}

func PreOrder(current *model.TreeNode) {
	if current != nil {
		fmt.Printf(" %v", current.Val)
		PreOrder(current.Left)
		PreOrder(current.Right)
	}
}

func main() {
	root := buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	PreOrder(root)
	fmt.Println("\n======")
	InOrder(root)
}
