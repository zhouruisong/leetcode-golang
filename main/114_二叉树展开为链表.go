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

	for i := 0; i < len(nodeSlice); i++ {
		fmt.Print(nodeSlice[i].Val, "->")
	}

	fmt.Println("=========")
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

//先序遍历递归
func PreOrder(current *model.TreeNode) {
	if current != nil {
		fmt.Printf(" %v", current.Val)
		PreOrder(current.Left)
		PreOrder(current.Right)
	}
}

func main() {
	root := &model.TreeNode{
		Val: 1,
	}
	tree := model.InitBinaryTree4(root)
	flatten(tree)
	PreOrder(tree)
}
