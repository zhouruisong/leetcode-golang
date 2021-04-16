package main

import "fmt"

/*
假定你知道某只股票每一天价格的变动。
你最多可以同时持有一只股票。但你最多只能进行两次交易（一次买进和一次卖出记为一次交易。买进和卖出均无手续费）。
请设计一个函数，计算你所能获得的最大收益。
示例1
输入
复制
[8,9,3,5,1,3]
返回值
复制
4
说明
第三天买进，第四天卖出，第五天买进，第六天卖出。总收益为4。
*/

func maxProfit(prices []int) int {
	// write code here
	if len(prices) <= 1 {
		return 0
	}
	var cost1, cost2, profit1, profit2 = prices[0], prices[0], 0, 0
	for i := range prices {
		if cost1 > prices[i] {
			cost1 = prices[i]
		} else if prices[i]-cost1 > profit1 {
			profit1 = prices[i] - cost1
		}

		if cost2 > prices[i]-profit1 {
			cost2 = prices[i] - profit1
		} else if prices[i]-cost2 > profit2 {
			profit2 = prices[i] - cost2
		}
	}
	return profit2
}

func main() {
	input := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(input))
}
