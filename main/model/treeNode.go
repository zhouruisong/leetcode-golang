package model

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InitBinaryTree(root *TreeNode) *TreeNode {
	l := TreeNode{}
	r := TreeNode{}
	root.Left = l.NewBinaryTreeNode(4)
	root.Right = r.NewBinaryTreeNode(8)

	root.Left.Left = (&TreeNode{}).NewBinaryTreeNode(11)
	root.Left.Left.Left = (&TreeNode{}).NewBinaryTreeNode(7)
	root.Left.Left.Right = (&TreeNode{}).NewBinaryTreeNode(2)

	root.Right.Left = (&TreeNode{}).NewBinaryTreeNode(13)
	root.Right.Right = (&TreeNode{}).NewBinaryTreeNode(4)
	root.Right.Right.Right = (&TreeNode{}).NewBinaryTreeNode(1)
	root.Right.Right.Left = (&TreeNode{}).NewBinaryTreeNode(6)

	return root
}

func InitBinaryTree2(root *TreeNode) *TreeNode {
	l := TreeNode{}
	r := TreeNode{}
	root.Left = l.NewBinaryTreeNode(9)
	root.Right = r.NewBinaryTreeNode(20)

	root.Left.Left = (&TreeNode{}).NewBinaryTreeNode(1)
	root.Left.Right = (&TreeNode{}).NewBinaryTreeNode(2)

	root.Right.Left = (&TreeNode{}).NewBinaryTreeNode(15)
	root.Right.Right = (&TreeNode{}).NewBinaryTreeNode(7)

	return root
}

func InitBinaryTree3(root *TreeNode) *TreeNode {
	l := TreeNode{}
	r := TreeNode{}
	root.Left = l.NewBinaryTreeNode(2)
	root.Right = r.NewBinaryTreeNode(2)

	root.Right.Left = (&TreeNode{}).NewBinaryTreeNode(3)
	root.Right.Right = (&TreeNode{}).NewBinaryTreeNode(4)

	root.Left.Left = (&TreeNode{}).NewBinaryTreeNode(4)
	root.Left.Right = (&TreeNode{}).NewBinaryTreeNode(3)

	return root
}

func InitBinaryTree5(root *TreeNode) *TreeNode {
	l := TreeNode{}
	root.Left = l.NewBinaryTreeNode(2)
	return root
}

func InitBinaryTree4(root *TreeNode) *TreeNode {
	l := TreeNode{}
	r := TreeNode{}
	root.Left = l.NewBinaryTreeNode(2)
	root.Right = r.NewBinaryTreeNode(5)

	root.Right.Right = (&TreeNode{}).NewBinaryTreeNode(6)

	root.Left.Left = (&TreeNode{}).NewBinaryTreeNode(3)
	root.Left.Right = (&TreeNode{}).NewBinaryTreeNode(4)

	return root
}

func (this *TreeNode) NewBinaryTreeNode(value int) *TreeNode {
	this.Val = value
	this.Left = nil
	this.Right = nil
	return this
}

func GenTree(x []int) *TreeNode {
	index := 0
	return GenTreeHelp(x, &index)
}

func GenTreeHelp(x []int, index *int) *TreeNode {
	if len(x) == *index {
		return nil
	}
	if x[*index] == -1 {
		*index = *index + 1
		return nil
	}
	//fmt.Print(*index, x[*index])
	(*index) = *index + 1
	fmt.Println(x, x[*index])
	node := &TreeNode{Val: x[*index]}
	//(*index) = *index + 1

	node.Left = GenTreeHelp(x, index)
	return node
}

//	序列化
func DumpTreeToString(tree *TreeNode) string {
	var str string = ""
	dumpTreeToStringHelper(tree, &str)
	return str
}

func dumpTreeToStringHelper(tree *TreeNode, str *string) {
	if tree == nil {
		*str += "#"
		return
	}
	*str += strconv.Itoa(tree.Val)
	dumpTreeToStringHelper(tree.Left, str)
	dumpTreeToStringHelper(tree.Right, str)
}

//	反序列化
func ConFromPreStr(str string) *TreeNode {
	var strIndex int = 0
	return conFromPreStrHelper(str, &strIndex)
}

func conFromPreStrHelper(str string, index *int) *TreeNode {
	if len(str) == *index {
		return nil
	}
	if (str)[*index] == '#' {
		(*index) = *index + 1
		return nil
	}
	v, _ := strconv.Atoi(string((str)[*index]))
	root := &TreeNode{Val: v}
	(*index) = *index + 1
	root.Left = conFromPreStrHelper(str, index)
	root.Right = conFromPreStrHelper(str, index)

	return root
}
