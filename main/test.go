package main

import (
	"fmt"
)

type number int

func (n number) nprinf() {
	fmt.Printf("===nprinf n=%p\n", &n)
	fmt.Printf("===nprinf n=%d\n", n)
}

func (n *number) nnprinf() {
	fmt.Printf("--nnprinf n=%p\n", &n)
	fmt.Println(*n)
}

//里面的defer被插入到赋值和return之间执行，t赋值给r后 执行defer 在return
func f1() (r int) {
	t := 5
	defer func() {
		t = t + 5
		fmt.Println(&t)
	}()

	fmt.Println(&r)
	fmt.Println(&t)

	return t
}

//1
func f2() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)

	return 1
}

//6
func f3() (r int) {
	defer func() {
		r = r + 5
	}()

	return 1
}

//6
func f4() (r int) {
	defer func(r *int) {
		*r = *r + 5
	}(&r)

	return 1
}

func main() {
	var n number
	fmt.Printf("main n=%p\n", &n)

	defer n.nprinf()
	defer n.nnprinf()

	//fmt.Println(f1())
	fmt.Println(f4())

	//defer func() {
	//	fmt.Printf("defer1 n=%p\n", &n)
	//	n.nprinf()
	//}()
	//
	//defer func() {
	//	fmt.Printf("defer2 n=%p\n", &n)
	//	n.nnprinf()
	//}()
	n = 3
}
