package main

/*
题目描述
实现一个特殊功能的栈，在实现栈的基本功能的基础上，再实现返回栈中最小元素的操作。
示例1
输入
复制
[[1,3],[1,2],[1,1],[3],[2],[3]]
返回值
复制
[1,2]
备注:
有三种操作种类，op1表示push，op2表示pop，op3表示getMin。你需要返回和op3出现次数一样多的数组，表示每次getMin的答案

1<=操作总数<=1000000
-1000000<=每个操作数<=1000000
数据保证没有不合法的操作
*/

/**
 * return a array which include all ans for op3
 * @param op int整型二维数组 operator
 * @return int整型一维数组
 */

type MyStack struct {
	Num []int
	Min []int
}

func (m *MyStack) Push(x int) {
	m.Num = append(m.Num, x)
	if len(m.Min) > 0 {
		min := m.Min[len(m.Min)-1]
		if x < min {
			m.Min = append(m.Min, x)
		} else {
			m.Min = append(m.Min, min)
		}
	} else {
		m.Min = append(m.Min, x)
	}
}

func (m *MyStack) Pop() int {
	x := m.Num[len(m.Num)-1]
	m.Num = m.Num[:len(m.Num)-1]
	m.Min = m.Min[:len(m.Min)-1]
	return x
}

func (m *MyStack) GetMin() int {
	return m.Min[len(m.Min)-1]
}

func getMinStack(op [][]int) []int {
	// write code here
	m := MyStack{
		Min: make([]int, 0),
		Num: make([]int, 0),
	}

	res := []int{}
	for _, v := range op {
		if v[0] == 1 {
			m.Push(v[1])
		} else if v[0] == 2 {
			m.Pop()
		} else if v[0] == 3 {
			res = append(res, m.GetMin())
		}
	}

	return res
}
