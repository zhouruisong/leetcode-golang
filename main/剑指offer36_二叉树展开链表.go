package main

import (
	"fmt"
	"main/model"
)

/*
二叉树展开为链表
输入一棵二叉搜索树，将该二叉搜索树转换成一个排序的循环双向链表。要求不能创建任何新的节点，只能调整树中节点指针的指向。
*/

var pre *model.TreeNode
var head *model.TreeNode

//O(N)
func treeToDoublyList(root *model.TreeNode) *model.TreeNode {
	if root == nil {
		return nil
	}

	dfs(root)
	//如何要求双链表首尾连接,下面两行代码注释掉
	head.Left = pre
	pre.Right = head
	return head
}

//循环双链表
func dfs(cur *model.TreeNode) {
	if cur == nil {
		return
	}

	dfs(cur.Left)

	if pre != nil {
		pre.Right = cur
	} else {
		head = cur
	}
	//单链表的话,注释这一行即可
	cur.Left = pre

	pre = cur

	dfs(cur.Right)
}

/*
       5
      / \
     3   6
    / \
   2   4
  /
 1
*/

func InOrder(current *model.TreeNode) {
	if current != nil {
		InOrder(current.Left)
		fmt.Printf(" %v", current.Val)
		InOrder(current.Right)
	}
}

func printDoubleLink(current *model.TreeNode, n int) {
	for current != nil && n >= 0 {
		fmt.Print(current.Val, "->")
		current = current.Right
		n--
	}
}

//非循环双链表
func Convert(pRootOfTree *model.TreeNode) *model.TreeNode {
	// write code here
	var pre *model.TreeNode
	var head *model.TreeNode
	var f func(cur *model.TreeNode)
	f = func(cur *model.TreeNode) {
		if cur == nil {
			return
		}

		f(cur.Left)

		if pre != nil {
			pre.Right = cur
		} else {
			head = cur
		}
		cur.Left = pre
		pre = cur

		f(cur.Right)
	}

	f(pRootOfTree)
	//如何要求双链表首尾连接,下面两行代码注释掉
	//head.Left = pre
	//pre.Right = head
	return head
}

func main() {
	root := &model.TreeNode{
		Val: 5,
	}
	root.Left = &model.TreeNode{Val: 3}
	root.Right = &model.TreeNode{Val: 6}

	root.Left.Left = &model.TreeNode{Val: 2}
	root.Left.Right = &model.TreeNode{Val: 4}

	root.Left.Left.Left = &model.TreeNode{Val: 1}

	//InOrder(root)

	//head2 := treeToDoublyList(root)
	head2 := Convert(root)
	printDoubleLink(head2, 7)
}
