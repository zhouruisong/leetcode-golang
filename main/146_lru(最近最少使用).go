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
		//插入首位
		this.AddNode(node)
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
 * lru design
 * @param operators int整型二维数组 the ops
 * @param k int整型 the k
 * @return int整型一维数组
 */

func LRU2(operators [][]int, k int) []int {
	// write code here
	//[[1,1,1],[1,2,2],[1,3,2],[2,1],[1,4,4],[2,2]],3
	res := make([]int, 0, len(operators))
	keys := make([]int, k)
	values := make([]int, k)

	for _, v := range operators {
		if v[0] == 1 {
			//插入
			//超过容量,第一个删除
			if len(keys) == k {
				keys = keys[1:]
				values = values[1:]
			}
			keys = append(keys, v[1])
			values = append(values, v[2])
		} else if v[0] == 2 {
			//获取
			index := -1
			for i := 0; i < len(keys); i++ {
				if v[1] == keys[i] {
					index = i
					break
				}
			}
			//未查到
			if index == -1 {
				res = append(res, -1)
			} else {
				res = append(res, values[index])
				//将获取的index对应的key和alue 放到队列尾部
				if index < k-1 {
					keys = append(keys[:index], append(keys[index+1:], keys[index])...)
					values = append(values[:index], append(values[index+1:], values[index])...)
				}
			}
		}
	}

	return res
}

func main() {
	c := Constructor(2)
	c.Put(1, 1)
	c.Put(2, 2)
	fmt.Println(c.Get(1))
	fmt.Println(c.Get(2))
}
