package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
输入一个正整数数组，把数组里所有数字拼接起来排成一个数，打印能拼接出的所有数字中最小的一个。例如输入数组{3，32，321}，则打印出这三个数字能排成的最小数字为321323
*/

func printMin(x []string) int {
	var num []string
	for _, v := range x {
		num = append(num, v)
	}

	fmt.Printf("num: %+v\n", num)

	ln := len(num)
	for i := 0; i < ln; i++ {
		sort.Slice(num, func(i, j int) bool {
			num1 := num[i] + num[j]
			num2 := num[j] + num[i]
			if strings.Compare(num1, num2) < 0 {
				return true
			}
			return false
		})
	}

	fmt.Printf("num: %+v\n", num)

	for i := 0; i < ln; i++ {
		fmt.Printf("%+v", num[i])
	}
	fmt.Print("\n")
	s := fmt.Sprintf("%v", num)
	fmt.Println(s)
	r := strings.ReplaceAll(strings.Trim(s, "[]"), " ", "")
	fmt.Printf("r=%v\n", r)
	//ret, _ := strconv.Atoi(r)

	//字符串转数字
	var n uint64
	for _, c := range []byte(r) {
		d := c - '0'
		n *= uint64(10)
		n1 := n + uint64(d)
		n = n1
		fmt.Println(n1)
	}

	return int(n)
}

func main() {
	x := []string{"3", "32", "321"}
	fmt.Println(printMin(x))

	arr := []string{"123", "456", "789"}
	s := fmt.Sprintf("%v", arr)
	fmt.Println(s)

	//[123 456 789] -> 123456789
	str := strings.ReplaceAll(strings.Trim(s, "[]"), " ", "")
	fmt.Println(str)

	//字符串转成int64
	var n uint64
	var n1 uint64
	for _, c := range []byte(str) {
		d := c - '0'
		n *= uint64(10)
		n1 = n + uint64(d)
		n = n1
	}

	fmt.Println(n)
}
