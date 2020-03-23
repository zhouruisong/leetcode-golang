package main

import "fmt"

/*
输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
*/

func main() {
	v1 := []int{7, 7, 9}
	var l1 *ListNode
	var tmp *ListNode
	tmp = nil
	lv1 := len(v1)
	for i := lv1 - 1; i >= 0; i-- {
		l1 = new(ListNode)
		l1.Val = v1[i]
		l1.Next = tmp
		tmp = l1
	}

	l1 = tmp
	for l1 != nil {
		fmt.Print(l1.Val)
		l1 = l1.Next
	}

	fmt.Println("=======")
	v2 := []int{5, 6, 4}

	var l2 *ListNode
	var tmp2 *ListNode
	tmp2 = nil
	lv2 := len(v2)
	for i := lv2 - 1; i >= 0; i-- {
		l2 = new(ListNode)
		l2.Val = v2[i]
		l2.Next = tmp2
		tmp2 = l2
	}

	l2 = tmp2
	for l2 != nil {
		fmt.Print(l2.Val)
		l2 = l2.Next
	}

	fmt.Println("\n=======")

	ret := addTwoNumbers(tmp, tmp2)
	for ret != nil {
		fmt.Println(ret.Val)
		ret = ret.Next
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := new(ListNode)
	head := result
	sum := 0 //进位
	tmp := 0
	v := 0 //本位数
	for l1 != nil && l2 != nil {
		tmp = sum + l1.Val + l2.Val //对位累加后再加上后面的进位
		v = tmp % 10                //本位数
		//fmt.Println(tmp)
		sum = tmp / 10 //进位
		node := new(ListNode)
		node.Val = v
		node.Next = nil
		result.Next = node
		result = result.Next
		l1 = l1.Next
		l2 = l2.Next
	}

	//不相等链表
	if l1 == nil && l2 != nil {
		for l2 != nil {
			tmp = sum + l2.Val
			v = tmp % 10
			sum = tmp / 10
			node := new(ListNode)
			node.Val = v
			node.Next = nil
			result.Next = node
			result = result.Next
			l2 = l2.Next
		}
	}

	if l1 != nil && l2 == nil {
		for l1 != nil {
			tmp = sum + l1.Val
			v = tmp % 10
			sum = tmp / 10
			node := new(ListNode)
			node.Val = v
			node.Next = nil
			result.Next = node
			result = result.Next
			l1 = l1.Next
		}
	}

	fmt.Printf("sum: %v\n", sum)

	//最高位是1需要进位
	if sum > 0 {
		node := new(ListNode)
		node.Val = sum
		result.Next = node
	}

	return head.Next
}
