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

func main() {
	x := int64(-1234567893889895620)
	r := reserve(x)
	fmt.Printf("ret=%+v\n", r)
}
