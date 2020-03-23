package main

import "fmt"

/*
给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。

如果你最多只允许完成一笔交易（即买入和卖出一支股票一次），设计一个算法来计算你所能获取的最大利润。

注意你不能在买入股票前卖出股票。

示例 1:

输入: [7,1,5,3,6,4]
输出: 5
解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格。
示例 2:

输入: [7,6,4,3,1]
输出: 0
解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。
*/

func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	result := 0
	left := 0
	tmp := 0
	for left < n {
		right := n - 1
		for right > left {
			tmp = prices[right] - prices[left]
			if tmp > 0 && tmp > result {
				result = tmp
			}
			right--
		}

		left++
	}

	return result
}

//优化算法
//定义两个边界 i 和 j
// i 在 j 的左边
// i 右边的数要比 price[i]大，否则移动 i
func maxProfit2(prices []int) int {
	result := 0
	i := 0
	for j := 0; j < len(prices); j++ {
		if j > i {
			if prices[j] >= prices[i] {
				tmp := prices[j] - prices[i]
				if tmp > result {
					result = tmp
				}
			} else {
				i = j
			}
		}
	}
	return result
}

func main() {
	//input := []int{7, 1, 5, 3, 6, 4}
	//input := []int{7, 6, 4, 3, 1}
	//input := []int{2, 1, 4}
	input := []int{3, 2, 6, 5, 0, 3}
	fmt.Println(maxProfit2(input))
}
