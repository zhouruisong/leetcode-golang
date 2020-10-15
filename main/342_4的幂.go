package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
逆向操作，判断是不是4的幂，就看跟4的幂是否相等
比较n和4的幂的大小，若n < 1 (30) 则返回false，n ==1 返回true
n > 4 ,进入下次循环，比较 n < 4 (31)，依次循环
*/
func isPowerOfFour(n int) bool {
	aa := 1
	for {
		if n < aa {
			return false
		}

		if n == aa {
			return true
		}

		aa = aa * 4
	}
}

/*
方法二，
利用3进制字符串，判断第一位为1，其他为0，符合条件则为3的幂
if a < 1 {
fmt.Println("false")
}
aa := strconv.FormatInt(int64(a),4)
fmt.Println(aa[0:1] =="1" && strings.Count(aa,"0") == len(aa)-1)
这个思路更简洁
*/

func isPowerOfFour2(n int) bool {
	if n < 1 {
		return false
	}

	aa := strconv.FormatInt(int64(n), 4)
	return aa[0:1] == "1" && strings.Count(aa, "0") == len(aa)-1
}

func main() {
	fmt.Println(isPowerOfFour(16))
	fmt.Println(isPowerOfFour2(256))
}
