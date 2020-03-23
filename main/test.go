package main

import (
	"fmt"
)

type number int

func (n number) nprinf() {
	fmt.Printf("nprinf n=%p\n", &n)
	fmt.Printf("nprinf n=%d\n", n)
}

func (n *number) nnprinf() {
	fmt.Printf("nnprinf n=%p\n", n)
	fmt.Println(*n)
}

func main() {
	var n number
	fmt.Printf("main n=%p\n", &n)

	defer func() {
		fmt.Printf("defer1 n=%p\n", &n)
		n.nprinf()
	}()

	defer func() {
		fmt.Printf("defer2 n=%p\n", &n)
		n.nnprinf()
	}()
	n = 3
}
