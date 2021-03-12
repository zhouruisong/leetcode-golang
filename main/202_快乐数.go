package main

import "fmt"

/*
编写一个算法来判断一个数 n 是不是快乐数。

「快乐数」定义为：

对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和。
然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。
如果 可以变为  1，那么这个数就是快乐数。
如果 n 是快乐数就返回 true ；不是，则返回 false 。

示例 1：

输入：19
输出：true
解释：
12 + 92 = 82
82 + 22 = 68
62 + 82 = 100
12 + 02 + 02 = 1
示例 2：

输入：n = 2
输出：false

提示：

1 <= n <= 231 - 1
*/

//o(logn)
func isHappy(n int) bool {
	hash := make(map[int]bool) // 开一个 map 判断是否循环
	for n != 1 {
		if _, ok := hash[n]; ok { // 出现了循环，证明不是快乐数
			return false
		}
		hash[n] = true
		n = step(n)
	}
	return true
}

//快慢指针o(logn)
func isHappy2(n int) bool {
	slow, fast := n, step(n)
	for fast != 1 && slow != fast {
		slow = step(slow)
		fast = step(step(fast))
	}
	return fast == 1
}

func step(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}
	return sum
}

func main() {
	fmt.Println(isHappy(19))
	fmt.Println(isHappy2(116))
}
