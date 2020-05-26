package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
* 给定一个整数，写一个函数来判断它是否是 3 的幂次方。
 *
 * 示例 1:
 *
 * 输入: 27
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入: 0
 * 输出: false
 *
 * 示例 3:
 *
 * 输入: 9
 * 输出: true
 *
 * 示例 4:
 *
 * 输入: 45
 * 输出: false
 *
 * 进阶：
 * 你能不使用循环或者递归来完成本题吗？
*/

/*
逆向操作，判断是不是3的幂，就看跟3的幂是否相等
比较n和3的幂的大小，若n < 1 (30) 则返回false，n ==1 返回true
n > 3 ,进入下次循环，比较 n < 3 (31)，依次循环
*/
func isPowerOfThree(n int) bool {
	aa := 1
	for {
		if n < aa {
			return false
		}

		if n == aa {
			return true
		}

		aa = aa * 3
	}
}

/*
方法二，
利用3进制字符串，判断第一位为1，其他为0，符合条件则为3的幂
if a < 1 {
fmt.Println("false")
}
aa := strconv.FormatInt(int64(a),3)
fmt.Println(aa[0:1] =="1" && strings.Count(aa,"0") == len(aa)-1)
这个思路更简洁
*/

func isPowerOfThree2(n int) bool {
	if n < 1 {
		return false
	}

	aa := strconv.FormatInt(int64(n), 3)
	return aa[0:1] == "1" && strings.Count(aa, "0") == len(aa)-1
}

func main() {
	fmt.Println(isPowerOfThree(26))
	fmt.Println(isPowerOfThree2(26))
}
