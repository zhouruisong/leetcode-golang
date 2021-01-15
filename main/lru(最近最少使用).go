package main

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
	head := &LinkNode{0,0,nil,nil}
	tail := &LinkNode{0,0,nil,nil}
	head.next = tail
	tail.pre = head

	return LRUCache{make(map[int]*LinkNode), capacity, head, tail}
}

func (this *LRUCache) Get(key int) int {
	cache := this.m

	if v,exist := cache[key];exist {
		this.MoveToHead(v)
		return v.val
	}else{
		return -1
	}
}


func (this *LRUCache) Put(key int, value int)  {
	tail := this.tail
	cache := this.m
	if v,exist := cache[key];exist {
		//更新值
		v.val = value
		//移动到最前端
		this.MoveToHead(v)
	}else{
		node := &LinkNode{key,value,nil,nil}
		//容量满了,删除最后的元素,更新map
		if len(cache)  == this.cap {
			delete(cache, tail.pre.key)
			tail.pre.pre.next = tail
			tail.pre = tail.pre.pre
		}

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
	//插入在首位
	head := this.head
	node.next = head.next
	node.pre = head
	head.next.pre = node
	head.next = node
}

func (this *LRUCache) MoveToHead(node *LinkNode) {
	this.RemoveNode(node)
	this.AddNode(node)
}
