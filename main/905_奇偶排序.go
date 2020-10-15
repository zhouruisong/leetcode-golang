package main

import "fmt"

func OddEven(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}

	isSorted := false //是否还需要排序，不需要奇数位偶数位排序的时候

	for isSorted == false {
		isSorted = true
		for i := 1; i < n-1; i = i + 2 { //奇数位
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				isSorted = false
			}
		}

		fmt.Println("1-->", arr)
		for i := 0; i < n-1; i = i + 2 { //偶数位
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				isSorted = false
			}
		}
		fmt.Println("0-->", arr)

	}

	return arr
}

func main() {
	arr := []int{9, 2, 8, 3, 4, 6, 1, 5, 10, 11}
	fmt.Println(OddEven(arr))
}
