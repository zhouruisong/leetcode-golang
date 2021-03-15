package model

import (
	"math/rand"
	"fmt"
)

type LinkNode struct {
	Pre  *LinkNode
	Next *LinkNode
	Val  int
}

type ListNode struct {
	Next *ListNode
	Val  int
}

func PrintList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, "->")
		head = head.Next
	}
	fmt.Println()
}

//	更具数组反序列化链表
func UnmarshalListBySlice(nums []int) *ListNode {
	head := &ListNode{Val: -1, Next: nil}
	tmp := head
	for _, v := range nums {
		tmp.Next = &ListNode{Val: v, Next: nil}
		tmp = tmp.Next
	}
	return head.Next
}

// 随机初始化链表
func UnmarshalListByRand(max_num int, len int) *ListNode {
	head := &ListNode{Val: -1, Next: nil}
	tmp := head

	for i := 0; i < len; i++ {
		tmp.Next = &ListNode{Val: rand.Intn(max_num), Next: nil}
		tmp = tmp.Next
	}
	return head.Next
}
