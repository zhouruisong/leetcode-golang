package main

import (
	"container/list"
	"main/model"
)

/*
给定一个二叉树，编写一个函数来获取这个树的最大宽度。树的宽度是所有层中的最大宽度。这个二叉树与满二叉树（full binary tree）结构相同，但一些节点为空。

每一层的宽度被定义为两个端点（该层最左和最右的非空节点，两端点间的null节点也计入长度）之间的长度。

示例 1:

输入:

           1
         /   \
        3     2
       / \     \
      5   3     9

输出: 4
解释: 最大值出现在树的第 3 层，宽度为 4 (5,3,null,9)。
示例 2:

输入:

          1
         /
        3
       / \
      5   3

输出: 2
解释: 最大值出现在树的第 3 层，宽度为 2 (5,3)。
示例 3:

输入:

          1
         / \
        3   2
       /
      5

输出: 2
解释: 最大值出现在树的第 2 层，宽度为 2 (3,2)。
示例 4:

输入:

          1
         / \
        3   2
       /     \
      5       9
     /         \
    6           7
输出: 8
解释: 最大值出现在树的第 4 层，宽度为 8 (6,null,null,null,null,null,null,7)。

*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func widthOfBinaryTree(root *model.TreeNode) int {
	if root == nil {
		return 0
	}

	nodes := list.New()
	nodes.PushFront(root)
	max := 1

	//两个队列
	nodeIndexList := list.New()
	nodeIndexList.PushFront(1)

	for nodes.Len() > 0 {
		currentLen := nodes.Len()
		for i := 0; i < currentLen; i++ {
			node := nodes.Remove(nodes.Back()).(*model.TreeNode)
			nodeIndex := nodeIndexList.Remove(nodeIndexList.Back()).(int)
			if node.Left != nil {
				nodes.PushFront(node.Left)
				nodeIndexList.PushFront(2 * nodeIndex)
			}

			if node.Right != nil {
				nodes.PushFront(node.Right)
				nodeIndexList.PushFront(2*nodeIndex + 1)
			}
		}

		if nodeIndexList.Len() >= 2 {
			last := nodeIndexList.Front().Value.(int)
			first := nodeIndexList.Back().Value.(int)
			if max < last-first+1 {
				max = last - first + 1
			}
		}
	}

	return max
}
