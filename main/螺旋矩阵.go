package main

import "fmt"

/* 给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。
*
* 示例 1:
*
* 输入:
* [
* ⁠[ 1, 2, 3 ],
* ⁠[ 4, 5, 6 ],
* ⁠[ 7, 8, 9 ]
* ]
* 输出: [1,2,3,6,9,8,7,4,5]
*
*
* 示例 2:
*
* 输入:
* [
* ⁠ [1, 2, 3, 4],
* ⁠ [5, 6, 7, 8],
* ⁠ [9,10,11,12]
* ]
* 输出: [1,2,3,4,8,12,11,10,9,5,6,7]
*
*
 */

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return []int{}
	}

	n := len(matrix[0])

	var ret []int
	left := 0      //左上
	right := n - 1 //右下
	upper := 0     //上右
	down := m - 1  //下右

	for right >= 0 && down >= 0 {
		//从左到右
		l := left
		for ; l <= right; l++ {
			ret = append(ret, matrix[left][l])
		}
		upper++
		if upper > down {
			fmt.Println("====1=====")
			break
		}
		fmt.Println(ret)
		fmt.Println("=========")

		//从上到下
		up := upper
		for ; up <= down; up++ {
			ret = append(ret, matrix[up][right])
		}
		right--
		if left > right {
			fmt.Println("====2=====")
			break
		}

		fmt.Println(ret)
		fmt.Println("=========")
		//从右到左
		r := right
		for ; r >= left; r-- {
			ret = append(ret, matrix[down][r])
		}
		down--
		if upper > down {
			fmt.Println("====3=====")
			break
		}

		fmt.Println(ret)
		fmt.Println("=========")
		//从下到上d
		d := down
		for ; d >= upper; d-- {
			ret = append(ret, matrix[d][left])
		}
		left++
		if left > right {
			fmt.Println("====4=====")
			break
		}
		fmt.Println(ret)
		fmt.Println("=========")
	}

	fmt.Println(ret)
	return ret
}

/*
* 给定一个正整数 n，生成一个包含 1 到 n^2 所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。
 *
 * 示例:
 *
 * 输入: 3
 * 输出:
 * [
 * ⁠[ 1, 2, 3 ],
 * ⁠[ 8, 9, 4 ],
 * ⁠[ 7, 6, 5 ]
 * ]
 *
*/
func genNum(n int) int {
	return n
}

func generateMatrix(n int) [][]int {
	if n == 0 {
		return [][]int{}
	}

	var matrix [][]int
	for i := 0; i < n; i++ {
		var tmp []int
		for j := 0; j < n; j++ {
			tmp = append(tmp, 0)
		}
		matrix = append(matrix, tmp)
	}

	left := 0      //左上
	right := n - 1 //右下
	upper := 0     //上右
	down := n - 1  //下右

	start := 1
	for right >= 0 && down >= 0 {
		//从左到右
		l := left
		for ; l <= right; l++ {
			matrix[left][l] = start
			start++
		}
		upper++
		if upper > down {
			break
		}

		//从上到下
		up := upper
		for ; up <= down; up++ {
			matrix[up][right] = start
			start++
		}
		right--
		if left > right {
			fmt.Println("====2=====")
			break
		}

		//从右到左
		r := right
		for ; r >= left; r-- {
			matrix[down][r] = start
			start++
		}
		down--
		if upper > down {
			fmt.Println("====3=====")
			break
		}

		//从下到上d
		d := down
		for ; d >= upper; d-- {
			matrix[d][left] = start
			start++
		}
		left++
		if left > right {
			fmt.Println("====4=====")
			break
		}
	}

	fmt.Println(matrix)
	return [][]int{}
}

func main() {
	//x := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	//x := [][]int{{1, 2, 3}}
	//x := [][]int{{1, 11}, {2, 12}, {3, 13}, {4, 14}, {5, 15}, {6, 16}, {7, 17}, {8, 18}, {9, 19}, {10, 20}}
	//x := [][]int{{1, 11}, {2, 12}, {3, 13}, {4, 14}}
	//spiralOrder(x)

	generateMatrix(3)
}
