package main

import (
	"fmt"
	"main/model"
)

/*
给定一个二叉树，返回其节点值自底向上的层次遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）

例如：
给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其自底向上的层次遍历为：

[
  [15,7],
  [9,20],
  [3]
]

 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }

*/

//层次遍历
func levelOrderBottom(root *model.TreeNode) [][]int {
	var ret [][]int
	if root == nil {
		return ret
	}

	p := root
	var nodes []*model.TreeNode
	nodes = append(nodes, p)

	for len(nodes) != 0 {
		var tmp []int
		old := nodes
		nodes = []*model.TreeNode{}

		for i := 0; i < len(old); i++ {
			tmp = append(tmp, old[i].Val)
			if old[i].Left != nil {
				nodes = append(nodes, old[i].Left)
			}
			if old[i].Left != nil {
				nodes = append(nodes, old[i].Right)
			}
		}

		ret = append(ret, tmp)
	}

	var result [][]int
	for i := len(ret) - 1; i >= 0; i-- {
		result = append(result, ret[i])
	}
	return result
}

func main() {
	root := &model.TreeNode{
		Val: 5,
	}
	tree := model.InitBinaryTree(root)

	PreOrder(tree)
	fmt.Println("\n=====")
	InOrder(tree)
	fmt.Println("\n=====")
	//
	PostOrder(tree)

	fmt.Println(levelOrderBottom(tree))
}
