package main

import (
	"fmt"
	"main/model"
)

/*
 * type ListNode struct{
 *   Val int
 *   Next *ListNode
 * }
 */

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param pHead ListNode类
 * @param k int整型
 * @return ListNode类
 */

func FindKthToTail(pHead *model.ListNode, k int) *model.ListNode {
	// write code here
	if k <= 0 {
		return nil
	}

	//快慢指针
	if pHead == nil {
		return nil
	}

	first := pHead
	for i := 0; i < k; i++ {
		if first == nil {
			return nil
		}
		first = first.Next
	}

	second := pHead
	for first != nil {
		first = first.Next
		second = second.Next
	}
	return second
}

func main() {
	v1 := []int{1, 2, 3, 4, 5}
	head := model.UnmarshalListBySlice(v1)
	model.PrintList(head)
	res := FindKthToTail(head, 2)
	fmt.Println(res.Val)
}
