package main

import (
	"fmt"
	"sort"
)

/*
给定一个未排序的整数数组，找出最长连续序列的长度。

要求算法的时间复杂度为 O(n)。

示例:

输入: [100, 4, 200, 1, 3, 2]
输出: 4
解释: 最长连续序列是 [1, 2, 3, 4]。它的长度为 4。
*/

/*
我们考虑枚举数组中的每个数 x，考虑以其为起点，不断尝试匹配x+1,x+2,⋯ 是否存在，假设最长匹配到了 x+y，
那么以 x 为起点的最长连续序列即为x,x+1,x+2,⋯,x+y，其长度为 y+1，我们不断枚举并更新答案即可。
对于匹配的过程，暴力的方法是 O(n) 遍历数组去看是否存在这个数，但其实更高效的方法是用一个哈希表存储数组中的数，
这样查看一个数是否存在即能优化至 O(1) 的时间复杂度。
*/

func longestConsecutive(nums []int) int {
	numSet := map[int]bool{}
	for _, num := range nums {
		numSet[num] = true
	}
	longestStreak := 0
	for num := range numSet {
		if !numSet[num-1] {
			currentNum := num
			currentStreak := 1
			for numSet[currentNum+1] {
				currentNum++
				currentStreak++
			}
			if longestStreak < currentStreak {
				longestStreak = currentStreak
			}
		}
	}
	return longestStreak
}

/**
 * max increasing subsequence
 * @param arr int整型一维数组 the array
 * @return int整型
 */
//o(nlogn)
func longestConsecutive2(arr []int) int {
	// write code here
	if len(arr) == 0 {
		return 0
	}

	n := len(arr)
	sort.Ints(arr)
	count := 1
	result := 1

	maxf := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	for i := 0; i < n-1; i++ {
		if arr[i+1]-arr[i] == 1 {
			count++
		} else if arr[i+1]-arr[i] == 0 {
			continue
		} else {
			count = 1
		}

		result = maxf(count, result)
	}
	return result
}

func main() {
	arr := []int{100, 4, 200, 1, 3, 2}
	fmt.Print(longestConsecutive(arr))
}
