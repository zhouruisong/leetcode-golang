package main

import "main/model"

func reverseLink(head *model.ListNode) *model.ListNode {
	cur := head
	var tmp *model.ListNode
	for cur != nil {
		next := cur.Next
		cur.Next = tmp
		tmp = cur
		cur = next
	}

	return tmp
}

func main() {
	v6 := []int{1, 2, 3, 4, 5, 6, 7}
	tmp1 := model.UnmarshalListBySlice(v6)
	model.PrintList(tmp1)
	tmp2 := reverseLink(tmp1)
	model.PrintList(tmp2)
}
