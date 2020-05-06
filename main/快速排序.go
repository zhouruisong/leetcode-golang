package main

import "fmt"

func QuickSort(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}

	splitdata := arr[0]       //基准
	low := make([]int, 0, 0)  //存储比我小的
	higt := make([]int, 0, 0) //存储比我大的
	mid := make([]int, 0, 0)  //存储和我相等的
	mid = append(mid, splitdata)

	for i := 1; i < n; i++ {
		if arr[i] < splitdata {
			low = append(low, arr[i])
		} else if arr[i] > splitdata {
			higt = append(higt, arr[i])
		} else {
			mid = append(mid, arr[i])
		}
	}

	low, higt = QuickSort(low), QuickSort(higt) //切割递归处理
	marr := append(append(low, mid...), higt...)
	return marr
}

func main() {
	arr := []int{9, 2, 8, 3, 4, 6, 1, 5, 10, 11}
	fmt.Println(QuickSort(arr))
}
