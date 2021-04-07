package main

import (
	"container/list"
	"fmt"
	"main/model"
)

/*
给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。

说明: 叶子节点是指没有子节点的节点。

示例:
给定如下二叉树，以及目标和 sum = 22，

              5
             / \
            4   8
           /   / \
          11  13  4
         /  \    / \
        7    2  5   1
返回:

[
   [5,4,11,2],
   [5,8,4,5]
]
*/

func pathSum(root *model.TreeNode, sum int) [][]int {
	if root == nil {
		return [][]int{}
	}

	//遍历每个节点
	nodeList := list.New()
	sumList := list.New()
	pathList := list.New()

	nodeList.PushFront(root)
	sumList.PushFront(sum - root.Val)
	pathList.PushFront([]int{root.Val})

	var result [][]int
	for nodeList.Len() > 0 {

		currentLen := nodeList.Len()
		for i := 0; i < currentLen; i++ {
			node := nodeList.Remove(nodeList.Back()).(*model.TreeNode)
			sumLeft := sumList.Remove(sumList.Back()).(int)
			path := pathList.Remove(pathList.Back()).([]int)

			tmp1 := []int{}
			tmp1 = append(tmp1, path...)
			tmp2 := []int{}
			tmp2 = append(tmp2, path...)

			if node.Left == nil && node.Right == nil && sumLeft == 0 {
				result = append(result, path)
			}

			if node.Left != nil {
				nodeList.PushFront(node.Left)
				sumList.PushFront(sumLeft - node.Left.Val)
				tmp1 = append(tmp1, node.Left.Val)
				pathList.PushFront(tmp1)
			}

			if node.Right != nil {
				nodeList.PushFront(node.Right)
				sumList.PushFront(sumLeft - node.Right.Val)
				tmp2 = append(tmp2, node.Right.Val)
				pathList.PushFront(tmp2)
			}
		}
	}

	return result
}

func pathSum2(root *model.TreeNode, sum int) [][]int {
	if root == nil {
		return [][]int{}
	}

	nodeList := []*model.TreeNode{root}
	sumList := []int{sum - root.Val}
	pathList := [][]int{}
	pathList = append(pathList, []int{root.Val})

	result := [][]int{}
	for len(nodeList) > 0 {
		n := len(nodeList)
		for i := 0; i < n; i++ {
			node := nodeList[0]
			nodeList = nodeList[1:]
			sumLeft := sumList[0]
			sumList = sumList[1:]
			path := pathList[0]
			pathList = pathList[1:]

			tmp1 := append([]int{}, path...)
			tmp2 := append([]int{}, path...)

			if node.Left == nil && node.Right == nil && sumLeft == 0 {
				result = append(result, path)
			}

			if node.Left != nil {
				nodeList = append(nodeList, node.Left)
				sumList = append(sumList, sumLeft-node.Left.Val)
				tmp1 = append(tmp1, node.Left.Val)
				pathList = append(pathList, tmp1)
			}

			if node.Right != nil {
				nodeList = append(nodeList, node.Right)
				sumList = append(sumList, sumLeft-node.Right.Val)
				tmp2 = append(tmp2, node.Right.Val)
				pathList = append(pathList, tmp2)
			}

		}
	}

	return result
}

func main() {
	root := &model.TreeNode{
		Val: 5,
	}
	tree := model.InitBinaryTree(root)
	fmt.Println(pathSum(tree, 22))
	fmt.Println(pathSum2(tree, 22))
}
