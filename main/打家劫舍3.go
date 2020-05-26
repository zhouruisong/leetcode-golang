package main

import (
	"fmt"
	"main/model"
	"math"
)

/*
* 在上次打劫完一条街道之后和一圈房屋后，小偷又发现了一个新的可行窃的地区。这个地区只有一个入口，我们称之为“根”。
 * 除了“根”之外，每栋房子有且只有一个“父“房子与之相连。一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。
 * 如果两个直接相连的房子在同一天晚上被打劫，房屋将自动报警。
 *
 * 计算在不触动警报的情况下，小偷一晚能够盗取的最高金额。
 *
 * 示例 1:
 *
 * 输入: [3,2,3,null,3,null,1]
 *
 * ⁠    3
 * ⁠   / \
 * ⁠  2   3
 * ⁠   \   \
 * ⁠    3   1
 *
 * 输出: 7
 * 解释: 小偷一晚能够盗取的最高金额 = 3 + 3 + 1 = 7.
 *
 * 示例 2:
 *
 * 输入: [3,4,5,1,3,null,1]
 *
 *     3
 * ⁠   / \
 * ⁠  4   5
 * ⁠ / \   \
 * ⁠1   3   1
 *
 * 输出: 9
 * 解释: 小偷一晚能够盗取的最高金额 = 4 + 5 = 9.
 *
 *
*/

/*
在解法一和解法二中，我们使用爷爷、两个孩子、4 个孙子来说明问题
首先来定义这个问题的状态
爷爷节点获取到最大的偷取的钱数呢

首先要明确相邻的节点不能偷，也就是爷爷选择偷，儿子就不能偷了，但是孙子可以偷
二叉树只有左右两个孩子，一个爷爷最多 2 个儿子，4 个孙子
根据以上条件，我们可以得出单个节点的钱该怎么算
4 个孙子偷的钱 + 爷爷的钱 VS 两个儿子偷的钱 哪个组合钱多，就当做当前节点能偷的最大钱数。这就是动态规划里面的最优子结构

由于是二叉树，这里可以选择计算所有子节点

4 个孙子投的钱加上爷爷的钱如下
int method1 = root.val + rob(root.left.left) + rob(root.left.right) + rob(root.right.left) + rob(root.right.right)
两个儿子偷的钱如下
int method2 = rob(root.left) + rob(root.right);
挑选一个钱数多的方案则
int result = Math.max(method1, method2);
*/

func rob(root *model.TreeNode) int {
	if root == nil {
		return 0
	}

	money := root.Val
	if root.Left != nil {
		money += (rob(root.Left.Left) + rob(root.Left.Right))
	}

	if root.Right != nil {
		money += (rob(root.Right.Left) + rob(root.Right.Right))
	}

	return int(math.Max(float64(money), float64(rob(root.Left)+rob(root.Right))))
}

func rob2(root *model.TreeNode) int {
	if root == nil {
		return 0
	}

	//当前root选： root.Val + (root.Left [不选]结果) + （root.Right【不选】结果）
	//当前root不选： (root.Left [选]结果) + （root.Right【选】结果）
	//返回两者的最大值
	selectVal, unSelectVal := nodeSelect(root)
	return Max(selectVal, unSelectVal)
}

func nodeSelect(root *model.TreeNode) (selectVal, unSelectVal int) {
	if root == nil {
		return 0, 0
	}

	//sLeftVal, unLeftVal当前root的左孩子选和不选的结果
	sLeftVal, unLeftVal := nodeSelect(root.Left)

	//sRightVal, unRightVal当前root的右孩子选和不选的结果
	sRightVal, unRightVal := nodeSelect(root.Right)

	//selectVal 当前节点[选】；当前节点能偷到的最大钱数=左孩子选择自己不偷时能得到的钱 +  右孩子选择不偷时能得到的钱 + 当前节点的钱数
	//unSelectVal 当前节点[不选】；当前节点能偷到的最大钱数 = 左孩子能偷到的钱 + 右孩子能偷到的钱
	selectVal = root.Val + unLeftVal + unRightVal
	unSelectVal = Max(sLeftVal, unLeftVal) + Max(sRightVal, unRightVal)
	return
}

func Max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	x := &model.TreeNode{Val: 3}
	x.Left = &model.TreeNode{Val: 4}
	x.Right = &model.TreeNode{Val: 5}
	x.Left.Left = &model.TreeNode{Val: 1}
	x.Left.Right = &model.TreeNode{Val: 3}
	x.Right.Right = &model.TreeNode{Val: 1}

	fmt.Println(rob2(x))
}
