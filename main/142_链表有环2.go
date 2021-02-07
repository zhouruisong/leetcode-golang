package main

import (
	"fmt"
	"main/model"
)

/*
* 给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
 *
 * 为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。
 *
 * 说明：不允许修改给定的链表。
 *
 *
 *
 * 示例 1：
 *
 * 输入：head = [3,2,0,-4], pos = 1
 * 输出：tail connects to node index 1
 * 解释：链表中有一个环，其尾部连接到第二个节点。
 *
 *
 *
 *
 * 示例 2：
 *
 * 输入：head = [1,2], pos = 0
 * 输出：tail connects to node index 0
 * 解释：链表中有一个环，其尾部连接到第一个节点。
 *
 *
 *
 *
 * 示例 3：
 *
 * 输入：head = [1], pos = -1
 * 输出：no cycle
 * 解释：链表中没有环。
你是否可以不用额外空间解决此题？
*/

//时间复杂度O(n)，空间复杂度O(n)
func detectCycle(head *model.ListNode) *model.ListNode {
	pos := -1
	//hash方式，key是每个节点的地址，value是节点位置
	hashMap := make(map[*model.ListNode]int)
	loc := 0
	for head != nil {
		if val, ok := hashMap[head]; ok {
			pos = val
			break
		}

		hashMap[head] = loc
		loc++
		head = head.Next
	}

	if pos >= 0 {
		fmt.Print("tail connects to node index ", pos)
	} else {
		fmt.Print("no cycle")
	}
	return head
}

//时间复杂度 O(n)，空间复杂度 O(1)
func detectCycle2(head *model.ListNode) *model.ListNode {
	if head == nil {
		return nil
	}

	//找到相遇位置
	hasCycleLoc := hasCycle(head)
	if hasCycleLoc == nil {
		return nil
	}

	p1 := head
	p2 := hasCycleLoc
	for p1 != p2 {
		p1 = p1.Next
		p2 = p2.Next
	}

	return p1
}

func hasCycle(head *model.ListNode) *model.ListNode { // 快慢指针。假如爱有天意，那么快慢指针终会相遇
	if head == nil {
		return nil
	}

	fastHead := head.Next // 快指针，每次走两步
	for fastHead != nil && head != nil && fastHead.Next != nil {
		if fastHead == head { // 快慢指针相遇，表示有环
			return fastHead
		}
		fastHead = fastHead.Next.Next
		head = head.Next // 慢指针，每次走一步
	}
	return nil
}

func detectCycle3(head *model.ListNode) *model.ListNode {
	// write code here
	if head == nil || head.Next == nil {
		return nil
	}

	slow, fast := head, head
	isCycle := false
	for fast.Next != nil && fast.Next.Next != nil {
		slow := slow.Next
		fast = fast.Next.Next
		if slow == fast {
			isCycle = true
			break
		}
	}

	if isCycle {
		fast = head
		for slow != fast {
			fast = fast.Next
			slow = slow.Next
		}
		return fast
	}
	return nil
}

func main() {
	v6 := []int{3, 2, 0, -4}
	tmp1 := model.UnmarshalListBySlice(v6)
	model.PrintList(tmp1)
	detectCycle3(tmp1)
}
