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
	res := append(append(low, mid...), higt...)
	return res
}

func QuickSort2(arr []int, low, high int) {
	if low > high {
		return
	}

	i := low
	j := high
	loc := arr[low]

	for i < j {
		for i < j && arr[j] >= loc {
			j--
		}
		arr[i], arr[j] = arr[j], arr[i]

		for i < j && arr[i] <= loc {
			i++
		}
		arr[i], arr[j] = arr[j], arr[i]
	}

	arr[i] = loc

	QuickSort2(arr, 0, i-1)
	QuickSort2(arr, i+1, high)
}

func main() {
	arr := []int{9, 2, 8, 3, 4, 6, 1, 5, 10, 11}
	QuickSort2(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
