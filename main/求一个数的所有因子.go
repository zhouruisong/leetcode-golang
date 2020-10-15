package main

import (
	"fmt"
	"sort"
)

/*
获取一个数的所有因子
*/
func yinzi(x int) []int {
	var ret []int
	for i := 1; i < x; i++ {
		if x%i == 0 {
			ret = append(ret, i)
		}
	}

	return ret
}

//效率较高
//1是任意数字的因子
//用数组按顺序存储因子，函数返回因子的个数
func factor(n int) int {
	count := 1
	arr := [100]int{1}
	for i := 2; i*i <= n; i++ {
		if (n % i) == 0 {
			arr[count] = i
			count++
			arr[count] = n / i
			count++
		}
	}

	//fmt.Printf("ret: %+v\n", arr[:count])
	sort.Ints(arr[:count])
	fmt.Printf("ret: %+v\n", arr[:count])
	return count
}

func main() {
	x := 87975123
	ret := yinzi(x)
	fmt.Printf("ret: %+v\n", ret)

	fmt.Printf("ret: %+v\n", factor(x))
}
