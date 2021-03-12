package main

import (
	"fmt"
	"main/model"
)

/* 给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。
*
*
*
* 示例：
* 二叉树：[3,9,20,null,null,15,7],
*
*
* ⁠   3
* ⁠  / \
* ⁠ 9  20
* ⁠   /  \
* ⁠  15   7
*
*
* 返回其层序遍历结果：
*
*
* [
* ⁠ [3],
* ⁠ [9,20],
* ⁠ [15,7]
* ]
*
*
 */

func levelOrder(root *model.TreeNode) [][]int {
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
			if old[i].Right != nil {
				nodes = append(nodes, old[i].Right)
			}
		}

		ret = append(ret, tmp)
	}

	return ret
}

func main() {
	root := &model.TreeNode{
		Val: 5,
	}
	/*
			     5
			  4     8
			11   13    4
		  7	  2      6   1

	*/

	tree := model.InitBinaryTree(root)
	fmt.Println(levelOrder(tree))
}
