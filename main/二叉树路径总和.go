package main

import (
	"container/list"
	"fmt"
	"main/model"
)

/*
给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和。

说明: 叶子节点是指没有子节点的节点。

示例:
给定如下二叉树，以及目标和 sum = 22，

              5
             / \
            4   8
           /   / \
          11  13  4
         /  \      \
        7    2      1
返回 true, 因为存在目标和为 22 的根节点到叶子节点的路径 5->4->11->2。
*/

func hasPathSum(root *model.TreeNode, sum int) bool {
	//递归方法===遍历每个节点，如果该节点是叶子，判断sum是否是0，非0返回false，否则返回true
	if root == nil {
		return false
	}

	sum -= root.Val

	if root.Left == nil && root.Right == nil {
		return sum == 0
	}

	return hasPathSum(root.Left, sum) || hasPathSum(root.Right, sum)
}

/*
我们可以用栈将递归转成迭代的形式。深度优先搜索在除了最坏情况下都比广度优先搜索更快。最坏情况是指满足目标和的 root->leaf 路径是最后被考虑的，这种情况下深度优先搜索和广度优先搜索代价是相通的。
利用深度优先策略访问每个节点，同时更新剩余的目标和。
所以我们从包含根节点的栈开始模拟，剩余目标和为 sum - root.val。
然后开始迭代：弹出当前元素，如果当前剩余目标和为 0 并且在叶子节点上返回 True；如果剩余和不为零并且还处在非叶子节点上，将当前节点的所有孩子以及对应的剩余和压入栈中。
*/
func hasPathSum2(root *model.TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	nodeStack := list.New()
	sumStack := list.New()
	pathStack := list.New()

	nodeStack.PushFront(root)
	sumStack.PushFront(sum - root.Val)
	pathStack.PushFront(fmt.Sprintf("%v", root.Val))

	for nodeStack.Len() > 0 {
		currentLen := nodeStack.Len()
		for i := 0; i < currentLen; i++ {
			node := nodeStack.Remove(nodeStack.Back()).(*model.TreeNode)
			val := sumStack.Remove(sumStack.Back()).(int)
			path := pathStack.Remove(pathStack.Back()).(string)

			if node.Left == nil && node.Right == nil && val == 0 {
				fmt.Printf("%v\n", path)
				return true
			}

			if node.Left != nil {
				nodeStack.PushFront(node.Left)
				sumStack.PushFront(val - node.Left.Val)
				pathStack.PushFront(fmt.Sprintf("%v->%v", path, node.Left.Val))
			}

			if node.Right != nil {
				nodeStack.PushFront(node.Right)
				sumStack.PushFront(val - node.Right.Val)
				pathStack.PushFront(fmt.Sprintf("%v->%v", path, node.Right.Val))
			}
		}
	}

	return false
}

func main() {
	root := &model.TreeNode{
		Val: 5,
	}
	tree := model.InitBinaryTree(root)
	fmt.Println(hasPathSum2(tree, 18))
}
