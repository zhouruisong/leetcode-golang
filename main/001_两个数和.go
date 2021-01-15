package main

/*题目：给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

示例: 给定 nums = [2, 7, 11, 15], target = 9 因为 nums[0] + nums[1] = 2 + 7 = 9 所以返回 [0, 1]*/

import (
	"fmt"
)

type hashNode struct {
	Id    int
	Index int
}

func twoSum(input []int, target int) map[int]int {
	m := make(map[int]hashNode)
	ret := make(map[int]int, 2)
	for k, v := range input {
		other := target - v
		val, ok := m[other]
		if !ok {
			h := hashNode{
				Id:    v,
				Index: k,
			}
			m[v] = h
			continue
		}

		ret[0] = k
		ret[1] = val.Index

		return ret
	}

	return ret
}

type Kv struct {
	Key int
	Val int
}

func twoSum2(nums []int, target int) []int {
	var result []int
	mp := make(map[int]Kv)
	for k, v := range nums {
		other := target - v
		val, ok := mp[other]
		if ok {
			result = append(result, k)
			result = append(result, val.Val)
			break
		} else {
			t := Kv{
				Key: v,
				Val: k,
			}
			mp[v] = t
		}
	}

	return result
}

func twoSum3(nums []int, target int) []int {
	mp := make(map[int]int)
	for k, v := range nums {
		other := target - v
		if j, ok := mp[other]; ok {
			return []int{j, k}
		} else {
			mp[v] = k
		}
	}

	return []int{}
}

func main() {
	input := []int{2, 7, 11, 15}
	target := 22
	//ret := twoSum(input, target)
	//for _, v := range ret {
	//	fmt.Printf("ret=%+v\n", v)
	//}

	ret2 := twoSum3(input, target)
	for _, v := range ret2 {
		fmt.Printf("ret2=%+v\n", v)
	}
}
