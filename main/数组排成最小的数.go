package main

import (
	"sort"
	"strings"
	"fmt"
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
	return 0
}

func main() {
	x := []string{"3", "32", "321"}
	printMin(x)
}
