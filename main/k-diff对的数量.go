package main

import (
	"fmt"
)

/*
给定整数nums和整数k的数组，返回数组中唯一的k-diff对的数量。
k-diff对是一个整数对（nums [i]，nums [j]），其中以下是正确的：
0 <= i，j <长度
我！= j
| nums [i]-nums [j] | == k
请注意| val | 表示val的绝对值。

范例1：
输入：nums = [3,1,4,1,5]，k = 2
输出2
说明：数组中有两个2-diff对，（1、3）和（3、5）。
尽管输入中有两个1，但我们只应返回唯一对的数量。

范例2：
输入：nums = [1,2,3,4,5]，k = 1
输出4
说明：数组中有四个1-diff对，分别为（1、2），（2、3），（3、4）和（4、5）。

范例3：
输入：nums = [1,3,1,5,4]，k = 0
输出1
说明：数组中有一个（0，diff）对（1，1）。

范例4：
输入：nums = [1,2,4,4,3,3,0,9,2,3]，k = 3
输出2

范例5：
输入：nums = [-1，-2，-3]，k = 1
输出2


限制条件：
1 <=数字长度<= 104
-107 <= nums [i] <= 107
0 <= k <= 107
*/
func findPairs(nums []int, k int) int {
	mp := make(map[int]bool, len(nums))
	res := [][]int{}

	for _, v := range nums {
		mp[v] = true
	}

	fmt.Println(mp)

	//去重使用
	used := make(map[string]bool, 0)
	for v, _ := range mp {
		other1 := v - k
		other2 := v + k

		if _, ok := mp[other1]; ok {
			key := fmt.Sprintf("%v%v", v-k, v)
			if _, ok := used[key]; !ok {
				res = append(res, []int{v - k, v})
				used[key] = true
			}

		} else if _, ok := mp[other2]; ok {
			key := fmt.Sprintf("%v%v", v, v+k)
			if _, ok := used[key]; !ok {
				res = append(res, []int{v, v + k})
				used[key] = true
			}
		}
	}

	fmt.Println(res)
	return len(res)
}

func main() {
	//nums := []int{1, 2, 3, 4, 5}
	//k := 1

	//nums := []int{3, 1, 4, 1, 5}
	//k := 2

	//nums := []int{1, 2, 4, 4, 3, 3, 0, 9, 2, 3}
	//k := 3

	nums := []int{-1, -2, -3}
	k := 1

	fmt.Println(findPairs(nums, k))
}
