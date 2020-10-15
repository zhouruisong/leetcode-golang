package main

import (
	"container/list"
	"fmt"
)

/*
 * 给定一个 n × n 的二维矩阵表示一个图像。
 *
 * 将图像顺时针旋转 90 度。
 *
 * 说明：
 *
 * 你必须在原地旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要使用另一个矩阵来旋转图像。
 *
 * 示例 1:
 *
 * 给定 matrix =
 * [
 * ⁠ [1,2,3],
 * ⁠ [4,5,6],
 * ⁠ [7,8,9]
 * ],
 *
 * 原地旋转输入矩阵，使其变为:
 * [
 * ⁠ [7,4,1],
 * ⁠ [8,5,2],
 * ⁠ [9,6,3]
 * ]
 *
 * 示例 2:
 *
 * 给定 matrix =
 * [
 * ⁠ [ 5, 1, 9,11],
 * ⁠ [ 2, 4, 8,10],
 * ⁠ [13, 3, 6, 7],
 * ⁠ [15,14,12,16]
 * ],
 *
 * 原地旋转输入矩阵，使其变为:
 * [
 * ⁠ [15,13, 2, 5],
 * ⁠ [14, 3, 4, 1],
 * ⁠ [12, 6, 8, 9],
 * ⁠ [16, 7,10,11]
 * ]
 */

func rotate(matrix [][]int) {
	n := len(matrix)
	if n <= 1 {
		return
	}

	//先把矩阵每一行存储到栈中，然后在弹栈，采用列赋值,空间负责度O（n*n）
	mList := list.New()
	for i := 0; i < n; i++ {
		tmp := make([]int, n)
		copy(tmp, matrix[i])
		mList.PushFront(tmp)
		//如果这样，修改matrix里面的值，会修改掉mList的存储值，这里需要拷贝一份
		//mList.PushFront(matrix[i])
	}

	for i := 0; i < n; i++ {
		node := mList.Remove(mList.Front()).([]int)
		k := 0
		for j := 0; j < n; j++ {
			matrix[j][i] = node[k]
			k++
		}
	}
}

func rotate2(matrix [][]int) {
	n := len(matrix)
	if n <= 1 {
		return
	}

	//以左上角到右下角为对角线交换每个位置，然后每一行反转即可
	i := 0
	tmp := 0
	for i < n {
		j := i
		for j < n {
			tmp = matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = tmp
			j++
		}
		i++
	}

	//fmt.Println(matrix)

	//每一行反转
	i = 0
	for i < n {
		sliceReserve(matrix[i])
		i++
	}
}

func sliceReserve(nums []int) {
	n := len(nums)
	i := 0
	j := n - 1
	for i <= j {
		//交换nums[i],nums[j]
		tmp := nums[i]
		nums[i] = nums[j]
		nums[j] = tmp
		i++
		j--
	}

	//fmt.Println(nums)
}

func main() {
	x := [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	//x := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate2(x)
	fmt.Println(x)
	//x1 := []int{1, 2, 3, 4}
	//sliceReserve(x1)
	//fmt.Println(x1)
}
