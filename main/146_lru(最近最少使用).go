package main

import "fmt"

/*
运用你所掌握的数据结构，设计和实现一个 LRU (最近最少使用) 缓存机制。它应该支持以下操作：获取数据 get 和 写入数据 put 。
获取数据 get(key) - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1 。

写入数据 put(key, value) - 如果密钥不存在，则写入其数据值。当缓存容量达到上限时，它应该在写入新数据之前删除最近最少使用的数据值，从而为新的数据值留出空间。

进阶:你是否可以在 O(1) 时间复杂度内完成这两种操作？

示例:

LRUCache cache = new LRUCache(2); // 缓存容量
cache.put(1, 1);
cache.put(2, 2);
cache.get(1);       // 返回  1
cache.put(3, 3);    // 该操作会使得密钥 2 作废
cache.get(2);       // 返回 -1 (未找到)
cache.put(4, 4);    // 该操作会使得密钥 1 作废
cache.get(1);       // 返回 -1 (未找到)
cache.get(3);       // 返回  3
cache.get(4);       // 返回  4
*/

//hashmap + 双链表
type LinkNode struct {
	key  int
	val  int
	pre  *LinkNode
	next *LinkNode
}

type LRUCache struct {
	m    map[int]*LinkNode
	cap  int
	head *LinkNode
	tail *LinkNode
}

func Constructor(capacity int) LRUCache {
	head := &LinkNode{0, 0, nil, nil}
	tail := &LinkNode{0, 0, nil, nil}
	head.next = tail
	tail.pre = head
	return LRUCache{make(map[int]*LinkNode), capacity, head, tail}
}

func (this *LRUCache) Get(key int) int {
	//查询map,有就返回,并且把查到的节点放到首部,未查到直接返回-1
	cache := this.m
	if v, exist := cache[key]; exist {
		this.MoveToHead(v)
		return v.val
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	head := this.head
	tail := this.tail
	cache := this.m
	//假若元素存在
	if v, exist := cache[key]; exist {
		//1.更新值
		v.val = value
		//2.移动到最前
		this.MoveToHead(v)
	} else {
		node := &LinkNode{key, value, nil, nil}
		//如果容量满了,删除尾部的
		if len(cache) == this.cap {
			//删除最后元素
			delete(cache, tail.pre.key)
			tail.pre.pre.next = tail
			tail.pre = tail.pre.pre
		}
		node.next = head.next
		node.pre = head
		head.next.pre = node
		head.next = node
		cache[key] = node
	}
}

func (this *LRUCache) RemoveNode(node *LinkNode) {
	//当前位置删除
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (this *LRUCache) AddNode(node *LinkNode) {
	//放到首位
	head := this.head
	node.next = head.next
	head.next.pre = node
	node.pre = head
	head.next = node
}

func (this *LRUCache) MoveToHead(node *LinkNode) {
	this.RemoveNode(node)
	this.AddNode(node)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	c := Constructor(2)
	c.Put(1, 1)
	c.Put(2, 2)
	fmt.Println(c.Get(1))
	fmt.Println(c.Get(2))
}
