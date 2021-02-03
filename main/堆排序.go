package main

import "fmt"

//o(nlogn)
func HeapSort(nums []int) []int {
	n := len(nums)
	if n <= 1 {
		return nums
	}

	for i := 0; i < n; i++ {
		lastLen := n - i
		HeapSortMax(nums, lastLen) //每一次最大值在头部，为了将最大值放到尾部，需要替换nums[0]和最后的位置lastLen
		fmt.Println(nums)

		if i < n {
			nums[0], nums[lastLen-1] = nums[lastLen-1], nums[0]
		}
		fmt.Println("ex", nums)
	}

	return nums
}

//最大节点在跟节点
func HeapSortMax(nums []int, n int) []int {
	if n <= 1 {
		return nums
	}

	depth := n/2 - 1              //深度, 最后一个非叶子节点
	for i := depth; i >= 0; i-- { //循环所有的三节点
		topMax := i      //假定最大的在i的位置
		left := 2*i + 1  //左孩子
		right := 2*i + 2 //右孩子

		//fmt.Println(i, nums[left], nums[topMax], nums)

		if left <= n-1 && nums[left] > nums[topMax] { //如果左边比我大，记录最大
			topMax = left
		}

		if right <= n-1 && nums[right] > nums[topMax] { //如果右边边比我大，记录最大
			topMax = right
		}

		if topMax != i { //确保i的值就是最大,交换nums[i]和nums[topMax]
			nums[i], nums[topMax] = nums[topMax], nums[i]
		}
		//fmt.Println(i, nums)
	}

	return nums
}

func main() {
	//nums := []int{1, 9, 2, 8, 3, 4, 6, 5, 10, 7}
	nums := []int{7, 3, 8, 5, 1, 2}
	fmt.Println(HeapSort(nums))
}
