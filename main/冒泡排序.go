package main

import "fmt"

func bubbleSort(nums []int) []int {
	n := len(nums)
	if n <= 1 {
		return nums
	}

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				//fmt.Println(nums)
			}
		}
	}

	return nums
}

func bubbleSort2(nums []int) []int {
	n := len(nums)
	if n <= 1 {
		return nums
	}

	for i := 0; i < n-1; i++ {
		//默认没有交换
		isChange := false
		for j := 0; j < n-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				isChange = true
				//fmt.Println(nums)
			}
			//优化
			if isChange {
				break
			}
		}
	}

	return nums
}

func main() {
	nums := []int{9, 2, 8, 3, 4, 6, 5, 10, 11}
	nums2 := []int{9, 2, 8, 3, 4, 6, 5, 10, 11}

	fmt.Println(bubbleSort(nums))
	fmt.Println("=========")
	fmt.Println(bubbleSort2(nums2))

}
