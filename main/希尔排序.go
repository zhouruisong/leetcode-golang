package main

import "fmt"

//步长收缩的算法
//多线程排序
func ShellSort(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}

	gap := len(arr) / 2
	for gap > 0 {
		for i := 0; i < gap; i++ { //处理每个元素的步长
			ShellSortStep(arr, i, gap)
		}
		gap /= 2
	}

	return arr
}

func ShellSortStep(arr []int, start, gap int) {
	n := len(arr)
	for i := start + gap; i < n; i += gap {
		backup := arr[i] //备份插入的数据
		j := i - gap     //上一个位置循环找到位置插入
		for j >= 0 && backup < arr[j] {
			arr[j+gap] = arr[j] // 从前往后移动
			j -= gap
		}

		arr[j+gap] = backup
	}

	fmt.Println(gap, arr)
}

func main() {
	nums := []int{1, 9, 2, 8, 3, 4, 6, 5, 10, 7}
	fmt.Println(ShellSort(nums))
}
