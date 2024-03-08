/**
Given an integer array nums, return true if you can partition the array into
two subsets such that the sum of the elements in both subsets is equal or false
otherwise.


 Example 1:


Input: nums = [1,5,11,5]
Output: true
Explanation: The array can be partitioned as [1, 5, 5] and [11].


 Example 2:


Input: nums = [1,2,3,5]
Output: false
Explanation: The array cannot be partitioned into equal sum subsets.



 Constraints:


 1 <= nums.length <= 200
 1 <= nums[i] <= 100


 Related Topics 数组 动态规划 👍 2021 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
package main

func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	if sum%2 != 0 {
		return false
	}

	half := sum / 2

	dp := make([][]bool, len(nums)+1)
	for i := range dp {
		dp[i] = make([]bool, half+1)
		dp[i][0] = true
	}

	for i := 1; i < len(nums)+1; i++ {
		for j := 1; j < half+1; j++ {
			if j-nums[i-1] < 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j-nums[i-1]] || dp[i-1][j]
			}
		}
	}
	return dp[len(nums)][half]
}

// 如果i和数组索引对齐，对于边界条件处理太不方便了，需要手动判断数组第一个值
//	if nums[0] <= half {
//		dp[0][nums[0]] = true
//	}
// 所以DP Table和数组索引错开，dp[i][j]代表数组索引前i-1个数包括i-1，值的和能否构成j
// ===========================================================
//func canPartition(nums []int) bool {
//	sum := 0
//	for _, num := range nums {
//		sum += num
//	}
//
//	if sum%2 != 0 {
//		return false
//	}
//
//	half := sum / 2
//
//	dp := make([][]bool, len(nums))
//	for i := range dp {
//		dp[i] = make([]bool, half+1)
//		dp[i][0] = true
//	}
//
//	if nums[0] <= half {
//		dp[0][nums[0]] = true
//	}
//
//	for i := 1; i < len(nums); i++ {
//		for j := 1; j < half+1; j++ {
//			if j-nums[i] < 0 {
//				dp[i][j] = dp[i-1][j]
//			} else {
//				dp[i][j] = dp[i-1][j-nums[i]] || dp[i-1][j]
//			}
//		}
//	}
//	return dp[len(nums)-1][half]
//}

//leetcode submit region end(Prohibit modification and deletion)
