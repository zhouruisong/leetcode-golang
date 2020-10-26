package main

import (
	"fmt"
	"math"
)

/*
给定正整数 n，找到若干个完全平方数（比如 1, 4, 9, 16, ...）使得它们的和等于 n。你需要让组成和的完全平方数的个数最少。
示例 1:
输入: n = 12
输出: 3
解释: 12 = 4 + 4 + 4.
示例 2:

输入: n = 13
输出: 2
解释: 13 = 4 + 9.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/perfect-squares
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
解法一 回溯法
相当于一种暴力的方法，去考虑所有的分解方案，找出最小的解，举个例子。

n = 12
先把 n 减去一个平方数，然后求剩下的数分解成平方数和所需的最小个数

把 n 减去 1, 然后求出 11 分解成平方数和所需的最小个数,记做 n1
那么当前方案总共需要 n1 + 1 个平方数

把 n 减去 4, 然后求出 8 分解成平方数和所需的最小个数,记做 n2
那么当前方案总共需要 n2 + 1 个平方数

把 n 减去 9, 然后求出 3 分解成平方数和所需的最小个数,记做 n3
那么当前方案总共需要 n3 + 1 个平方数

下一个平方数是 16, 大于 12, 不能再分了。

接下来我们只需要从 (n1 + 1), (n2 + 1), (n3 + 1) 三种方案中选择最小的个数,
此时就是 12 分解成平方数和所需的最小个数了

至于求 11、8、3 分解成最小平方数和所需的最小个数继续用上边的方法去求

直到如果求 0 分解成最小平方数的和的个数, 返回 0 即可
代码的话，就是回溯的写法，或者说是 DFS。
*/

func numSquares(n int) int {
	return numSquaresHelper(n)
}

func numSquaresHelper(n int) int {
	if n == 0 {
		return 0
	}
	count := math.MaxInt32
	//依次减去一个平方数
	for i := 1; i*i <= n; i++ {
		//选最小的
		count = (int)(math.Min(float64(count), float64(numSquaresHelper(n-i*i)+1)))
	}
	return count
}

//动态规划
func numSquares2(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0

	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	//依次求出 1, 2... 直到 n 的解
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt32
		//依次减去一个平方数
		for j := 1; j*j <= i; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]
}

//BFS
func numSquares3(n int) int {
	queue := make([]int, 0)
	visited := make([]bool, n+1)
	level := 0

	queue = append(queue, n)

	for len(queue) > 0 {
		size := len(queue)
		level++
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			for j := 1; j*j <= cur; j++ {
				next := cur - j*j
				if next == 0 {
					return level
				}
				if visited[next] == false {
					queue = append(queue, next)
					visited[next] = true
				}
			}
		}
	}
	return -1
}
func main() {
	fmt.Println(numSquares3(12))
}
