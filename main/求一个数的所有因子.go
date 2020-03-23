package main

import "fmt"

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

func main() {
	x := 28
	ret := yinzi(x)
	fmt.Printf("ret: %+v\n", ret)
}
