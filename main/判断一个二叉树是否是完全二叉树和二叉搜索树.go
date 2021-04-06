package main

import (
	"container/list"
	"fmt"
	"main/model"
)

func judgeIt(root *model.TreeNode) []bool {
	// write code here
	//判断一个二叉树是否是搜索树，中序遍历后，是升序就可以
	res := []int{}
	result := []bool{true, true}
	dfs(root, &res)
	//cur := res[0]
	for i := 1; i < len(res); i++ {
		if res[i-1] > res[i] {
			result[0] = false
			break
		}
	}

	result[1] = isCompleteTree(root)
	return result
}

func dfs(root *model.TreeNode, res *[]int) {
	if root != nil {
		dfs(root.Left, res)
		*res = append(*res, root.Val)
		dfs(root.Right, res)
		return
	}
}

func isCompleteTree(root *model.TreeNode) bool {
	//层次遍历直至遇到第一个空节点
	//完全二叉树在遇到空节点之后剩余的应当全是空节点
	qu := list.New()
	qu.PushBack(root)
	//top := qu.Front().Value.(*TreeNode)
	for qu.Front().Value.(*model.TreeNode) != nil {
		node := qu.Remove(qu.Front()).(*model.TreeNode)
		qu.PushBack(node.Left)
		qu.PushBack(node.Right)
		//top = qu.Front().Value.(*TreeNode)
	}
	//top = qu.Front().Value.(*TreeNode)
	for qu.Len() != 0 && qu.Front().Value.(*model.TreeNode) == nil {
		qu.Remove(qu.Front())
		//top = qu.Front().Value.(*TreeNode)
	}

	return qu.Len() == 0
}

/*
 * 输入:
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   3
 * ⁠/ \  /
 * 4  5 6
 */

func main() {
	root := &model.TreeNode{
		Val: 1,
	}

	root.Left = &model.TreeNode{
		Val: 2,
	}
	root.Left.Left = &model.TreeNode{
		Val: 4,
	}
	root.Left.Right = &model.TreeNode{
		Val: 5,
	}

	root.Right = &model.TreeNode{
		Val: 3,
	}
	root.Right.Left = &model.TreeNode{
		Val: 6,
	}

	fmt.Println(judgeIt(root))
}
