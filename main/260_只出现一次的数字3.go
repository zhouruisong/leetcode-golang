package main

/*
给定一个整数数组 nums，其中恰好有两个元素只出现一次，其余所有元素均出现两次。 找出只出现一次的那两个元素。

示例 :

输入: [1,2,1,3,2,5]
输出: [3,5]
注意：

结果输出的顺序并不重要，对于上面的例子， [5, 3] 也是正确答案。
你的算法应该具有线性时间复杂度。你能否仅使用常数空间复杂度来实现？
 */

func singleNumber(nums []int) []int {
	//异或运算o(n)
	n := len(nums)
	if n == 0 {
		return []int{}
	}

	// 得到 (3 ^ 5)
	xorRes := 0
	for _, v := range nums {
		xorRes ^= v
	}

	temp := 1 // 用来标志第几位是 1
	for {
		if (xorRes & 1) == 1 {
			break
		}
		temp = temp << 1
		xorRes = xorRes >> 1 // 右移，从低到高
	}

	//分成两组,分别异或
	res := make([]int, 2)
	for _, v := range nums {
		if v&temp == 0 {
			res[0] ^= v
		} else {
			res[1] ^= v
		}
	}

	return res
}
