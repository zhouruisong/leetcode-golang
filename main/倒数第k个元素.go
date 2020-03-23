package main

import (
	"fmt"
	"practies/model"
)

//倒数第K个元素,两个指针p1,p2，p1先走n-k个长度，然后p2指向开始位置，p1继续走到链表尾部，p2这时候所指的位置就是答案
func kThNumber(head *model.ListNode, k int) int {
	if head == nil || k <= 0 {
		return -1
	}

	p1 := head
	p2 := head

	for p1 != nil && k-1 > 0 {
		p1 = p1.Next
		k--
	}

	//超出范围
	if k-1 < 0 {
		return -1
	}

	if p1 == nil {
		return -1
	}

	for p1.Next != nil {
		p1 = p1.Next
		p2 = p2.Next
	}

	return p2.Val
}

func main() {
	v1 := []int{1, 2, 5, 3, 6}
	head := model.UnmarshalListBySlice(v1)
	model.PrintList(head)
	fmt.Println(kThNumber(head, 2))
}
