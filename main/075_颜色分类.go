package main

import (
	"fmt"
)

/*
给定一个包含红色、白色和蓝色，一共 n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。

此题中，我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。

注意:
不能使用代码库中的排序函数来解决这道题。

示例:

输入: [2,0,2,1,1,0]
输出: [0,0,1,1,2,2]
进阶：

一个直观的解决方案是使用计数排序的两趟扫描算法。
首先，迭代计算出0、1 和 2 元素的个数，然后按照0、1、2的排序，重写当前数组。
你能想出一个仅使用常数空间的一趟扫描算法吗？
*/

/*
我们可以考虑对数组进行两次遍历。在第一次遍历中，我们将数组中所有的 00 交换到数组的头部。在第二次遍历中，
我们将数组中所有的 11 交换到头部的 00 之后。此时，所有的 22 都出现在数组的尾部，这样我们就完成了排序。
*/

/*
适配各种判断条件
*/
//函数指针
type doFunc func(c, target int) bool

func check(c, target int) bool {
	return (c % target) == 0
}

func sortColorsFunc(nums []int) {
	f := check
	swapColorsFunc(nums, 2, f)
	//count0 := swapColorsFunc(nums, 0, f) // 把 0 排到前面
	//swapColorsFunc(nums[count0:], 1, f)  // nums[:count0] 全部是 0 了，对剩下的 nums[count0:] 把 1 排到前面
}

func swapColorsFunc(colors []int, target int, f doFunc) (countTarget int) {
	for i, c := range colors {
		if f(c, target) {
			colors[i], colors[countTarget] = colors[countTarget], colors[i]
			countTarget++
		}
	}
	return
}

func swapColors(colors []int, target int) (countTarget int) {
	for i, c := range colors {
		if c == target {
			colors[i], colors[countTarget] = colors[countTarget], colors[i]
			countTarget++
		}
	}
	return
}

func sortColors(nums []int) {
	count0 := swapColors(nums, 0) // 把 0 排到前面
	swapColors(nums[count0:], 1)  // nums[:count0] 全部是 0 了，对剩下的 nums[count0:] 把 1 排到前面
}

/*
双指针
方法一需要进行两次遍历，那么我们是否可以仅使用一次遍历呢？我们可以额外使用一个指针，即使用两个指针分别用来交换 00 和 11。
*/
func sortColors2(nums []int) {
	p0, p2 := 0, len(nums)-1
	for i := 0; i <= p2; i++ {
		//双指针p0,p2,p0指向头部,p2指向尾部
		//从左向右遍历,当前位置遇到0 交换p0, 遇到2, 循环交换p2,直到当前位置不是2为止
		for ; i <= p2 && nums[i] == 2; p2-- {
			nums[i], nums[p2] = nums[p2], nums[i]
		}
		if nums[i] == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			p0++
		}
	}
}

func main() {
	x := []int{2, 0, 2, 1, 1, 0}
	//x := []int{2, 3, 2, 1, 1, 4}
	//sortColors(x)
	sortColorsFunc(x)
	fmt.Print(x)
}
