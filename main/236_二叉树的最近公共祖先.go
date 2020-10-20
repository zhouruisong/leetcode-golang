package main

import "main/model"

/*
给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

例如，给定如下二叉树:  root = [3,5,1,6,2,0,8,null,null,7,4]

      3
     /  \
    5    1
   / \  / \
  6   2 0   8
     / \
    7   4

示例 1:

输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
输出: 3
解释: 节点 5 和节点 1 的最近公共祖先是节点 3。
示例 2:

输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
输出: 5
解释: 节点 5 和节点 4 的最近公共祖先是节点 5。因为根据定义最近公共祖先节点可以为节点本身。

https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/solution/236-er-cha-shu-de-zui-jin-gong-gong-zu-xian-hou-xu/
*/

func lowestCommonAncestor(root, p, q *model.TreeNode) *model.TreeNode {
	type void struct{}
	var emptyValue void
	fatherMap := make(map[int]*model.TreeNode)
	//用map实现set，map的值为struct{}则不额外占用空间
	visit := make(map[int]void)
	var dfs func(*model.TreeNode)
	// 构造子节点指向祖先节点的指针
	dfs = func(root *model.TreeNode) {
		if root.Left != nil {
			fatherMap[root.Left.Val] = root
			dfs(root.Left)
		}
		if root.Right != nil {
			fatherMap[root.Right.Val] = root
			dfs(root.Right)
		}
	}
	dfs(root)
	// 从p节点开始不断往它的祖先移动，并用数据结构记录已经访问过的祖先节点。
	for p != nil {
		visit[p.Val] = emptyValue
		p = fatherMap[p.Val]
	}
	// 从q节点开始不断往它的祖先移动，如果有祖先已经被访问过，即意味着这是 p 和 q 的深度最深的公共祖先，即 LCA 节点。
	for q != nil {
		if _, exist := visit[q.Val]; exist {
			return q
		}
		q = fatherMap[q.Val]
	}
	return nil
}

func lowestCommonAncestor2(root, p, q *model.TreeNode) *model.TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Right, p, q)
	if l == nil {
		return r
	}
	if r == nil {
		return l
	}
	return root
}
