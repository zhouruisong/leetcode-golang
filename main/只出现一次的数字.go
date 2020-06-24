package main

import "fmt"

/*
* 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
*
* 说明：
*
* 你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
*
* 示例 1:
*
* 输入: [2,2,1]
* 输出: 1
*
*
* 示例 2:
*
* 输入: [4,1,2,1,2]
* 输出: 4
*
 */

//空间分配多
func singleNumber(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}

	mp := make(map[int]int)
	for i := 0; i < n; i++ {
		if _, ok := mp[nums[i]]; ok {
			mp[nums[i]]++
			continue
		}

		mp[nums[i]]++
	}

	loc := 0
	for k, v := range mp {
		if v == 1 {
			loc = k
			break
		}
	}

	return loc
}

//异或
func singleNumber2(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}

	ret := 0
	for i := 0; i < n; i++ {
		ret ^= nums[i]
	}
	return ret
}

func main() {
	arr := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(arr))
}
