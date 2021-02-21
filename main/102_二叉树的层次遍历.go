package main

import (
	"container/list"
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
			if old[i].Right != nil {
				nodes = append(nodes, old[i].Right)
			}
		}

		ret = append(ret, tmp)
	}

	return ret
}

func levelOrderBottom2(root *model.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := [][]int{}
	p := root
	qu := list.New()
	qu.PushBack(p)

	for qu.Len() > 0 {
		ln := qu.Len()
		tmp := []int{}
		for i := 0; i < ln; i++ {
			top := qu.Remove(qu.Front()).(*model.TreeNode)
			tmp = append(tmp, top.Val)
			if top.Left != nil {
				qu.PushBack(top.Left)
			}
			if top.Right != nil {
				qu.PushBack(top.Right)
			}
		}

		if len(tmp) > 0 {
			res = append(res, tmp)
		}
	}

	return res
}

func main() {
	root := &model.TreeNode{
		Val: 5,
	}
	tree := model.InitBinaryTree(root)

	//PreOrder(tree)
	fmt.Println("\n=====")
	//InOrder(tree)
	fmt.Println("\n=====")
	//
	//PostOrder(tree)

	fmt.Println(levelOrderBottom(tree))
	fmt.Println(levelOrderBottom2(tree))
}
