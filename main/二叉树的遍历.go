package main

import (
	"fmt"
	"main/model"
)

//中序遍历
func InOrder(current *model.TreeNode) {
	if current != nil {
		InOrder(current.Left)
		fmt.Printf(" %v", current.Val)
		InOrder(current.Right)
	}
}

//先序遍历
func PreOrder(current *model.TreeNode) {
	if current != nil {
		fmt.Printf(" %v", current.Val)
		PreOrder(current.Left)
		PreOrder(current.Right)
	}
}

//后序遍历
func PostOrder(current *model.TreeNode) {
	if current != nil {
		PostOrder(current.Left)
		PostOrder(current.Right)
		fmt.Printf(" %v", current.Val)
	}
}
