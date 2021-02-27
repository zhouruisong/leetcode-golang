package main

/*
题目描述
一个数组A中存有N（N&gt0）个整数，在不允许使用另外数组的前提下，将每个整数循环向右移M（M>=0）个位置，即将A中的数据由（A0 A1 ……AN-1 ）
变换为（AN-M …… AN-1 A0 A1 ……AN-M-1 ）（最后M个数循环移至最前面的M个位置）。如果需要考虑程序移动数据的次数尽量少，要如何设计移动的方法？
示例1
输入
6,2,[1,2,3,4,5,6]
返回值
[5,6,1,2,3,4]
*/

/**
 * 旋转数组
 * @param n int整型 数组长度
 * @param m int整型 右移距离
 * @param a int整型一维数组 给定数组
 * @return int整型一维数组
 */
func reverse(a []int, left, right int) {
	for left < right {
		a[left], a[right] = a[right], a[left]
		left++
		right--
	}
}

func solve(n int, m int, a []int) []int {
	// write code here
	/*
	   三次翻转
	   假设 n=7且 k=3
	   原始数组                  : 1 2 3 4 5 6 7
	   反转所有数字后             : 7 6 5 4 3 2 1
	   反转前 k 个数字后          : 5 6 7    4 3 2 1
	   反转后 n-k 个数字后        : 5 6 7    1 2 3 4 --> 结果
	*/
	m = m % n
	if m == 0 {
		return a
	}

	reverse(a, 0, n-m-1)
	reverse(a, n-m, n-1)
	reverse(a, 0, n-1)
	return a
}
