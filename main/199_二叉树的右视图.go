package main

import "main/model"

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
	var nodes []*model.TreeNode
	nodes = append(nodes, p)

	for len(nodes) != 0 {
		old := nodes
		//取最右侧数据
		index := len(old) - 1
		ret = append(ret, old[index].Val)

		nodes = []*model.TreeNode{}
		for i := 0; i < len(old); i++ {
			if old[i].Left != nil {
				nodes = append(nodes, old[i].Left)
			}

			if old[i].Right != nil {
				nodes = append(nodes, old[i].Right)
			}
		}
	}

	return ret
}
