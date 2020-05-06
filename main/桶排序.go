package main

import "fmt"

//分治法，也成基数法
//99, 799 123 246, 1650, 2370
//800万考生 分数排序 0-750桶
func radixSort(arr []int) []int {
	max := 11 //最大数
	for bit := 1; max/bit > 0; bit *= 10 {
		//按照数量级分段
		arr = BitSort(arr, bit) //每次处理一个级别的排序
		fmt.Println(arr)
	}

	return arr
}

func BitSort(arr []int, bit int) []int {
	n := len(arr)
	bitcounts := make([]int, 10) //统计长度0-9
	for i := 0; i < n; i++ {
		num := (arr[i] / bit) % 10 //分层处理nit = 100三位数不参与排序，bit=1000 四位数不参与排序
		bitcounts[num]++
	}

	fmt.Println(bitcounts)

	//0,2,1,3,4,5
	//1 0 3 0 0 1
	//标记每个数在哪个位置
	for i := 1; i < 10; i++ {
		bitcounts[i] += bitcounts[i-1] // 叠加计算位置
	}
	fmt.Println(bitcounts)

	tmp := make([]int, 10) //开辟临时数组
	for i := n - 1; i >= 0; i-- {
		num := (arr[i] / bit) % 10
		tmp[bitcounts[num]-1] = arr[i] //计算排序的尾椎
		bitcounts[num]--
	}

	for i := 0; i < n; i++ {
		arr[i] = tmp[i] //保存数组
	}

	return arr

}
func main() {
	arr := []int{1, 9, 2, 8, 3, 4, 6, 5, 10, 11}
	fmt.Println(radixSort(arr))
}
