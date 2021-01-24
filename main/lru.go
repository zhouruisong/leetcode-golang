package main

import "fmt"

//hashmap+双链表

type LinkNode2 struct {
	key  int
	val  int
	pre  *LinkNode2
	next *LinkNode2
}

type Cache struct {
	m    map[int]*LinkNode2
	cap  int
	head *LinkNode2
	tail *LinkNode2
}

func Constructor2(capacity int) Cache {
	head := &LinkNode2{0, 0, nil, nil}
	tail := &LinkNode2{0, 0, nil, nil}
	head.next = tail
	tail.pre = head

	return Cache{make(map[int]*LinkNode2), capacity, head, tail}
}

func (this *Cache) Get(key int) int {
	m := this.m

	if v, exist := m[key]; exist {
		this.MoveToHead(v)
		return v.val
	} else {
		return -1
	}
}

func (this *Cache) Put(key int, value int) {
	m := this.m

	if val, exist := m[key]; exist {
		val.val = value
		this.MoveToHead(val)
	} else {
		node := &LinkNode2{
			key: key,
			val: value,
		}
		//删除最后一个元素
		if len(m) == this.cap {
			tail := this.tail
			delete(m, tail.pre.key)
			tail.pre.pre.next = tail
			tail.pre = tail.pre.pre
		}

		this.Add(node)
		m[key] = node
	}
}

func (this *Cache) Add(node *LinkNode2) {
	//插入首位
	head := this.head

	node.next = head.next
	node.pre = head

	head.next.pre = node
	head.next = node
}

func (this *Cache) RemoveLastNode(node *LinkNode2) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (this *Cache) MoveToHead(node *LinkNode2) {
	this.RemoveLastNode(node)
	this.Add(node)
}

func main() {
	cache := Constructor2(2)
	cache.Put(1, 1)
	cache.Put(2, 2)

	fmt.Println(cache.Get(1)) // 返回  1
	cache.Put(3, 3) // 该操作会使得密钥 2 作废

	fmt.Println(cache.Get(2)) // 返回 -1 (未找到)

	cache.Put(4, 4)           // 该操作会使得密钥 1 作废
	fmt.Println(cache.Get(1)) // 返回 -1 (未找到)
	fmt.Println(cache.Get(3)) // 返回  3
	fmt.Println(cache.Get(4)) // 返回  4
}
