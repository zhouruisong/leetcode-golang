package main

import "fmt"

func firstMissingPositive(nums []int) int {
	//hashmap o(n)空间复杂度
	mp := make(map[int]bool, len(nums))
	for _, v := range nums {
		mp[v] = true
	}

	for i := 1; i <= len(nums); i++ {
		if _, ok := mp[i]; !ok {
			return i
		}
	}

	return len(nums) + 1
}

func main() {
	//arr := []int{3, 4, -1, 1}
	arr := []int{1, 2, 0}
	fmt.Println(firstMissingPositive(arr))
}
