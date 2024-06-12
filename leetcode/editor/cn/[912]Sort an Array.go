/**
Given an array of integers nums, sort the array in ascending order and return
it.

 You must solve the problem without using any built-in functions in O(nlog(n))
time complexity and with the smallest space complexity possible.


 Example 1:


Input: nums = [5,2,3,1]
Output: [1,2,3,5]
Explanation: After sorting the array, the positions of some numbers are not
changed (for example, 2 and 3), while the positions of other numbers are changed (
for example, 1 and 5).


 Example 2:


Input: nums = [5,1,1,2,0,0]
Output: [0,0,1,1,2,5]
Explanation: Note that the values of nums are not necessairly unique.



 Constraints:


 1 <= nums.length <= 5 * 10⁴
 -5 * 10⁴ <= nums[i] <= 5 * 10⁴


 Related Topics 数组 分治 桶排序 计数排序 基数排序 排序 堆（优先队列） 归并排序 👍 1004 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
package main

// // time limit exceeded
// // 因为相同元素过多，退化成n^2
//func sortArray(nums []int) []int {
//	var quickSort func(lo, hi int)
//	quickSort = func(lo, hi int) {
//		if lo >= hi {
//			return
//		}
//
//		pivot := lo
//		left := lo + 1
//		right := hi
//
//		for left <= right {
//			for left <= hi && nums[left] <= nums[pivot] {
//				left++
//			}
//			for lo < right && nums[right] > nums[pivot] {
//				right--
//			}
//
//			if left >= right {
//				break
//			}
//			sortArraySwap(nums, left, right)
//			left++
//			right--
//		}
//		sortArraySwap(nums, pivot, right)
//
//		quickSort(lo, right-1)
//		quickSort(right+1, hi)
//	}
//
//	quickSort(0, len(nums)-1)
//	return nums
//}

func sortArray(nums []int) []int {
	var quickSort func(lo, hi int)
	quickSort = func(lo, hi int) {
		if lo >= hi {
			return
		}

		pivot := lo
		left := lo + 1
		right := hi

		for left <= right {
			for left <= hi && nums[left] < nums[pivot] {
				left++
			}
			for lo < right && nums[right] > nums[pivot] {
				right--
			}

			if left >= right {
				break
			}
			sortArraySwap(nums, left, right)
			left++
			right--
		}
		sortArraySwap(nums, pivot, right)

		quickSort(lo, right-1)
		quickSort(right+1, hi)
	}

	quickSort(0, len(nums)-1)
	return nums
}

func sortArraySwap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

//leetcode submit region end(Prohibit modification and deletion)
