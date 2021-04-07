package main

import (
	"fmt"
	"main/model"
)

/*
给定一个二叉树，返回其节点值的锯齿形层次遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

例如：
给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回锯齿形层次遍历如下：

[
  [3],
  [20,9],
  [15,7]
]
*/

func zigzagLevelOrder(root *model.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	p := root
	nodeList := []*model.TreeNode{p}
	var ret [][]int
	count := 2

	for len(nodeList) > 0 {
		var tmp []int
		n := len(nodeList)
		for i := 0; i < n; i++ {
			node := nodeList[0]
			nodeList = nodeList[1:]
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				nodeList = append(nodeList, node.Left)
			}
			if node.Right != nil {
				nodeList = append(nodeList, node.Right)
			}
		}

		if count%2 != 0 {
			l := 0
			r := len(tmp) - 1
			for l < r {
				tmp[l], tmp[r] = tmp[r], tmp[l]
				l++
				r--
			}
		}
		ret = append(ret, tmp)

		count++
	}

	return ret
}

func zlevelOrder(root *model.TreeNode) [][]int {
	// write code here
	//借助队
	if root == nil {
		return [][]int{}
	}

	queueNodes := []*model.TreeNode{root}
	res := [][]int{}

	count := 2
	for len(queueNodes) > 0 {
		n := len(queueNodes)
		tmp := []int{}
		for i := 0; i < n; i++ {
			node := queueNodes[0]
			queueNodes = queueNodes[1:]
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				queueNodes = append(queueNodes, node.Left)
			}
			if node.Right != nil {
				queueNodes = append(queueNodes, node.Right)
			}
		}
		if count%2 != 0 {
			i := 0
			j := len(tmp) - 1
			for i < j {
				tmp[i], tmp[j] = tmp[j], tmp[i]
				i++
				j--
			}
		}

		res = append(res, tmp)
		count++
	}

	return res
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
	fmt.Println(zigzagLevelOrder(tree))
	fmt.Println(zlevelOrder(gen()))
}

func gen() *model.TreeNode {
	root := &model.TreeNode{
		Val: 5,
	}
	root.Left = &model.TreeNode{Val: 4}
	root.Right = &model.TreeNode{Val: 8}

	root.Left.Left = &model.TreeNode{Val: 11}
	root.Left.Left.Left = &model.TreeNode{Val: 7}
	root.Left.Left.Right = &model.TreeNode{Val: 2}

	root.Right.Left = &model.TreeNode{Val: 13}
	root.Right.Right = &model.TreeNode{Val: 4}
	root.Right.Right.Left = &model.TreeNode{Val: 6}
	root.Right.Right.Right = &model.TreeNode{Val: 1}

	return root
}
