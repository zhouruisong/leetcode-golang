package main

import "fmt"

func maxSumDivThree(nums []int) int {
	/*暴力法是找出所有和的情况，然后遍历一遍，取%3==0的最大值
	优化思路是：我们只需要动态更新遍历过的元素之和中 n%3==0，n%3==1，n%3==2的三个最大值即可舍弃其它和的情况
	复杂度o（n)*/
	ans := make([]int, 3) //ans[0],ans[1],ans[2]分别保存遍历过的元素的和中%3==0，%3==1，%3==2的最大值
	temp := make([]int, 3)
	maxf := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	for _, num := range nums { //遍历nums
		for _, a := range ans { //将ans中每个元素与num求和
			mod := (num + a) % 3
			if mod == 0 { //如果当前和 %3==0  且大于之前的temp[0]则更新temp[0]
				temp[0] = maxf(num+a, temp[0])
			} else if mod == 1 { //如果当前和 %3==1  且大于之前的temp[1]则更新temp[1]
				temp[1] = maxf(num+a, temp[1])
			} else if mod == 2 { //如果当前和 %3==2  且大于之前的temp[2]则更新temp[2]
				temp[2] = maxf(num+a, temp[2])
			}
		}
		copy(ans, temp) //将修正过的temp赋给ans
		//ans = temp //将修正过的temp赋给ans
	}

	return ans[0]
}

func main() {
	nums := []int{3, 6, 5, 1, 8}
	fmt.Println(maxSumDivThree(nums))
}
