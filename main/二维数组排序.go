package main

import (
	"fmt"
	"sort"
)

type myvalue2 [][]int

//type myvalue2 [][]string

func (p myvalue2) Len() int {
	return len(p)
}

// 想了想能用循环，实现所有字段比对一次，更精确点，免得二维数组的元素前面n个元素都是一样的
func (p myvalue2) Less(i, j int) bool {
	for k := 0; k < len(p[i]); k++ {
		if p[i][k] > p[j][k] {
			return false
		} else if p[i][k] == p[j][k] {
			continue
		} else {
			return true
		}
	}
	return true
}

func (p myvalue2) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	var m myvalue2
	//m = [][]string{
	//	{"a", "a", "g", "c", "z", "d"},
	//	{"a", "c", "g", "v", "z", "d"},
	//	{"r", "a", "g", "c", "z", "a"},
	//	{"a", "a", "g", "b", "a", "v"},
	//	{"a", "a", "g", "b", "a", "a"},
	//	{"a", "a", "g", "b", "z", "a"},
	//}

	//"123"
	//"132"
	//"213"
	//"231"
	//"312"
	//"321"
	m = [][]int{
		{2, 1, 3},
		{1, 2, 3},
		{3, 2, 1},
	}
	sort.Sort(m)
	fmt.Println(m)
}
