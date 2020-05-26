package main

import "fmt"

/*
下面给出几个求卡特兰数的公式，用h(n)表示卡特兰数的第n项，其中h(0)=1，h(1)=1
公式一：h(n)= h(0)*h(n-1)+h(1)*h(n-2) + ... + h(n-1)h(0) (n>=2)
*/

/*
Catalan数的典型应用：

1.由n个+1和n个-1组成的排列中，满足前缀和>=0的排列有Catalan(N)种。

2.括号化问题。左括号和右括号各有n个时，合法的括号表达式的个数有Catalan(N)种。

3.有n+1个数连乘，乘法顺序有Catalan(N)种,相当于在式子上加括号。

4.n个数按照特定顺序入栈，出栈顺序随意，可以形成的排列的种类有Catalan(N)种。

5.给定N个节点，能构成Catalan(N)种种形状不同的二叉树。

6.n个非叶节点的满二叉树的形态数为Catalan(N)。

7.对于一个n*n的正方形网格，每次只能向右或者向上移动一格，那么从左下角到右上角的不同种类有Catalan(N)种。

8.对于在n位的2进制中，有m个0，其余为1的catalan数为：C（n,m）-C(n,m-1)。

9.对凸n+2边形进行不同的三角形分割（只连接顶点对形成n个三角形）数为Catalan(N)。

10.将有2n个元素的集合中的元素两两分为n个子集，若任意两个子集都不交叉，那么我们称此划分为一个不交叉划分。此时不交叉的划分数是Catalan(N)。

11.n层的阶梯切割为n个矩形的切法数也是Catalan(N)。

12.在一个2*n的格子中填入1到2n这些数值使得每个格子内的数值都比其右边和上边的所有数值都小的情况数也是Catalan(N)。
*/

func main() {
	i, j := 2, 0
	n := 20 //求500以内的卡特兰数
	arr := make([]uint64, 0, 0)
	arr = append(arr, 1)
	arr = append(arr, 1)
	tmp := uint64(0)
	for ; i <= n; i++ {
		for ; j < i; j++ {
			tmp += (arr[j] * arr[i-j-1])
		}

		arr = append(arr, tmp)
		tmp = 0
		j = 0
	}

	fmt.Println(arr)
}
