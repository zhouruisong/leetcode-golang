package main

import (
	"fmt"
	"main/model"
)

/*
给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。

示例:

输入: [1,2,3,null,5,null,4]
输出: [1, 3, 4]
解释:

   1            <---
 /   \
2     3         <---
 \     \
  5     4       <---
*/
func rightSideView(root *model.TreeNode) []int {
	if root == nil {
		return []int{}
	}

	ret := []int{}
	p := root
	//队列
	nodes := []*model.TreeNode{p}

	for len(nodes) != 0 {
		n := len(nodes)
		//取最右侧数据
		index := n - 1
		ret = append(ret, nodes[index].Val)
		for i := 0; i < n; i++ {
			node := nodes[0]
			nodes = nodes[1:]
			if node.Left != nil {
				nodes = append(nodes, node.Left)
			}

			if node.Right != nil {
				nodes = append(nodes, node.Right)
			}
		}
	}

	return ret
}

func main() {
	root := &model.TreeNode{Val: 1}
	root.Left = &model.TreeNode{Val: 2}
	root.Left.Right = &model.TreeNode{Val: 5}
	root.Right = &model.TreeNode{Val: 3}
	root.Right.Right = &model.TreeNode{Val: 4}
	fmt.Println(rightSideView(root))
}
