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

func levelOrderBottom(root *model.TreeNode) [][]int {
	var ret [][]int
	if root == nil {
		return ret
	}

	p := root
	nodes := []*model.TreeNode{p} //队列

	for len(nodes) != 0 {
		var tmp []int
		n := len(nodes)

		for i := 0; i < n; i++ {
			node := nodes[0]
			nodes = nodes[1:]
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				nodes = append(nodes, node.Left)
			}
			if node.Right != nil {
				nodes = append(nodes, node.Right)
			}
		}

		ret = append(ret, tmp)
	}

	n := len(ret)
	res := [][]int{}
	for i := n - 1; i >= 0; i-- {
		res = append(res, ret[i])
	}
	return res
}

func main() {
	root := &model.TreeNode{
		Val: 5,
	}
	tree := model.InitBinaryTree(root)
	fmt.Println(levelOrderBottom(tree))
}
