package main

import "fmt"

/*
* 假设按照升序排序的数组在预先未知的某个点上进行了旋转。
 *
 * ( 例如，数组 [0,0,1,2,2,5,6] 可能变为 [2,5,6,0,0,1,2] )。
 *
 * 编写一个函数来判断给定的目标值是否存在于数组中。若存在返回 true，否则返回 false。
 *
 * 示例 1:
 *
 * 输入: nums = [2,5,6,0,0,1,2], target = 0
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入: nums = [2,5,6,0,0,1,2], target = 3
 * 输出: false
 *
 * 进阶:
 *
 *
 * 这是 搜索旋转排序数组 的延伸题目，本题中的 nums  可能包含重复元素。
 * 这会影响到程序的时间复杂度吗？会有怎样的影响，为什么？
 *
 *
*/
/*
直接二分，如果中间值=target值，直接返回true，然后比较left和mid值，如果二者相等，不停的右移left，此时mid也在不停的右移，
直到left值和mid值不相等，此时，mid有可能在转折点后面，也可能在转折点前面，如果left值大于mid值说明，mid在转折点后面
此时mid的后面是有序的子序列，然后判断，mid，right值和target的大小关系，如果在此间，这在这中间二分，不在的话，则right=mid-1，
重新二分。回到前面，mid如果在转折点前面，这前面序列是有序的，判断mid，left值和target大小，如果target在此间，则在这中间二分，
否则left=mid+1，重新二分
*/

func search(nums []int, target int) bool {
	if len(nums) == 0 || (len(nums) == 1 && nums[0] != target) {
		return false
	}
	if len(nums) == 1 && nums[0] == target {
		return true
	}
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return true
		}
		if nums[left] == nums[mid] {
			left++
			continue
		}
		//mid后面是顺序的，mid在转折点后面
		if nums[left] > nums[mid] {
			if nums[mid] < target && nums[right] >= target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else { //mid在转折点前面
			if nums[mid] > target && nums[left] <= target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return false
}

func main() {
	fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 0))
}
