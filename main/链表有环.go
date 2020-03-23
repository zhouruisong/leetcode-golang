package main

/*
方法一：hash表（8ms）

将遍历过的节点记录下来，如果又遍历到了，表示链表有环，时间复杂度O(n)，空间复杂度O(n)

func hasCycle(head *ListNode) bool {    // hash表
    hash := make(map[*ListNode]int)     // 开一个哈希表记录该节点是否已经遍历过，值记录节点索引
    for head != nil {
        if _,ok := hash[head]; ok {     // 该节点遍历过，形成了环
            return true
        }
        hash[head] = head.Val           // 记录该节点已经遍历过
        head = head.Next
    }
    return false
}
方法二：快慢指针（12ms）

快指针一次走两步，慢指针一次走一步，如果链表有环，那么两个指针始终会相遇。

时间复杂度 O(n)，空间复杂度 O(1)

func hasCycle(head *ListNode) bool {    // 快慢指针。假如爱有天意，那么快慢指针终会相遇
    if head == nil {
        return false
    }
    fastHead := head.Next       // 快指针，每次走两步
    for fastHead != nil && head != nil && fastHead.Next != nil {
        if fastHead == head {   // 快慢指针相遇，表示有环
            return true
        }
        fastHead = fastHead.Next.Next
        head = head.Next        // 慢指针，每次走一步
    }
    return false
}
方法三：走自己的路让别人无路可走思路（8ms）

在评论区看到的代码，表示差点被笑死。

将遍历过的节点值修改为一个后台不太可能设置的值，来标志这个节点我们遍历到了。

这个思路虽然有点沙雕，但是一样能达到 O(1) 的空间复杂度，2333

func hasCycle(head *ListNode) bool {    // 走自己的路让别人无路可走思路
    for head != nil {
        if head.Val == 18464187 {       // 这是自己走过的路，说明有环
            return true
        }
        head.Val = 18464187
        head = head.Next
    }
    return false
}
 */