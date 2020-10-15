package main

import (
	"fmt"
	"main/model"
	"time"
)

/*
给定一个整数 n，生成所有由 1 ... n 为节点所组成的二叉搜索树。
示例:
输入: 3
输出:
[
  [1,null,3,2],
  [3,2,null,1],
  [3,1,null,null,2],
  [2,1,3],
  [1,null,2,null,3]
]
解释:
以上的输出对应以下 5 种不同结构的二叉搜索树：

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3
*/

/*

 */
func generateTrees(n int) []*model.TreeNode {
	if n == 0 {
		return []*model.TreeNode{}
	}

	return dfs(1, n)
}

//递归
func dfs(start, end int) []*model.TreeNode {
	if start > end {
		return []*model.TreeNode{nil} //返回 []{nil}, len([]{nil}) = 1
	}
	tmp := []*model.TreeNode{}
	if start == end {
		return append(tmp, &model.TreeNode{Val: start})
	}

	for i := start; i <= end; i++ {
		left := dfs(start, i-1)
		right := dfs(i+1, end)

		for j := 0; j < len(left); j++ {
			for k := 0; k < len(right); k++ {
				tree := &model.TreeNode{
					Left:  left[j],
					Right: right[k],
					Val:   i}
				tmp = append(tmp, tree)
				//PreOrder(tree)
			}
		}
	}

	return tmp
}

//动态规划
/*
考虑 [] 的所有解
null

考虑 [ 1 ] 的所有解
1

考虑 [ 1 2 ] 的所有解
  2
 /
1

 1
  \
   2

考虑 [ 1 2 3 ] 的所有解
    3
   /
  2
 /
1

   2
  / \
 1   3

     3
    /
   1
    \
     2

   1
     \
      3
     /
    2

  1
    \
     2
      \
       3
仔细分析，可以发现一个规律。首先我们每次新增加的数字大于之前的所有数字，所以新增加的数字出现的位置只可能是根节点或者是根节点的右孩子，
右孩子的右孩子，右孩子的右孩子的右孩子等等，总之一定是右边。其次，新数字所在位置原来的子树，改为当前插入数字的左孩子即可，因为插入数字是最大的。

对于下边的解
  2
 /
1

然后增加 3
1.把 3 放到根节点
    3
   /
  2
 /
1

2. 把 3 放到根节点的右孩子
   2
  / \
 1   3

对于下边的解
 1
  \
   2

然后增加 3
1.把 3 放到根节点
     3
    /
   1
    \
     2

2. 把 3 放到根节点的右孩子，原来的子树作为 3 的左孩子
      1
        \
         3
        /
      2

3. 把 3 放到根节点的右孩子的右孩子
  1
    \
     2
      \
       3
以上就是根据 [ 1 2 ] 推出 [ 1 2 3 ] 的所有过程，可以写代码了。由于求当前的所有解只需要上一次的解，
所有我们只需要两个 list，pre 保存上一次的所有解， cur 计算当前的所有解。
*/
func generateTrees2(n int) []*model.TreeNode {
	if n == 0 {
		return []*model.TreeNode{}
	}

	pre := []*model.TreeNode{nil}
	//每次增加一个数字
	for i := 1; i <= n; i++ {
		cur := []*model.TreeNode{}
		//遍历之前的所有解
		ln := len(pre)
		for m := 0; m < ln; m++ {
			root := pre[m]
			//将新增加的数字放到根节点
			insert := &model.TreeNode{Val: i, Left: root}
			cur = append(cur, insert)

			//插入到右孩子，右孩子的右孩子...最多找 n 次孩子
			for j := 0; j < n; j++ {
				treeCopy := treeCopy(root)
				right := treeCopy
				//找到要插入的位置
				for k := 0; k < j; k++ {
					if right == nil {
						break
					}
					right = right.Right
				}

				//到达 null 提前结束
				if right == nil {
					break
				}

				//保存当前右孩子的位置的子树作为插入节点的左孩子
				rightTree := right.Right
				insert2 := &model.TreeNode{Val: i}
				right.Right = insert2    //右孩子是插入的节点
				insert2.Left = rightTree //插入节点的左孩子更新为插入位置之前的子树

				//加入结果中
				cur = append(cur, treeCopy)
			}
		}

		pre = cur
	}

	return pre
}

//拷贝树
func treeCopy(root *model.TreeNode) *model.TreeNode {
	if root == nil {
		return root
	}

	//递归拷贝，拷贝每个节点的值
	newRoot := &model.TreeNode{Val: root.Val}
	newRoot.Left = treeCopy(root.Left)
	newRoot.Right = treeCopy(root.Right)
	return newRoot
}

func main() {
	start1 := time.Now()
	generateTrees2(10)
	//fmt.Println("")
	fmt.Println(time.Since(start1).String())

	start2 := time.Now()
	generateTrees(10)
	//fmt.Println("")
	fmt.Println(time.Since(start2).String())

	//for _, v := range tree {
	//	PreOrder(v)
	//	fmt.Println("")
	//}
}

//先序遍历递归
func PreOrder(current *model.TreeNode) {
	if current != nil {
		fmt.Printf(" %v", current.Val)
		PreOrder(current.Left)
		PreOrder(current.Right)
	}
}
