package main

import "fmt"

/*
实现 int sqrt(int x) 函数。

计算并返回 x 的平方根，其中 x 是非负整数。

由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。

示例 1:

输入: 4
输出: 2
示例 2:

输入: 8
输出: 2
说明: 8 的平方根是 2.82842...,
     由于返回类型是整数，小数部分将被舍去。
*/

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}

	N := x / 2
	r := 1
	next := 0
	cur := 0
	nr := 0
	for ; r <= N; r++ {
		nr = r + 1
		cur = r * r
		next = nr * nr
		if cur == x {
			break
		}
		if cur < x && next > x {
			break
		}
	}

	return r
}

//二分查找方法
func mySqrt2(x int) int {
	//二分查找方法
	if x < 2 {
		return x
	}

	left := 2
	right := x/2

	for left < right {
		mid := (right + left) / 2 + 1
		if mid * mid <= x {
			left = mid
		} else {
			right = mid - 1
		}
	}

	return right
}

func main() {
	x := 8
	//fmt.Println(mySqrt(x))
	fmt.Println(mySqrt2(x))

}
