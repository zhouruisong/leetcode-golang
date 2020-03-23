package main

import "fmt"

//时间复杂度o(N)
func MyPower(x float64, n int) float64 {
	var N float64
	N = (float64)(n)
	if n < 0 {
		x = 1 / x
		N = -N
	}

	var i float64
	var ret float64
	i = 0
	ret = 1.0
	for ; i < N; i++ {
		ret = ret * x
	}

	return ret
}

//时间复杂度o(logN)
func MyPowerOther(x float64, n int) float64 {
	var N float64
	N = (float64)(n)
	if n < 0 {
		x = 1 / x
		N = -N
	}
	var ret float64
	var i float64
	i = N
	ret = 1.0

	var current float64
	current = x

	for ; i != 0; i /= 2 {
		if int(i)%2 == 1 {
			ret = ret * current
		}

		current = current * current
	}

	return ret
}

func main() {
	x := 2.0
	y := -4
	//fmt.Printf("ret=%+v\n", MyPower(x, y))
	fmt.Printf("ret=%+v\n", MyPowerOther(x, y))
}
