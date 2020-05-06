package main

import "fmt"

/*
给定一个整数 n，求以 1 ... n 为节点组成的二叉搜索树有多少种？
示例:
输入: 3
输出: 5
解释:
给定 n = 3, 一共有 5 种不同结构的二叉搜索树:

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3
*/

/*
用个数组记录之前的情况就快多了
首先 我们知道 n==1 与 n==2 返回的就是本身
在这个基础上进行计算
举例：
1 2 3
变成二叉搜索树无非就是谁做root节点的问题，当然1,2,3谁都可以
我们已经知道n==1与n==2的情况，保存在vector<int> work中，work[1]=1,work[2]=2
那个三个节点的1,2,3是多少中情况呢？很明显是1,2,3都做一次root节点
1、1为root节点，此时右边的2个节点是右子树，而2个节点作为一个树的情况我们已经知道了，直接查询work[2]=2
2、2为root节点，此时左右都是1个节点的子树，也就是work[1]*work[1];
3、3为root节点，此时左边的2个节点是左子树，同1为work[2]=2
因此work[3]=2+1+2=5
之后的同理迭代运算即可
*/

func numTrees(n int) int {
	if n == 1 || n == 2 {
		return n
	}

	res := []int{1, 1, 2}
	index := 3
	tmp := 0 //累加
	for ; index <= n; index++ {
		for i := 1; i <= index; i++ {
			//fmt.Println(res[i-1], res[index-i])
			tmp += (res[i-1] * res[index-i])
		}
		res = append(res, tmp)
		tmp = 0
	}

	return res[len(res)-1]
}

//dp
func numTrees2(n int) int {
	g := make([]int, n+1)
	g[0] = 1
	g[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			g[i] += g[j-1] * g[i-j]
		}
	}
	return g[n]
}

func main() {
	fmt.Println(numTrees(3))
}
