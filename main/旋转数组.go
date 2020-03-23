package main

import "fmt"

/*
* 给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
*
* 示例 1:
*
* 输入: [1,2,3,4,5,6,7] 和 k = 3
* 输出: [5,6,7,1,2,3,4]
* 解释:
* 向右旋转 1 步: [7,1,2,3,4,5,6]
* 向右旋转 2 步: [6,7,1,2,3,4,5]
* 向右旋转 3 步: [5,6,7,1,2,3,4]
*
*
* 示例 2:
*
* 输入: [-1,-100,3,99] 和 k = 2
* 输出: [3,99,-1,-100]
* 解释:
* 向右旋转 1 步: [99,-1,-100,3]
* 向右旋转 2 步: [3,99,-1,-100]
*
* 说明:
*
*
* 尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。
* 要求使用空间复杂度为 O(1) 的 原地 算法。
*
*
 */

type node struct {
	Val  int
	Next *node
}

//环形链表解法
func rotate(nums []int, k int) {
	n := len(nums)
	head := &node{Val: -1, Next: nil}
	tmp := head
	for i := 0; i < n; i++ {
		tmp.Next = &node{Val: nums[i], Next: nil}
		tmp = tmp.Next
	}

	head = head.Next
	//for i := 0; i < n; i++ {
	//	fmt.Println(head.Val)
	//	head = head.Next
	//}

	//环形链表
	tmp.Next = head
	p1 := tmp

	j := k % n
	//查找到底k个位置，在该处断开即可
	for i := 0; i < n-j; i++ {
		p1 = p1.Next
		head = head.Next
	}

	p1.Next = nil
	i := 0

	for head != nil {
		nums[i] = head.Val
		head = head.Next
		i++
	}

	fmt.Println(nums)
}

//[1,2,3,4,5,6,7] 和 k = 3
//数组最后元素向右侧移动k次
//o(kn)时间复杂度，空间复杂度o(1)
func rotate2(nums []int, k int) {
	n := len(nums)
	var tmp int
	for i := 0; i < k; i++ {
		tmp = nums[n-1]
		for j := n - 1; j > 0; j-- {
			nums[j] = nums[j-1]
		}
		nums[0] = tmp
	}

	fmt.Println(nums)
}


func main() {
	rotate2([]int{1, 2, 3, 4, 5, 6, 7}, 2)
}
