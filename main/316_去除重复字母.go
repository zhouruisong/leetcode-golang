package main

import (
	"container/list"
	"fmt"
)

/*
给你一个仅包含小写字母的字符串，请你去除字符串中重复的字母，使得每个字母只出现一次。需保证返回结果的字典序最小（要求不能打乱其他字符的相对位置）

示例 1:

输入: "bcabc"
输出: "abc"
示例 2:

输入: "cbacdcbc"
输出: "acdb"
*/

func removeDuplicateLetters(s string) string {
	if len(s) == 0 {
		return ""
	}

	st := list.New()

	//记录栈中是否存在该元素
	mp := make(map[uint8]bool, 256)

	//记录每个元素出现的次数
	count := make(map[uint8]int, 256)
	for i:=0;i<len(s);i++ {
		count[s[i]]++
	}

	for i := 0; i < len(s); i++ {
		//每遍历一个字符 计数-1
		count[s[i]]--

		if mp[s[i]] {
			continue
		}

		//要入栈的字符比栈顶的小,先弹栈
		for st.Len() > 0 && st.Front().Value.(uint8) > s[i] {
			// fmt.Print(s[i])

			//s[i]后面不存在栈顶元素了 break
			if count[st.Front().Value.(uint8)] == 0 {
				break
			}

			top := st.Remove(st.Front()).(uint8)
			mp[top] = false
		}

		st.PushFront(s[i])
		mp[s[i]] = true
	}

	var ret []byte

	for st.Len() > 0 {
		node := st.Remove(st.Front()).(uint8)
		ret = append(ret, node)
	}

	//交换顺序
	for left, right := 0, len(ret)-1; left < right; left, right = left+1, right-1 {
		ret[left], ret[right] = ret[right], ret[left]
	}

	return string(ret)
}

func main() {
	fmt.Print(removeDuplicateLetters("cbacdcbc"))
}
