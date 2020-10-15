package main

import (
	"container/list"
	"fmt"
)

/*
 * 这里有一个非负整数数组 arr，你最开始位于该数组的起始下标 start 处。当你位于下标 i 处时，你可以跳到 i + arr[i] 或者 i -
 * arr[i]。
 *
 * 请你判断自己是否能够跳到对应元素值为 0 的 任意 下标处。
 *
 * 注意，不管是什么情况下，你都无法跳到数组之外。
 *
 * 示例 1：
 *
 * 输入：arr = [4,2,3,0,3,1,2], start = 5
 * 输出：true
 * 解释：
 * 到达值为 0 的下标 3 有以下可能方案：
 * 下标 5 -> 下标 4 -> 下标 1 -> 下标 3
 * 下标 5 -> 下标 6 -> 下标 4 -> 下标 1 -> 下标 3
 *
 *
 * 示例 2：
 *
 * 输入：arr = [4,2,3,0,3,1,2], start = 0
 * 输出：true
 * 解释：
 * 到达值为 0 的下标 3 有以下可能方案：
 * 下标 0 -> 下标 4 -> 下标 1 -> 下标 3
 *
 *
 * 示例 3：
 *
 * 输入：arr = [3,0,2,1,2], start = 2
 * 输出：false
 * 解释：无法到达值为 0 的下标 1 处。
 *
 * 提示：
 *
 * 1 <= arr.length <= 5 * 10^4
 * 0 <= arr[i] < arr.length
 * 0 <= start < arr.length
 */

//广度优先搜索 o(n)
/*
我们可以使用广度优先搜索的方法得到从 start 开始能够到达的所有位置，如果其中某个位置对应的元素值为 0，那么就返回 True。
具体地，我们初始时将 start 加入队列。在每一次的搜索过程中，我们取出队首的节点 u，它可以到达的位置为 u + arr[u] 和 u - arr[u]。
如果某个位置落在数组的下标范围 [0, len(arr)) 内，并且没有被搜索过，则将该位置加入队尾。只要我们搜索到一个对应元素值为 0 的位置，
我们就返回 True。在搜索结束后，如果仍然没有找到符合要求的位置，我们就返回 False。
*/
func canReach(arr []int, start int) bool {
	if arr[start] == 0 {
		return true
	}

	n := len(arr)
	q := list.New()
	q.PushBack(start)
	mp := make(map[int]bool)
	mp[start] = true

	fmt.Println("===", start)
	for q.Len() > 0 {
		u := q.Remove(q.Front()).(int)

		if u+arr[u] < n && !mp[u+arr[u]] {
			if arr[u+arr[u]] == 0 {
				fmt.Println("===", u+arr[u])
				return true
			}
			q.PushBack(u + arr[u])
			mp[u+arr[u]] = true
			fmt.Println("===", u+arr[u])
		}

		if u-arr[u] >= 0 && !mp[u-arr[u]] {
			if arr[u-arr[u]] == 0 {
				fmt.Println("===", u-arr[u])
				return true
			}
			q.PushBack(u - arr[u])
			mp[u-arr[u]] = true
			fmt.Println("---", u-arr[u])
		}
	}

	return false
}

func main() {
	fmt.Println(canReach([]int{4, 2, 3, 0, 3, 1, 2}, 5))
	//fmt.Println(canReach([]int{4, 2, 3, 0, 3, 1, 2}, 0))
}
