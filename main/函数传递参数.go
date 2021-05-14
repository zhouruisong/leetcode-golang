package main

import (
	//"time"
	"fmt"
	"unsafe"
)

func testFunc(p1 map[string]bool, p2 []int) {
	//fmt.Println(cap(p2)) //长度1 容量1
	//sliece扩容是新分配空间,空间长度是原来长度的两倍,空间地址变了

	x := (*[3]uintptr)(unsafe.Pointer(&p2))
	fmt.Printf("in===%v\n", x[0])

	//p2[0] = "777"          //长度1 容量1
	p2 = append(p2, 2)

	//
	//fmt.Println(x)
	//
	p2 = append(p2, 3)

	fmt.Printf("in===%v\n", x[0])

	p2 = append(p2, 4) //扩容,原来容量3不够了,翻倍扩容, 长度4 容量变为6

	fmt.Printf("in===%v\n", x[0])

	fmt.Println(cap(p2))
	//
	//x := (*[3]uintptr)(unsafe.Pointer(&p2))
	//fmt.Println(x)

	//
	//fmt.Println(x)
	//p2 = append(p2, "666") //不扩容,原来容量4够了,长度4 容量4
	//fmt.Println(cap(p2))
	//
	//fmt.Println(x)
	//
	//p2 = append(p2, "999") //扩容,原来容量4不够了,翻倍扩容, 长度5 容量变为8
	//fmt.Println(cap(p2))

	//fmt.Println(x)
	//
	//p1["222"] = true
	//p1["33"] = true
	//p1["55"] = true

	fmt.Println(p1)
}

func main() {
	//runtime.GOMAXPROCS(runtime.NumCPU())
	//
	//var wg sync.WaitGroup
	//for i = 0; i < 500; i++ {
	//	wg.Add(1)
	//	go func() {
	//		fmt.Println(i)
	//		wg.Done()
	//	}()
	//	//time.Sleep(1 * time.Second)   // 设置时间延时1秒
	//}
	//wg.Wait()

	list := make([]int, 0, 3)
	list = append(list, 1)

	mp := make(map[string]bool, 1)
	mp["1111"] = true

	//fmt.Printf("mp:%p, list:%p\n", &mp, &list)

	x1 := (*[3]uintptr)(unsafe.Pointer(&list))
	fmt.Println(x1[0])

	//不扩容的数据会带回来,扩容的数据不会带回来!!!
	testFunc(mp, list)

	fmt.Println(x1[0])
	//这里slice的地址指向的还是原来的空间地址
	fmt.Println(len(list), cap(list))

	fmt.Println(list, list[1:2], list[2:3])
	//fmt.Println(list[1:2]) //未扩容的数据,地址还是原来的地址,可以发现第二个切片的值是2
	//fmt.Println(list[2:3]) //未扩容的数据,地址还是原来的地址, 可以发现第二个切片的值是3
	fmt.Println(mp)
}
