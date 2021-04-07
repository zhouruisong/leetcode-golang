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

//单链表,使用中序遍历
func treeToList2(root *model.TreeNode) {
	// write code here
	head := new(model.TreeNode)
	pre = head

	var f func(cur *model.TreeNode)

	f = func(cur *model.TreeNode) {
		if cur == nil {
			return
		}

		//清除left、right，构建单链表形式
		left, right := cur.Left, cur.Right
		cur.Left, cur.Right = nil, nil

		pre.Right = cur
		pre = pre.Right

		f(left)
		f(right)
	}

	f(root)

	return
}

//循环双链表,使用中序遍历
func treeToDoublyList2(root *model.TreeNode) {
	// write code here
	head := new(model.TreeNode)
	pre = head

	var f func(cur *model.TreeNode)

	f = func(cur *model.TreeNode) {
		if cur == nil {
			return
		}

		//清除left、right，构建单链表形式
		left, right := cur.Left, cur.Right
		cur.Left, cur.Right = nil, nil

		pre.Right = cur
		cur.Left = pre
		pre = pre.Right

		f(left)
		f(right)
	}

	f(root)
	head.Left = pre
	pre.Right = head
	return
}

//使用水平遍历
func treeToList(root *model.TreeNode) {
	if root == nil {
		return
	}

	stack := []*model.TreeNode{root}
	var pre *model.TreeNode
	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if pre != nil {
			pre.Left, pre.Right = nil, curr
		}
		//前序遍历,右边节点先入栈,后出栈
		if curr.Right != nil {
			stack = append(stack, curr.Right)
		}
		if curr.Left != nil {
			stack = append(stack, curr.Left)
		}
		pre = curr
	}
}

/*
       1              1
      / \              \
     2   5       ->     2
    / \   \              \
   3   4   6              3
                           \
                            4
                             \
                              5
                               \
                                6

输入：root = [1,2,5,3,4,null,6]
输出：[1,null,2,null,3,null,4,null,5,null,6]
*/
func main() {
	root := &model.TreeNode{
		Val: 1,
	}
	root.Left = &model.TreeNode{Val: 2}
	root.Left.Left = &model.TreeNode{Val: 3}
	root.Left.Right = &model.TreeNode{Val: 4}

	root.Right = &model.TreeNode{Val: 5}
	root.Right.Right = &model.TreeNode{Val: 6}

	fmt.Println("循环双链表")
	InOrder(root)
	fmt.Println("")
	//head := treeToDoublyList(root) //循环双链表
	treeToDoublyList2(root)
	//Convert(root)
	printDoubleLink(root, 7)

	root2 := &model.TreeNode{
		Val: 1,
	}
	root2.Left = &model.TreeNode{Val: 2}
	root2.Left.Left = &model.TreeNode{Val: 3}
	root2.Left.Right = &model.TreeNode{Val: 4}

	root2.Right = &model.TreeNode{Val: 5}
	root2.Right.Right = &model.TreeNode{Val: 6}

	fmt.Println("\n单链表1")
	InOrder(root2)
	fmt.Println("")
	treeToList(root2) //单链表
	InOrder(root2)

	root3 := &model.TreeNode{
		Val: 1,
	}
	root3.Left = &model.TreeNode{Val: 2}
	root3.Left.Left = &model.TreeNode{Val: 3}
	root3.Left.Right = &model.TreeNode{Val: 4}

	root3.Right = &model.TreeNode{Val: 5}
	root3.Right.Right = &model.TreeNode{Val: 6}

	fmt.Println("\n单链表2")
	InOrder(root3)
	fmt.Println("")
	treeToList2(root3) //单链表2
	InOrder(root3)
}
