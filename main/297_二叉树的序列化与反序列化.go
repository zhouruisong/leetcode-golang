package main

import (
	"fmt"
	"main/model"
	"strconv"
	"strings"
)

/*
序列化是将一个数据结构或者对象转换为连续的比特位的操作，进而可以将转换后的数据存储在一个文件或者内存中，
同时也可以通过网络传输到另一个计算机环境，采取相反方式重构得到原数据。

请设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 / 反序列化算法执行逻辑，
你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。

示例:

你可以将以下二叉树：

    1
   / \
  2   3
     / \
    4   5

序列化为 "[1,2,3,null,null,4,5]"
*/

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *model.TreeNode) string {
	if root == nil {
		return "X"
	}

	return strconv.Itoa(root.Val) + "," + this.serialize(root.Left) + "," + this.serialize(root.Right)
}

func buildTree(list *[]string) *model.TreeNode {
	rootVal := (*list)[0]
	*list = (*list)[1:]
	if rootVal == "X" {
		return nil
	}

	Val, _ := strconv.Atoi(rootVal)
	root := &model.TreeNode{Val: Val}
	root.Left = buildTree(list)
	root.Right = buildTree(list)
	return root
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *model.TreeNode {
	list := strings.Split(data, ",")
	return buildTree(&list)
}

func main() {
	root := &model.TreeNode{Val: 1}
	root.Left = &model.TreeNode{Val: 2}
	root.Right = &model.TreeNode{Val: 3}

	root.Right.Left = &model.TreeNode{Val: 4}
	root.Right.Right = &model.TreeNode{Val: 5}
	c := Constructor()
	str := c.serialize(root)
	fmt.Println(str)
	root = c.deserialize(str)
	fmt.Println(model.DumpTreeToString(root))
}
