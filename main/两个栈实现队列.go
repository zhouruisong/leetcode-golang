package main

var stack1 []int
var stack2 []int

func Push(node int) {
	stack1 = append(stack1, node)
}

func Pop() int {
	if len(stack2) == 0 {
		for i := len(stack1) - 1; i >= 0; i-- {
			stack2 = append(stack2, stack1[i])
		}
		stack1 = make([]int, 0)
	}
	if len(stack2) > 0 {
		ret := stack2[len(stack2)-1]
		stack2 = stack2[:len(stack2)-1]
		return ret
	}
	return -1
}

func main() {

}
