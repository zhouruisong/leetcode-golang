package main

import (
	"fmt"
	"math"
)

//计算小于等于N的所有素数
func sushu(N uint64, smp map[uint64]uint64) {
	//fmt.Println(i)
	k := uint64(2)
	for i := uint64(2); i <= N; i++ {
		//t := (uint64)(math.Sqrt(float64(i))
		t := (uint64)(math.Sqrt(float64(i)))
		for ; k <= t; k++ {
			//fmt.Println(k)
			if i%k == 0 {
				break
			}
		}

		//是素数直接返回
		if k > t {
			//fmt.Println(i)
			smp[i] = i
		}

		k = uint64(2)
	}

	for _, v := range smp {
		fmt.Println(v)
	}
}

func primePalindrome(N int) int {
	y := uint64(0)
	tmp := uint64(0)
	limit := uint64(10 * 10)

	for i := uint64(N); i <= limit; i++ {
		//判断回文数
		tmp = i
		for tmp != 0 {
			y = y*10 + tmp%10
			tmp = tmp / 10
		}

		//是回文数在判断是否是素数(质数)
		if y == i {
			//fmt.Println(i)
			k := uint64(2)
			t := (uint64)(math.Sqrt(float64(y)))
			for ; k < t; k++ {
				if y%k == 0 {
					break
				}
			}

			//是素数直接返回
			if k == t {
				return int(i)
			}
		}
		y = uint64(0)
	}

	return 0
}

func main() {
	smp := make(map[uint64]uint64)
	//limit := uint64(10 * 10 * 10 * 10 * 10 * 10 * 10 * 10)
	//limit := uint64(10 * 10)
	limit := uint64(20)

	sushu(limit, smp)

	//fmt.Println(primePalindrome(9989900))
	//fmt.Println(primePalindrome(13))
}
