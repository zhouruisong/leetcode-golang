package main

import (
	"container/list"
	"fmt"
	"main/model"
)

//中序遍历递归
func InOrder(current *model.TreeNode) {
	if current != nil {
		InOrder(current.Left)
		fmt.Printf(" %v", current.Val)
		InOrder(current.Right)
	}
}

//中序遍历非递归(借助栈)
func InOrder2(current *model.TreeNode) {
	//打印节点， 将左子树全部入队，一个一个出队，再将右子树入队
	if current == nil {
		return
	}

	var result []int
	nodelist := list.New()

	for current != nil || nodelist.Len() != 0 {
		for current != nil {
			nodelist.PushFront(current)
			current = current.Left
		}

		if nodelist.Len() != 0 {
			current = nodelist.Remove(nodelist.Front()).(*model.TreeNode)
			result = append(result, current.Val)
			current = current.Right
		}
	}

	fmt.Println(result)
}

func inorderTraversal(root *model.TreeNode) []int {
	//打印节点， 将左子树全部入队，一个一个出队，再将右子树入队
	if root == nil {
		return []int{}
	}

	var result []int
	nodelist := list.New()

	for root != nil || nodelist.Len() != 0 {
		for root != nil {
			nodelist.PushFront(root)
			root = root.Left
		}

		if nodelist.Len() != 0 {
			root = nodelist.Remove(nodelist.Front()).(*model.TreeNode)
			result = append(result, root.Val)
			root = root.Right
		}
	}

	return result
}

//先序遍历递归
func PreOrder(current *model.TreeNode) {
	if current != nil {
		fmt.Printf(" %v", current.Val)
		PreOrder(current.Left)
		PreOrder(current.Right)
	}
}

//先序遍历非递归(借助栈)
func PreOrder2(current *model.TreeNode) {
	//打印节点， 将左子树全部入队，一个一个出队，再将右子树入队
	if current == nil {
		return
	}

	var result []int
	nodelist := list.New()

	for current != nil || nodelist.Len() != 0 {
		for current != nil {
			result = append(result, current.Val)
			nodelist.PushFront(current)
			current = current.Left
		}

		if nodelist.Len() != 0 {
			current = nodelist.Remove(nodelist.Front()).(*model.TreeNode)
			current = current.Right
		}
	}

	fmt.Println(result)
}

//后序遍历递归
func PostOrder(current *model.TreeNode) {
	if current != nil {
		PostOrder(current.Left)
		PostOrder(current.Right)
		fmt.Printf(" %v", current.Val)
	}
}

//后序遍历非递归(借助栈)
func PostOrder3(root *model.TreeNode) []int {
	var res []int
	var stack = []*model.TreeNode{root}
	for len(stack) > 0 {
		if root != nil {
			//根右左
			res = append(res, root.Val)
			stack = append(stack, root.Left)  //左节点，因为先进 所以后出
			stack = append(stack, root.Right) //右节点，因为后进 所以先出
		}
		index := len(stack) - 1 //栈顶
		root = stack[index]     //出栈
		stack = stack[:index]
	}

	//反转 变成后序遍历 左右根
	l, r := 0, len(res)-1
	for l < r {
		res[l], res[r] = res[r], res[l]
		l++
		r--
	}
	return res
}

/* 中序遍历 inorder = [9,3,15,20,7]
*  后序遍历 postorder=[9,15,7,20,3]
*
* 返回如下的二叉树：
*
* ⁠   3
* ⁠  / \
* ⁠ 9  20
* ⁠   /  \
* ⁠  15   7
 */

func main() {
	tree := &model.TreeNode{
		Val: 3,
	}
	model.InitBinaryTree3(tree)

	InOrder(tree)
	fmt.Println("\n=====")
	InOrder2(tree)
	fmt.Println("\n=====")

	fmt.Println(inorderTraversal(tree))
	fmt.Println("\n=====")

	//PreOrder(tree)
	//fmt.Println("\n=====")
	//PreOrder2(tree)
	//
	//fmt.Println("\n=====")
	//res := PostOrder3(tree)
	//fmt.Println(res)
	//PostOrder(tree)
	//fmt.Println("\n===111==")
	//PostOrder2(tree)

}
