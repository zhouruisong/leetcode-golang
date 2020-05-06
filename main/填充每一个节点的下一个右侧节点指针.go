package main

import "fmt"

/*
给定一个完美二叉树，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下：

struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。
初始状态下，所有 next 指针都被设置为 NULL。
*/

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

//bfs(breadth-first-search),广度遍历
func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	pre := root
	for pre.Left != nil {
		parent := pre
		for parent != nil {
			parent.Left.Next = parent.Right
			if parent.Next != nil {
				parent.Right.Next = parent.Next.Left
			}

			parent = parent.Next
		}

		pre = pre.Left
	}

	return root
}

//dfs(depth-first-search) 深度遍历
func connect2(root *Node) *Node {
	if root == nil {
		return nil
	}
	l := root.Left
	r := root.Right
	for l != nil {
		l.Next = r  //连接
		l = l.Right //往右
		r = r.Left  //往左
	}
	connect(root.Left)
	connect(root.Right)
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

	root.Left.Left = (&Node{}).NewNode(6)
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
