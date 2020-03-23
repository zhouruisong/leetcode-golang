package main

import "fmt"

/*

输入：intersectVal = 8, listA = [4,1,8,4,5], listB = [5,0,1,8,4,5], skipA = 2, skipB = 3
输出：Reference of the node with value = 8
输入解释：相交节点的值为 8 （注意，如果两个列表相交则不能为 0）。从各自的表头开始算起，链表 A 为 [4,1,8,4,5]，链表 B 为 [5,0,1,8,4,5]。
在 A 中，相交节点前有 2 个节点；在 B 中，相交节点前有 3 个节点。

开两个指针分别遍历这两个链表，在第一次遍历到尾部的时候，指向另一个链表头部继续遍历，这样会抵消长度差。

如果链表有相交，那么会在中途相等，返回相交节点；

如果链表不相交，那么最后会 nil == nil，返回 nil；

代码

func getIntersectionNode(headA, headB *ListNode) *ListNode {
    curA,curB := headA,headB
    for curA != curB {
        if curA == nil {    // 如果第一次遍历到链表尾部，就指向另一个链表的头部，继续遍历，这样会抵消长度差。如果没有相交，因为遍历长度相等，最后会是 nil ==  nil
            curA = headB
        } else {
            curA = curA.Next
        }
        if curB == nil {
            curB = headA
        } else {
            curB = curB.Next
        }
    }
    return curA
}
 */

func main(){
	defer fmt.Printf("defer 1\n")
	defer fmt.Printf("defer 2\n")

	panic("error")
}