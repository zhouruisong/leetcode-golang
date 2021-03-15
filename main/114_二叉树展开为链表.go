package main

import (
	"container/list"
	"fmt"
	"main/model"
)

/*
给定一个二叉树，原地将它展开为链表。

例如，给定二叉树

    1
   / \
  2   5
 / \   \
3   4   6
将其展开为：

1
 \
  2
   \
    3
     \
      4
       \
        5
         \
          6

https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list/

*/
func flatten(root *model.TreeNode) {
	if root == nil {
		return
	}

	//二叉树先序遍历
	nodeSlice := []*model.TreeNode{}
	nodeList := list.New()
	current := root
	for current != nil || nodeList.Len() != 0 {
		for current != nil {
			nodeSlice = append(nodeSlice, current)
			nodeList.PushFront(current)
			current = current.Left
		}

		if nodeList.Len() != 0 {
			current = nodeList.Remove(nodeList.Front()).(*model.TreeNode)
			current = current.Right
		}
	}

	//for i := 0; i < len(nodeSlice); i++ {
	//	fmt.Print(nodeSlice[i].Val, "->")
	//}
	//
	//fmt.Println("=========")
	//节点全部在右子树
	root = nodeSlice[0]
	head := nodeSlice[0]
	for i := 0; i < len(nodeSlice)-1; i++ {
		nodeSlice[i].Right = nil
		nodeSlice[i].Left = nil
		head.Right = nodeSlice[i+1]
		head = head.Right
	}
}

//前序遍历和展开分开进行
func flatten2(root *model.TreeNode) {
	list := []*model.TreeNode{}
	stack := []*model.TreeNode{}
	node := root

	for node != nil || len(stack) > 0 {
		for node != nil {
			list = append(list, node)
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		node = node.Right
		stack = stack[:len(stack)-1]
	}

	for i := 1; i < len(list); i++ {
		prev, curr := list[i-1], list[i]
		prev.Left, prev.Right = nil, curr
	}
}

//前序遍历和展开同步进行
func flatten3(root *model.TreeNode) {
	if root == nil {
		return
	}
	stack := []*model.TreeNode{root}
	var prev *model.TreeNode
	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if prev != nil {
			prev.Left, prev.Right = nil, curr
		}
		left, right := curr.Left, curr.Right
		//前序遍历,右节点先入栈
		if right != nil {
			stack = append(stack, right)
		}
		if left != nil {
			stack = append(stack, left)
		}
		prev = curr
	}
}

//先序遍历递归
func PreOrder(current *model.TreeNode) {
	if current != nil {
		fmt.Printf("%v->", current.Val)
		PreOrder(current.Left)
		PreOrder(current.Right)
	}
}

func main() {
	root := &model.TreeNode{
		Val: 1,
	}
	tree := model.InitBinaryTree4(root)
	flatten3(tree)
	PreOrder(tree)
}
