package main

import "fmt"

/*
给定一个二叉树

struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。
初始状态下，所有 next 指针都被设置为 NULL。
进阶：

你只能使用常量级额外空间。
使用递归解题也符合要求，本题中递归程序占用的栈空间不算做额外的空间复杂度。

*/
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	cur := root //指针
	for cur != nil {
		var pre *Node  //前置节点
		var down *Node //下降节点，节点为下一层的左边第一节点
		for cur != nil {
			if cur.Left != nil { //左节点判断
				if pre != nil {
					pre.Next = cur.Left //pre不为空 设置Next
				} else {
					down = cur.Left //pre为空 设置下降节点
				}
				pre = cur.Left //设置前置节点
			}

			if cur.Right != nil { //右节点判断，同上
				if pre != nil {
					pre.Next = cur.Right
				} else {
					down = cur.Right
				}
				pre = cur.Right
			}
			cur = cur.Next //同层移动
		}
		cur = down //下降
	}
	return root
}

func main() {
	root := &Node{
		Val: 1,
	}

	l := Node{}
	r := Node{}
	root.Left = l.NewNode(2)
	root.Right = r.NewNode(3)

	root.Right.Left = (&Node{}).NewNode(4)
	root.Right.Right = (&Node{}).NewNode(5)

	root.Left.Right = (&Node{}).NewNode(7)

	fmt.Println(connect(root))
}

func (this *Node) NewNode(value int) *Node {
	this.Val = value
	this.Left = nil
	this.Right = nil
	this.Next = nil
	return this
}
