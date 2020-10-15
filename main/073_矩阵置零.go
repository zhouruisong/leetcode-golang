package main

import "fmt"

/*
给定一个 m x n 的矩阵，如果一个元素为 0，则将其所在行和列的所有元素都设为 0。请使用原地算法。

示例 1:

输入:
[
  [1,1,1],
  [1,0,1],
  [1,1,1]
]
输出:
[
  [1,0,1],
  [0,0,0],
  [1,0,1]
]
示例 2:

输入:
[
  [0,1,2,0],
  [3,4,5,2],
  [1,3,1,5]
]
输出:
[
  [0,0,0,0],
  [0,4,5,0],
  [0,3,1,0]
]
进阶:

一个直接的解决方案是使用  O(mn) 的额外空间，但这并不是一个好的解决方案。
一个简单的改进方案是使用 O(m + n) 的额外空间，但这仍然不是最好的解决方案。
你能想出一个常数空间的解决方案吗？
*/

func setZeroes(matrix [][]int) {
	//先找到为0的位置，记录下来
	//在从每个位置开始，横向从当前位置前和后，上和下遍历置零
	m := len(matrix)
	if m == 0 {
		return
	}

	n := len(matrix[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}

	fmt.Println()

	var location [][]int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				location = append(location, []int{i, j})
			}
		}
	}

	for _, v := range location {
		hang := v[0]
		lie := v[1]
		setzero(hang, lie, m, n, matrix)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}
}

func setzero(i, j, m, n int, matrix [][]int) {
	for k := 0; k <= j; k++ {
		matrix[i][k] = 0
	}

	for k := j + 1; k < n; k++ {
		matrix[i][k] = 0
	}

	for k := 0; k < i; k++ {
		matrix[k][j] = 0
	}

	for k := i + 1; k < m; k++ {
		matrix[k][j] = 0
	}
}

func main() {
	//x := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	x := [][]int{{1, 1, 1}, {1, 0, 1}}

	setZeroes(x)
}
