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

func main() {
	root := &model.TreeNode{
		Val: 5,
	}
	tree := model.InitBinaryTree(root)
	fmt.Println(zigzagLevelOrder(tree))
}
