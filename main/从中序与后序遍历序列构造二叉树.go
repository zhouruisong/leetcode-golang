package main

import (
	"fmt"
	"main/model"
)

/*
 * 根据一棵树的中序遍历与后序遍历构造二叉树。
 *
 * 注意:
 * 你可以假设树中没有重复的元素。
 *
 * 例如，给出
 *
 * 中序遍历 inorder = [9,3,15,20,7]
 * 后序遍历 postorder=[9,15,7,20,3]
 *
 * 返回如下的二叉树：
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 *
 */

func buildTree(inorder []int, postorder []int) *model.TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}

	root := 0
	for k, v := range inorder {
		if v == postorder[len(postorder)-1] {
			root = k
			break
		}
	}

	n := len(postorder)
	return &model.TreeNode{
		Val:   inorder[root],
		Left:  buildTree(inorder[:root], postorder[:root]),
		Right: buildTree(inorder[root+1:], postorder[root:n-1]),
	}
}

//中序遍历
func InOrder(current *model.TreeNode) {
	if current != nil {
		InOrder(current.Left)
		fmt.Printf(" %v", current.Val)
		InOrder(current.Right)
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

func main() {
	root := buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3})
	InOrder(root)
	fmt.Println("\n========")
	PostOrder(root)
}
