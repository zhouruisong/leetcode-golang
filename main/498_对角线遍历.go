package main

/*
给定一个含有 M x N 个元素的矩阵（M 行，N 列），请以对角线遍历的顺序返回这个矩阵中的所有元素，对角线遍历如下图所示。

 

示例:

输入:
[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]

输出:  [1,2,4,7,5,3,6,8,9]
 */

func findDiagonalOrder(matrix [][]int) []int {
	//从左上到右下
	if len(matrix) == 0 {
		return []int{}
	}

	res, temp := []int{}, []int{}
	row, col, flag, i, j := len(matrix), len(matrix[0]), 0, 0, 0
	if row == 1 {
		return matrix[0]
	}

	for j <= col-1 {
		m, n := i, j
		//添加对角线下方元素
		temp = append(temp, matrix[i][j])

		//查询对角线上方元素
		for m != 0 && n != col-1 {
			m--
			n++
			temp = append(temp, matrix[m][n])
		}
		//间隔反转
		if flag%2 != 0 {
			resserve(temp)
		}
		flag++
		res = append(res, temp...)
		temp = temp[:0]
		if i != row-1 {
			i++
		} else {
			j++
		}
	}

	return res
}

func resserve(arr []int) {
	i, j := 0, len(arr)-1
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}
