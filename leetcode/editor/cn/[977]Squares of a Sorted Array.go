//Given an integer array nums sorted in non-decreasing order, return an array
//of the squares of each number sorted in non-decreasing order.
//
//
// Example 1:
//
//
//Input: nums = [-4,-1,0,3,10]
//Output: [0,1,9,16,100]
//Explanation: After squaring, the array becomes [16,1,0,9,100].
//After sorting, it becomes [0,1,9,16,100].
//
//
// Example 2:
//
//
//Input: nums = [-7,-3,2,3,11]
//Output: [4,9,9,49,121]
//
//
//
// Constraints:
//
//
// 1 <= nums.length <= 10⁴
// -10⁴ <= nums[i] <= 10⁴
// nums is sorted in non-decreasing order.
//
//
//
//Follow up: Squaring each element and sorting the new array is very trivial,
//could you find an
//O(n) solution using a different approach?
//
// Related Topics 数组 双指针 排序 👍 1085 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
package main

import (
	"slices"
)

func sortedSquares(nums []int) []int {
	var pivot int = len(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] >= 0 {
			pivot = i
			break
		}
	}

	var l, r int = pivot - 1, pivot

	rlt := make([]int, len(nums))
	var idx int

	for i := 0; i < len(nums); i++ {
		nums[i] = nums[i] * nums[i]
	}

	for l >= 0 && r < len(nums) {
		if nums[l] >= nums[r] {
			rlt[idx] = nums[r]
			r++
		} else {
			rlt[idx] = nums[l]
			l--
		}
		idx++
	}

	if l >= 0 {
		slices.Reverse(nums[:l+1])
		copy(rlt[idx:], nums[:l+1])
	} else {
		copy(rlt[idx:], nums[r:])
	}

	return rlt
}

//leetcode submit region end(Prohibit modification and deletion)
