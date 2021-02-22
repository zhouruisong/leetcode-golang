package main

import (
	"container/list"
	"fmt"
	"main/model"
)

/*
* 给定一个二叉树，它的每个结点都存放一个 0-9 的数字，每条从根到叶子节点的路径都代表一个数字。
 *
 * 例如，从根到叶子节点路径 1->2->3 代表数字 123。
 *
 * 计算从根到叶子节点生成的所有数字之和。
 *
 * 说明: 叶子节点是指没有子节点的节点。
 *
 * 示例 1:
 *
 * 输入: [1,2,3]
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   3
 * 输出: 25
 * 解释:
 * 从根到叶子节点路径 1->2 代表数字 12.
 * 从根到叶子节点路径 1->3 代表数字 13.
 * 因此，数字总和 = 12 + 13 = 25.
 *
 * 示例 2:
 *
 * 输入: [4,9,0,5,1]
 * ⁠   4
 * ⁠  / \
 * ⁠ 9   0
 * / \
 * 5   1
 * 输出: 1026
 * 解释:
 * 从根到叶子节点路径 4->9->5 代表数字 495.
 * 从根到叶子节点路径 4->9->1 代表数字 491.
 * 从根到叶子节点路径 4->0 代表数字 40.
 * 因此，数字总和 = 495 + 491 + 40 = 1026.
*/

func sumNumbers(root *model.TreeNode) int {
	if root == nil {
		return 0
	}

	nodeList := list.New()
	pathList := list.New()

	nodeList.PushFront(root)
	pathList.PushFront([]int{root.Val})

	max := 0
	for nodeList.Len() > 0 {
		currLen := nodeList.Len()
		for i := 0; i < currLen; i++ {
			node := nodeList.Remove(nodeList.Back()).(*model.TreeNode)
			path := pathList.Remove(pathList.Back()).([]int)

			tmpleft := []int{}
			tmprigth := []int{}
			tmpleft = append(tmpleft, path...)
			tmprigth = append(tmprigth, path...)

			if node.Left == nil && node.Right == nil {
				//fmt.Println(path)
				//字符串转数字
				tmp := 0
				for _, v := range path {
					tmp = tmp*10 + v
				}

				max += tmp
			}

			if node.Left != nil {
				nodeList.PushFront(node.Left)
				pathList.PushFront(append(tmpleft, node.Left.Val))
			}

			if node.Right != nil {
				nodeList.PushFront(node.Right)
				pathList.PushFront(append(tmprigth, node.Right.Val))
			}
		}
	}

	return max
}

func pow10(n int) int {
	pow := 1
	x := 10
	for n > 0 {
		if n&1 == 1 {
			pow = pow * x
		}

		x = x * x
		n = n >> 1
	}

	return pow
}

func main() {
	root := &model.TreeNode{
		Val: 4,
	}

	l := model.TreeNode{}
	r := model.TreeNode{}
	root.Left = l.NewBinaryTreeNode(9)
	root.Right = r.NewBinaryTreeNode(0)

	root.Left.Left = (&model.TreeNode{}).NewBinaryTreeNode(5)
	root.Left.Right = (&model.TreeNode{}).NewBinaryTreeNode(1)

	fmt.Println(sumNumbers(root))

	fmt.Println(pow10(0))
}
