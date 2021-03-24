package main

import "fmt"

/*
* 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
*
* 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
*
*
*
* 示例:
*
* 输入："23"
* 输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
*
*
* 说明:
* 尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。
*
 */

/*
解法思路很简单，就是来一个数字就拼接一次
例如 "23" 这个例子
初始化 结果为 ""
先取出 2 对应的列表 [a b c], 用 "" 和 a b c 分别拼接并加入结果 得到 [a b c]
再取出 3 对应的列表 [d e f], 分别用a b c 和 d e f 拼接并加入结果得到[ad ae af bd be bf cd ce cf]
*/

var num2string map[byte][]string = map[byte][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	res := []string{}
	if digits == "" {
		return res
	}

	res = append(res, "")
	for i := 0; i < len(digits); i++ {
		nums := num2string[digits[i]]
		pre := res
		res = []string{}
		for _, num := range nums {
			for _, v := range pre {
				res = append(res, v+num)
			}
		}
	}

	return res
}

/*广度优先搜索
比起深度优先更加简单和高效，几乎无中间内存使用。

执行用时 : 0 ms, 在所有 golang 提交中击败了100.00%的用户
内存消耗 : 2 MB, 在所有 golang 提交中击败了100.00%的用户
*/

var table []string = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func letterCombinations2(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	result := []string{""}

	for i := 0; i < len(digits); i++ {
		t := table[digits[i]-'0']
		temp := []string{}
		for j := 0; j < len(t); j++ {
			for z := 0; z < len(result); z++ {
				temp = append(temp, result[z]+string([]byte{t[j]}))
			}
		}
		result = temp
	}

	return result
}

func main() {
	x := "23"
	//fmt.Println(letterCombinations(x))
	fmt.Println(letterCombinations2(x))

}
