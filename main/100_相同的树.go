package main

import (
	"fmt"
	"main/model"
	"time"
)

/*
* 给定两个二叉树，编写一个函数来检验它们是否相同。
 *
 * 如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。
 *
 * 示例 1:
 *
 * 输入:       1         1
 * ⁠         / \       / \
 * ⁠        2   3     2   3
 *
 * ⁠       [1,2,3],   [1,2,3]
 *
 * 输出: true
 *
 * 示例 2:
 *
 * 输入:      1          1
 * ⁠         /           \
 * ⁠        2             2
 *
 * ⁠       [1,2],     [1,null,2]
 *
 * 输出: false
 *
 *
 * 示例 3:
 *
 * 输入:       1         1
 * ⁠         / \       / \
 * ⁠        2   1     1   2
 *
 * ⁠       [1,2,1],   [1,1,2]
 *
 * 输出: false
*/

//O(n)
func isSameTree(p *model.TreeNode, q *model.TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false
	} else if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

/**
一开始我没有加【*ret = append(*ret, 0)】这个，导致一直调试不通过
后台想想，两个树不同但可能某序遍历相同，但遍历加上根节点那必然不同咯
仔细想想，这不就是把leetcode输入序列还原嘛,null就是我追加的0，既根节点
这个递归居然执行用时100%，估计很多人都用递归吧
代码
*/

//O(n)
func isSameTree2(p *model.TreeNode, q *model.TreeNode) bool {
	var tp, tq []int
	f(p, &tp)
	f(q, &tq)
	if len(tp) != len(tq) {
		return false
	}
	for i, v := range tp {
		if tq[i] != v {
			return false
		}
	}
	return true
}

func f(root *model.TreeNode, ret *[]int) {
	if root == nil {
		*ret = append(*ret, 0)
		return
	}
	*ret = append(*ret, root.Val)
	f(root.Left, ret)
	f(root.Right, ret)
}

func main() {
	tree1 := &model.TreeNode{Val: 1}
	tree1.Right = &model.TreeNode{Val: 2}
	tree1.Left = &model.TreeNode{Val: 1}

	tree2 := &model.TreeNode{Val: 1}
	tree2.Right = &model.TreeNode{Val: 2}
	tree2.Left = &model.TreeNode{Val: 1}

	s1 := time.Now()
	fmt.Println(isSameTree(tree1, tree2))
	fmt.Println(time.Since(s1).String())

	s2 := time.Now()
	fmt.Println(isSameTree2(tree1, tree2))
	fmt.Println(time.Since(s2).String())
}
