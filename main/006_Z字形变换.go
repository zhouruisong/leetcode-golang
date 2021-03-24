package main

import (
	"fmt"
	"strings"
)

/*
将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。

比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：

P   A   H   N
A P L S I I G
Y   I   R
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。

请你实现这个将字符串进行指定行数变换的函数：

string convert(string s, int numRows);


示例 1：

输入：s = "PAYPALISHIRING", numRows = 3
输出："PAHNAPLSIIGYIR"
示例 2：
输入：s = "PAYPALISHIRING", numRows = 4
输出："PINALSIGYAHRPI"
解释：
P     I    N
A   L S  I G
Y A   H R
P     I
示例 3：

输入：s = "A", numRows = 1
输出："A"
*/

/*
其实这个题官方的第一种方法已经很秀了，在我看来是最舒服的实现，下面是我的golang版本实现，z字变换，
其实说白了就是让字符串在第0-给定的行数之间摇摆行走，我们可以根据画出的二维数组拼接，就得到了结果，
关键在变换上，通过循环字符串的每个字符，判断它应该放在哪一个数组中，我们模仿z字变换轨迹，从第0个数组开始插入，
当写到最大行数时，触底使其反弹，从相反的方向开始写入，依次执行这个步骤，正-反-正-反，就得到了结果

链接：https://leetcode-cn.com/problems/zigzag-conversion/solution/golangban-shi-xian-jia-tong-su-jie-xi-by-user1775a/

numRows = 3
tmp[0]:LCIR
tmp[1]:ETOESIIG
tmp[2]:EDHN
来回上下扫，所以tmp[0]先放L，tmp[1]放E，tmp[2]放E
这时候，扫到底了，往上扫，tmp[1]放T，tmp[0]放C
这时候，扫到顶了，往下扫，tmp[1]放O，tmp[2]放D
这时候，扫到底了，继续。。。
*/
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	res := make([]string, numRows)
	numRows--
	index := 0
	sort := true
	for i, _ := range s {
		res[index] += string(s[i])
		if index == 0 {
			sort = true
		} else if index == numRows {
			sort = false
		}

		//从上倒下
		if sort {
			index++
		} else { //跳到指定行
			index--
		}
	}
	return strings.Join(res, "")
}

func main() {
	s := "PAYPALISHIRING"
	numRows := 3
	fmt.Println(convert(s, numRows)) //PAHNAPLSIIGYIR
}
