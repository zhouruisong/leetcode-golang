package main

import (
	"container/list"
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
	nodeList := []*model.TreeNode{}
	nodeList = append(nodeList, p)

	var ret [][]int
	count := 2
	for len(nodeList) > 0 {
		var tmp []int
		old := nodeList
		nodeList = []*model.TreeNode{}

		for i := 0; i < len(old); i++ {
			tmp = append(tmp, old[i].Val)
			if old[i].Left != nil {
				nodeList = append(nodeList, old[i].Left)
			}
			if old[i].Right != nil {
				nodeList = append(nodeList, old[i].Right)
			}
		}

		if count%2 == 0 {
			ret = append(ret, tmp)
		} else {
			l := 0
			r := len(tmp) - 1
			for l < r {
				tmp[l], tmp[r] = tmp[r], tmp[l]
				l++
				r--
			}
			ret = append(ret, tmp)
		}

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

	queueNodes := list.New()
	queueNodes.PushBack(root)
	res := [][]int{}

	count := 1
	for queueNodes.Len() > 0 {
		n := queueNodes.Len()
		tmp := []int{}
		for i := 0; i < n; i++ {
			node := queueNodes.Remove(queueNodes.Front()).(*model.TreeNode)
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				queueNodes.PushBack(node.Left)
			}
			if node.Right != nil {
				queueNodes.PushBack(node.Right)
			}
		}
		if len(tmp) > 0 {
			if count%2 == 0 {
				i := 0
				j := len(tmp) - 1
				for i < j {
					tmp[i], tmp[j] = tmp[j], tmp[i]
					i++
					j--
				}
			}

			res = append(res, tmp)
		}
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
