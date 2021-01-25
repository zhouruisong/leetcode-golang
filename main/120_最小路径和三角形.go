package main

import (
	"fmt"
	"math"
)

/*
 * 给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。
 *
 * 例如，给定三角形：
 *
 * [
 * ⁠    [2],
 * ⁠   [3,4],
 * ⁠  [6,5,7],
 * ⁠ [4,1,8,3]
 * ]
 *
 *
 * 自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。
 *
 * 说明：
 *
 * 如果你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题，那么你的算法会很加分。
 */

func minimumTotal(triangle [][]int) int {
	/*
		这题用DP解就可以了，简单描述下过程
		初始：
		[
		[2],
		[3,4],
		[6,5,7]
		]
		第一轮过后：3+2=5，4+2=6
		[
		[2],
		[5,6],
		[6,5,7]
		]
		第二轮过后：6+5=11，[5+5=10,5+6=11,取小值10]，7+6=13
		[
		[2],
		[5,6],
		[11,10,13]
		]
	*/
	m := len(triangle)
	if m < 1 {
		return 0
	}

	if m == 1 {
		return triangle[0][0]
	}

	for i := 0; i < len(triangle); i++ {
		fmt.Println(triangle[i])
	}

	fmt.Println("")

	for i := 1; i < m; i++ {
		for j := 0; j < i; j++ {
			tmp := triangle[i-1][j] + triangle[i][j]
			if j-1 >= 0 {
				tmp = int(math.Min(float64(tmp), float64(triangle[i-1][j-1]+triangle[i][j])))
			}
			triangle[i][j] = tmp
			for i := 0; i < len(triangle); i++ {
				fmt.Println(triangle[i])
			}
			fmt.Println("--------")
		}

		//对角线
		triangle[i][i] = triangle[i][i] + triangle[i-1][i-1]

		for i := 0; i < len(triangle); i++ {
			fmt.Println(triangle[i])
		}
		fmt.Println("========")

	}

	//fmt.Println(triangle)
	//找出最后一行的最小值即可
	tmp := triangle[len(triangle)-1][0]
	for _, v := range triangle[len(triangle)-1] {
		if v < tmp {
			tmp = v
		}
	}

	return tmp
}

func main() {
	x := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	fmt.Println(minimumTotal(x))
}
