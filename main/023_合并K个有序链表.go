package main

import "main/model"

/*
合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。

示例:

输入:
[
  1->4->5,
  1->3->4,
  2->6
]
输出: 1->1->2->3->4->4->5->6
*/

func mergeList(list1 *model.ListNode, list2 *model.ListNode) *model.ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	p1 := list1
	p2 := list2
	head := p2
	if p1.Val <= p2.Val {
		head = p1
		p1 = p1.Next
	} else {
		head = p2
		p2 = p2.Next
	}

	tmp := head
	for p1 != nil && p2 != nil {
		if p1.Val <= p2.Val {
			head.Next = p1
			p1 = p1.Next
		} else {
			head.Next = p2
			p2 = p2.Next
		}
		head = head.Next
	}

	if p1 != nil {
		head.Next = p1
	}

	if p2 != nil {
		head.Next = p2
	}

	return tmp
}

func mergeKLists(lists []*model.ListNode) *model.ListNode {
	n := len(lists)
	if n == 0 {
		return nil
	}

	head := lists[0]
	//tmp := head
	for i := 1; i < n; i++ {
		head = mergeList(head, lists[i])
	}

	return head
}

/*
[1,4,5],[1,3,4],[2,6]
*/
func main() {
	v1 := []int{}
	v2 := []int{1}
	//v3 := []int{2, 6}
	//v1 := []int{1, 4, 5}
	//v2 := []int{2, 4, 5, 6, 8}
	//v3 := []int{3, 7, 8, 9, 11}
	//v4 := []int{5, 7, 8, 16, 18}
	//v5 := []int{12, 14, 15, 16, 18}
	//v6 := []int{22, 24, 25, 26, 28}

	tmp1 := model.UnmarshalListBySlice(v1)
	tmp2 := model.UnmarshalListBySlice(v2)
	//tmp3 := model.UnmarshalListBySlice(v3)
	//tmp4 := model.UnmarshalListBySlice(v4)
	//tmp5 := model.UnmarshalListBySlice(v5)
	//tmp6 := model.UnmarshalListBySlice(v6)

	//model.PrintList(tmp1)
	//model.PrintList(tmp2)
	//model.PrintList(tmp3)
	//model.PrintList(tmp4)
	//model.PrintList(tmp5)
	//model.PrintList(tmp6)

	var lists []*model.ListNode
	lists = append(lists, tmp1)
	lists = append(lists, tmp2)
	//lists = append(lists, tmp3)
	//lists = append(lists, tmp4)
	//lists = append(lists, tmp5)
	//lists = append(lists, tmp6)

	t := mergeKLists(lists)
	model.PrintList(t)
}
