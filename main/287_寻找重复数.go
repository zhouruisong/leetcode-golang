package main

/*
综上
1.数组中有一个重复的整数 <==> 链表中存在环
2.找到数组中的重复整数 <==> 找到链表的环入口

至此，问题转换为 142 题。那么针对此题，快、慢指针该如何走呢。根据上述数组转链表的映射关系，可推出
142 题中慢指针走一步 slow = slow.next ==> 本题 slow = nums[slow]
142 题中快指针走两步 fast = fast.next.next ==> 本题 fast = nums[nums[fast]]
*/

func findDuplicate(nums []int) int {
	//采用寻找循环链表的方式,快慢指针
	slow, fast := 0, 0
	slow = nums[slow]
	fast = nums[nums[fast]]

	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	//slow为环形链表的入口位置
	pre1 := 0
	pre2 := slow

	for pre1 != pre2 {
		pre1 = nums[pre1]
		pre2 = nums[pre2]
	}

	return pre1
}
