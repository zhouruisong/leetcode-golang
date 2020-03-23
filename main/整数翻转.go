package main

import "fmt"

func reserve(x int64) int64 {
	var y int64
	y = 0
	for x != 0 {
		y = y*10 + x%10

		if !(-(1<<63) <= y && y <= (1<<63)-1) {
			return 0
		}
		x /= 10
	}

	return y
}

func reverse2(x int) int {
	var result int
	if x < 0 {
		x = -x
	}

	step := 10
	num := 0
	var tmp []int
	for x > 0 {
		yu := x % step
		num++
		x = x / step
		tmp = append(tmp, yu)
	}

	fmt.Println(tmp)

	//实现10的几次方累加
	t := 1
	n := len(tmp)
	for i := 0; i < n; i++ {
		k := n - i - 1
		for k > 0 {
			t = t * step
			k--
		}

		fmt.Println(tmp[i], i, t)
		result = result + t*tmp[i]
		t = 1
	}

	fmt.Println(result)
	fmt.Println(1 << 31)
	fmt.Println(1>>31 - 1)

	if -(1<<31) <= result && result <= (1<<31-1) {
		return 0
	}

	return 0
}

func main() {
	x := int64(-1234567893889895620)
	r := reserve(x)
	fmt.Printf("ret=%+v\n", r)
}
