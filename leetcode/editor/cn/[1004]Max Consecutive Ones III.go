/**
Given a binary array nums and an integer k, return the maximum number of
consecutive 1's in the array if you can flip at most k 0's.


 Example 1:


Input: nums = [1,1,1,0,0,0,1,1,1,1,0], k = 2
Output: 6
Explanation: [1,1,1,0,0,1,1,1,1,1,1]
Bolded numbers were flipped from 0 to 1. The longest subarray is underlined.

 Example 2:


Input: nums = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], k = 3
Output: 10
Explanation: [0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1]
Bolded numbers were flipped from 0 to 1. The longest subarray is underlined.



 Constraints:


 1 <= nums.length <= 10⁵
 nums[i] is either 0 or 1.
 0 <= k <= nums.length


 Related Topics 数组 二分查找 前缀和 滑动窗口 👍 696 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
package main

func longestOnes(nums []int, k int) int {
	var left, right, flip, maxLen int
	for right < len(nums) {
		addV := nums[right]
		right++
		if addV == 0 {
			flip++
		}

		for flip > k {
			if nums[left] == 0 {
				flip--
			}
			left++
		}
		maxLen = max(maxLen, right-left)
	}
	return maxLen
}

//leetcode submit region end(Prohibit modification and deletion)
