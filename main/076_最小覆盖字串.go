package main

import "fmt"

/*
给你一个字符串 S、一个字符串 T 。请你设计一种算法，可以在 O(n) 的时间复杂度内，从字符串 S 里面找出：包含 T 所有字符的最小子串。

示例：
输入：S = "ADOBECODEBANC", T = "ABC"
输出："BANC"

提示：
如果 S 中不存这样的子串，则返回空字符串 ""。
如果 S 中存在这样的子串，我们保证它是唯一的答案。
*/

/*
使用两个哈希表，一个记录所需要的字符还需要多少，另一个记录当前滑动窗口中各字符的数量有多少，count记录滑动窗口中已经包含了tt中多少字符。
开始左指针不动，右指针向右移动，直到窗口中包含了tt中全部字符，开始移动左指针收缩，收缩到窗口中刚好不满足tt中的字符为止。如此循环，
一旦count == t.length说明当前是满足条件的，记录此时的大小和字符串
*/

func minWindow(s string, t string) string {
	//	s := "ADOBECODEBANC"
	//	t := "ABC"
	need := make(map[uint8]int)
	tt := len(t)
	for i := 0; i < tt; i++ {
		need[t[i]]++
	}

	window := make(map[uint8]int)
	count := 0
	l, r := 0, 0
	n := len(s)
	min := n
	ans := ""

	for r < n {
		c := s[r]
		window[c]++

		//计算windows中是否包括了need的所有元素
		if window[c] <= need[c] {
			count++
		}

		//
		for count == tt {
			//
			if r-l+1 <= min {
				min = r - l + 1
				ans = s[l : r+1]
				//fmt.Println(ans)
			}

			window[s[l]]--
			if window[s[l]] < need[s[l]] {
				//fmt.Printf("%c\n",s[l])
				count--
			}

			l++
			if l >= n {
				break
			}
		}

		r++
	}

	fmt.Println(min)
	return ans

}

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	fmt.Print(minWindow(s, t))
}
