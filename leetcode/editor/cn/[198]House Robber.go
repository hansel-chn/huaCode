//You are a professional robber planning to rob houses along a street. Each
//house has a certain amount of money stashed, the only constraint stopping you from
//robbing each of them is that adjacent houses have security systems connected and
//it will automatically contact the police if two adjacent houses were broken
//into on the same night.
//
// Given an integer array nums representing the amount of money of each house,
//return the maximum amount of money you can rob tonight without alerting the
//police.
//
//
// Example 1:
//
//
//Input: nums = [1,2,3,1]
//Output: 4
//Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
//Total amount you can rob = 1 + 3 = 4.
//
//
// Example 2:
//
//
//Input: nums = [2,7,9,3,1]
//Output: 12
//Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5
//(money = 1).
//Total amount you can rob = 2 + 9 + 1 = 12.
//
//
//
// Constraints:
//
//
// 1 <= nums.length <= 100
// 0 <= nums[i] <= 400
//
//
// Related Topics 数组 动态规划 👍 3079 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
package main

func rob(nums []int) int {
	dp := make([][2]int, len(nums))

	dp[0][0] = 0
	dp[0][1] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i][1] = dp[i-1][0] + nums[i]
		dp[i][0] = max(dp[i-1][0], dp[i-1][1])
	}

	return max(dp[len(nums)-1][0], dp[len(nums)-1][1])
}

//leetcode submit region end(Prohibit modification and deletion)
