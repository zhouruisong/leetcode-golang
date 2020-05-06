package main

import "fmt"

//o(n*n)
func InsertSort(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}

	//把2位置插入到9的位置
	//1 9 2
	//1 2 9
	for i := 0; i < n; i++ {
		bakcup := arr[i]
		//从2的位置往后移动，找到插入位置
		j := i - 1 //从上一个位置循环找到位置插入
		for j >= 0 && bakcup < arr[j] {
			arr[j+1] = arr[j] //从前往后移动
			j--
		}

		arr[j+1] = bakcup
		fmt.Println(arr)
	}

	return arr
}

func main() {
	arr := []int{1, 9, 2, 8, 3, 4, 6, 5, 10, 11}
	fmt.Println(InsertSort(arr))
}
