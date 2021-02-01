package main

import (
	"fmt"
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var x []int
	if len(nums1) != 0 {
		x = append(x, nums1...)
	}
	if len(nums2) != 0 {
		x = append(x, nums2...)
	}

	fmt.Println(x)

	n := len(x)
	if n == 0 {
		return 0.0
	}

	var arr []int
	arr = append(arr, x...)
	fmt.Println(arr)

	mid := (n - 1) / 2
	left := 0
	right := n - 1
	k := -1
	for mid != k {
		k = partion(x, left, right)
		if k < mid {
			left = k + 1
		} else if k > mid {
			right = k - 1
		}
	}

	fmt.Println(x)

	//原来的数组已经改变，返回的索引是原来数组的位置
	// fmt.Println(nums)
	// fmt.Println(x)

	//fmt.Println(mid, x[mid])
	ans := float64(1.0)
	//偶数
	if n%2 == 0 {
		ans = (float64)(float64(arr[mid])+float64(arr[mid+1])) / 2.0
	} else {
		ans = float64(arr[mid])
	}

	fmt.Println(ans)
	return ans
}

//快排序思想查找
func findMidium(nums []int) int {
	var x []int
	x = append(x, nums...)

	fmt.Println(nums)

	n := len(nums)
	mid := (n - 1) / 2
	left := 0
	right := n - 1
	k := -1
	for mid != k {
		k = partion(nums, left, right)
		if k < mid {
			left = k + 1
		} else if k > mid {
			right = k - 1
		}
	}

	//原来的数组已经改变，返回的索引是原来数组的位置
	fmt.Println(x)
	fmt.Println(nums)
	//fmt.Println(x)

	fmt.Println(mid, x[mid])
	ans := float32(1.0)
	//偶数
	if n%2 == 0 {
		ans = (float32)(float32(x[mid])+float32(x[mid+1])) / 2.0
	} else {
		ans = float32(x[mid])
	}

	fmt.Println(ans)

	return nums[mid]

}

func swap(x *int, y *int) {
	other := *x
	*x = *y
	*y = other
}

//找到一个数轴，这个位置左边的都小于该数，右边的都大于该数
func partion(nums []int, low, high int) int {
	if low > high {
		return 0
	}

	i := low
	j := high
	val := nums[low]
	//fmt.Println(nums)
	//fmt.Println("\n=====")

	for i < j {
		for i < j && nums[j] > val {
			j--
		}
		swap(&nums[i], &nums[j])
		//fmt.Println(nums)

		for i < j && nums[i] < val {
			i++
		}
		swap(&nums[i], &nums[j])
		//fmt.Println(nums)
	}

	nums[i] = val
	//fmt.Println("\n=====")

	//fmt.Println(nums, i)
	partion(nums, 0, i-1)
	partion(nums, i+1, high)
	return i
}

func findMedianSortedArrays3(nums1 []int, nums2 []int) float64 {
	//把数组长度小的作为nums1
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	m := len(nums1)
	n := len(nums2)

	// 分割线左边的所有元素需要满足的个数 m + (n - m + 1) / 2;
	totalLeft := (m + n + 1) / 2

	// 在 nums1 的区间 [0, m] 里查找恰当的分割线，
	// 使得 nums1[i - 1] <= nums2[j] && nums2[j - 1] <= nums1[i]
	left := 0
	right := m
	for left < right {
		i := left + (right-left+1)/2
		j := totalLeft - i
		if nums1[i-1] > nums2[j] {
			// 下一轮搜索的区间 [left, i - 1]
			right = i - 1
		} else {
			// 下一轮搜索的区间 [i, right]
			left = i
		}
	}

	i := left
	j := totalLeft - i

	nums1LeftMax := math.MinInt32
	if i > 0 {
		nums1LeftMax = nums1[i-1]
	}

	nums1RightMin := math.MaxInt32
	if i != m {
		nums1RightMin = nums1[i]
	}

	nums2LeftMax := math.MinInt32
	if j > 0 {
		nums2LeftMax = nums2[j-1]
	}

	nums2RightMin := math.MaxInt32
	if j != n {
		nums2RightMin = nums2[j]
	}

	if ((m + n) % 2) == 1 {
		return math.Max(float64(nums1LeftMax), float64(nums2LeftMax))
	} else {
		return (math.Max(float64(nums1LeftMax), float64(nums2LeftMax)) + math.Min(float64(nums1RightMin), float64(nums2RightMin))) / 2
	}
}

func main() {
	//x := []int{1,13,2,55,5,8,7,9}
	//low := 0
	//high := len(x) - 1

	////快速排序
	//if low < high {
	//partion(x, 0, high)
	//partion(x, 0, mid-1)
	//partion(x, mid+1, high)
	//}

	//fmt.Println(x)
	//findMidium(x)
	fmt.Println(findMedianSortedArrays3([]int{1, 3}, []int{2}))
}
