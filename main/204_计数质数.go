package main

import (
	"fmt"
	"time"
)

/*
	* 统计所有小于非负整数 n 的质数的数量。
	 *
	 * 示例:
	 *
	 * 输入: 10
	 * 输出: 4
	 * 解释: 小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。
	 *
*/

//o(n*sqrt(n))
func countPrimes(n int) int {
	start := time.Now()
	defer func() {
		end := time.Since(start)
		fmt.Printf("countPrimes time cost = %v\n", end)
	}()

	count := 0
	for i := 2; i < n; i++ {
		if isPrime(i) {
			count++
		}
	}
	return count
}

// 判断整数 n 是否是素数
func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

/*
首先从 2 开始，我们知道 2 是一个素数，那么 2 × 2 = 4, 3 × 2 = 6, 4 × 2 = 8... 都不可能是素数了。
然后我们发现 3 也是素数，那么 3 × 2 = 6, 3 × 3 = 9, 3 × 4 = 12... 也都不可能是素数了。
看到这里，你是否有点明白这个排除法的逻辑了呢？先看我们的第一版代码：
*/

func countPrimes2(n int) int {
	start := time.Now()
	defer func() {
		end := time.Since(start)
		fmt.Printf("countPrimes2 time cost = %v\n", end)
	}()

	s := make([]byte, n)
	//fmt.Println(s)
	for i := 2; i < n; i++ {
		if s[i] == 0 {
			// i 的倍数不可能是素数了
			for j := 2 * i; j < n; j += i {
				s[j] = 1
			}
		}
	}

	//fmt.Println(s)
	count := 0
	for k := 2; k < n; k++ {
		if s[k] != 1 {
			//fmt.Println(k)
			count++
		}
	}
	return count
}

//优化countPrimes2后的，时间复杂度o(nloglogn)
func countPrimes3(n int) int {
	start := time.Now()
	defer func() {
		end := time.Since(start)
		fmt.Printf("countPrimes3 time cost = %v\n", end)
	}()

	s := make([]uint8, n)
	for i := 2; i*i <= n; i++ {
		if 0 == s[i] {
			// i 的倍数不可能是素数了
			for j := i * i; j < n; j += i {
				s[j] = 1
			}
		}
	}

	count := 0
	for k := 2; k < n; k++ {
		if 0 == s[k] {
			//fmt.Println(k)
			count++
		}
	}
	return count
}

//优化countPrimes3后的，时间复杂度o(nloglogn)
func countPrimes4(n int) int {
	start := time.Now()
	defer func() {
		end := time.Since(start)
		fmt.Printf("countPrimes4 time cost = %v\n", end)
	}()

	s := make([]byte, n)
	for i := 2; i*i <= n; i++ {
		if 0 == s[i] {
			// i 的倍数不可能是素数了
			for j := i * i; j < n; j += i {
				s[j] = 1
			}
		}
	}

	count := 0
	for k := 2; k < n; k++ {
		if 0 == s[k] {
			//fmt.Println(k)
			count++
		}
	}
	return count
}

func main() {
	x := 100000000
	//fmt.Println(countPrimes3(x))

	fmt.Println(countPrimes4(x))
	//
	//fmt.Println(countPrimes2(x))
	//
	//fmt.Println(countPrimes(x))

}
