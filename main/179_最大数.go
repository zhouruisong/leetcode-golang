package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
* 给定一组非负整数 nums，重新排列每个数的顺序（每个数不可拆分）使之组成一个最大的整数。
 *
 * 注意：输出结果可能非常大，所以你需要返回一个字符串而不是整数。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [10,2]
 * 输出："210"
 *
 * 示例 2：
 *
 *
 * 输入：nums = [3,30,34,5,9]
 * 输出："9534330"
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [1]
 * 输出："1"
 *
 *
 * 示例 4：
 *
 *
 * 输入：nums = [10]
 * 输出："10"
*/

func largestNumber(nums []int) string {
	if len(nums) <= 1 {
		s := fmt.Sprintf("%v", nums)
		r := strings.ReplaceAll(strings.Trim(s, "[]"), " ", "")
		return r
	}

	var num []string
	count := 0
	for _, v := range nums {
		if v == 0 {
			count++
		}
		num = append(num, fmt.Sprintf("%v", v))
	}

	if count == len(nums) {
		return "0"
	}

	//fmt.Println(num)

	sort.Slice(num, func(i, j int) bool {
		num1 := num[i] + num[j]
		num2 := num[j] + num[i]
		if strings.Compare(num1, num2) > 0 { //大于交换
			return true
		}
		return false
	})

	//fmt.Println(num)

	s := fmt.Sprintf("%v", num)
	//fmt.Println(s)
	r := strings.ReplaceAll(strings.Trim(s, "[]"), " ", "")
	//fmt.Printf("r=%v\n", r)
	return r
}

func main() {
	//x := []int{3, 30, 34, 5, 9} //9534330
	//x := []int{10, 2} //210
	//x := []int{0, 0} //0
	//x := []int{0, 1} //10
	//x := []int{1, 0} //10
	fmt.Println(largestNumber([]int{0, 0})) //0
	fmt.Println(largestNumber([]int{0, 1})) //10
	fmt.Println(largestNumber([]int{1, 0})) //10
	fmt.Println(largestNumber([]int{10}))   //10

	fmt.Println(largestNumber([]int{8308, 8308, 830})) //"83088308830"
}
