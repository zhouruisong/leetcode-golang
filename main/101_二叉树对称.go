package main

import (
	"fmt"
	"main/model"
)

/*
给定一个二叉树，检查它是否是镜像对称的。

例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

    1
   / \
  2   2
 / \ / \
3  4 4  3
但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:

    1
   / \
  2   2
   \   \
   3    3
说明:

如果你可以运用递归和迭代两种方法解决这个问题，会很加分。
*/

func isSymmetric(root *model.TreeNode) bool {
	/*
		除了递归的方法外，我们也可以利用队列进行迭代。队列中每两个连续的结点应该是相等的，而且它们的子树互为镜像。
		最初，队列中包含的是 root 以及 root。该算法的工作原理类似于 BFS，但存在一些关键差异。每次提取两个结点并比较它们的值。
		然后，将两个结点的左右子结点按相反的顺序插入队列中。当队列为空时，或者我们检测到树不对称（即从队列中取出两个不相等的连续结点）时，
		该算法结束。
	*/
	if root == nil {
		return true
	}

	//两个root在里面,采用栈
	node := []*model.TreeNode{root, root}

	for len(node) != 0 {
		node1 := node[len(node)-1]
		node = node[:len(node)-1]
		node2 := node[len(node)-1]
		node = node[:len(node)-1]

		if node1 == nil && node2 == nil {
			continue
		}

		if node1 == nil || node2 == nil {
			return false
		}

		if node1.Val != node2.Val {
			return false
		}

		node = append(node, node1.Left)
		node = append(node, node2.Right)
		node = append(node, node1.Right)
		node = append(node, node2.Left)
	}

	return true
}

//递归
func isSymmetric2(root *model.TreeNode) bool {
	return mirrorBinary(root, root)

}

func mirrorBinary(t1 *model.TreeNode, t2 *model.TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}

	if t1 == nil || t2 == nil {
		return false
	}

	return (t1.Val == t2.Val) &&
		mirrorBinary(t1.Left, t2.Right) &&
		mirrorBinary(t1.Right, t2.Left)
}

func main() {
	root := &model.TreeNode{
		Val: 1,
	}
	tree := model.InitBinaryTree3(root)
	fmt.Println(isSymmetric(tree), isSymmetric2(tree))
}
