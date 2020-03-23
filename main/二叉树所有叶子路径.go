package main

import (
	"container/list"
	"fmt"
	"main/model"
)

/*
给定一个二叉树，返回所有从根节点到叶子节点的路径。

说明: 叶子节点是指没有子节点的节点。

示例:

输入:

   1
 /   \
2     3
 \
  5

输出: ["1->2->5", "1->3"]

解释: 所有根节点到叶子节点的路径为: 1->2->5, 1->3
*/

func current_path(current *model.TreeNode, path string, paths *[]string) {
	if current != nil {
		path += fmt.Sprintf("%v", current.Val)
		if current.Right == nil && current.Left == nil {
			*paths = append(*paths, path)
		} else {
			path += "->"
			current_path(current.Left, path, paths)
			current_path(current.Right, path, paths)
		}
	}
}

func binaryTreePaths2(root *model.TreeNode) []string {
	var paths []string
	current_path(root, "", &paths)
	return paths
}

func binaryTreePaths(root *model.TreeNode) []string {
	var paths []string
	if root == nil {
		return paths
	}

	nodeLiist := list.New()
	pathList := list.New()

	nodeLiist.PushBack(root)
	pathList.PushBack(fmt.Sprintf("%v", root.Val))

	for nodeLiist.Len() != 0 {
		node := nodeLiist.Remove(nodeLiist.Front()).(*model.TreeNode)
		path := pathList.Remove(pathList.Front()).(string)
		if node.Right == nil && node.Left == nil {
			paths = append(paths, path)
		}

		if node.Left != nil {
			nodeLiist.PushBack(node.Left)
			pathList.PushBack(fmt.Sprintf("%v->%v", path, node.Left.Val))
		}
		if node.Right != nil {
			nodeLiist.PushBack(node.Right)
			pathList.PushBack(fmt.Sprintf("%v->%v", path, node.Right.Val))
		}
	}

	/*
			          5
		             / \
		            4   8
		           /   / \
		          11  13  4
		         /  \      \
		        7    2      1
	*/

	return paths
}

func main() {
	root := &model.TreeNode{
		Val: 5,
	}
	tree := model.InitBinaryTree(root)
	//PreOrder(tree)
	fmt.Println(binaryTreePaths2(tree))
}
