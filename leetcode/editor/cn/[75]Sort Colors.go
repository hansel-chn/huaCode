//Given an array nums with n objects colored red, white, or blue, sort them in-
//place so that objects of the same color are adjacent, with the colors in the
//order red, white, and blue.
//
// We will use the integers 0, 1, and 2 to represent the color red, white, and
//blue, respectively.
//
// You must solve this problem without using the library's sort function.
//
//
// Example 1:
//
//
//Input: nums = [2,0,2,1,1,0]
//Output: [0,0,1,1,2,2]
//
//
// Example 2:
//
//
//Input: nums = [2,0,1]
//Output: [0,1,2]
//
//
//
// Constraints:
//
//
// n == nums.length
// 1 <= n <= 300
// nums[i] is either 0, 1, or 2.
//
//
//
// Follow up: Could you come up with a one-pass algorithm using only constant
//extra space?
//
// Related Topics 数组 双指针 排序 👍 1886 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
package main

func sortColors(nums []int) {
	var p0, p, p2 = 0, 0, len(nums) - 1
	for p <= p2 {
		switch nums[p] {
		case 0:
			nums[p], nums[p0] = nums[p0], nums[p]
			p0++
			p++
		case 1:
			p++
		case 2:
			nums[p], nums[p2] = nums[p2], nums[p]
			p2--
		}
	}
}

//func sortColors(nums []int) {
//	count := make([]int, 3)
//	for _, num := range nums {
//		count[num]++
//	}
//
//	idx:=0
//	for i := 0; i < 3; i++ {
//		for j := 0; j < count[i]; j++ {
//			nums[idx] = i
//			idx++
//		}
//	}
//}

//leetcode submit region end(Prohibit modification and deletion)
