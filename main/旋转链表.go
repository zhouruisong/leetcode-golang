package main

/*给定一个链表，旋转链表，将链表每个节点向右移动 k 个位置，其中 k 是非负数。

示例 1:

输入: 1->2->3->4->5->NULL, k = 2
输出: 4->5->1->2->3->NULL
解释:
向右旋转 1 步: 5->1->2->3->4->NULL
向右旋转 2 步: 4->5->1->2->3->NULL
示例 2:

输入: 0->1->2->NULL, k = 4
输出: 2->0->1->NULL
解释:
向右旋转 1 步: 2->0->1->NULL
向右旋转 2 步: 1->2->0->NULL
向右旋转 3 步: 0->1->2->NULL
向右旋转 4 步: 2->0->1->NULL

解题思路:
1.遍历链表，保存长度
2.找到分割点，然后构建循环链表
3.在分割点分割，变单链表

ListNode* rotateRight(ListNode* head, int k) {
	if (!head)return NULL;
	ListNode *p1 = head;
	int len = 1;
	while (p1->next) {
		p1 = p1->next;
		len++;
	}
	p1->next = head;
	k %= len;
	for (int i = 1; i < len - k + 1; i++) {
		p1 = p1->next;
		head = head->next;
	}
	p1->next = NULL;

	return head;
}

*/

func main() {

}
