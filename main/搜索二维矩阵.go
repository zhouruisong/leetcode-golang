package main

import "fmt"

/*
编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：

每行中的整数从左到右按升序排列。
每行的第一个整数大于前一行的最后一个整数。
示例 1:

输入:
matrix = [
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
target = 3
输出: true
示例 2:

输入:
matrix = [
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
target = 13
输出: false
*/

//o(m+n)
func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}

	n := len(matrix[0])
	//采用二分查找思想
	row, col := 0, n-1
	for col >= 0 && row < m {
		if target == matrix[row][col] {
			return true
		} else if target > matrix[row][col] { //向下一行查找
			row++
		} else if target < matrix[row][col] { //向右侧查找
			col--
		}
	}
	return false
}

//o(m+n)
func searchMatrix2(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	row := len(matrix)
	col := len(matrix[0])
	//1. 规定起始点
	i, j := row-1, 0
	for i >= 0 && j < col {
		if matrix[i][j] == target { //2.找到则返回
			return true
		} else if matrix[i][j] < target { //3.小于target，向右查找
			j++
		} else { //4.大于target，向上查找
			i--
		}
	}

	return false
}

func main() {
	m := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 50},
	}
	//m := [][]int{{2}, {3}}
	fmt.Println(searchMatrix(m, 32))
}
