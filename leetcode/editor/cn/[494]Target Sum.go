/**
You are given an integer array nums and an integer target.

 You want to build an expression out of nums by adding one of the symbols '+'
and '-' before each integer in nums and then concatenate all the integers.


 For example, if nums = [2, 1], you can add a '+' before 2 and a '-' before 1
and concatenate them to build the expression "+2-1".


 Return the number of different expressions that you can build, which evaluates
to target.


 Example 1:


Input: nums = [1,1,1,1,1], target = 3
Output: 5
Explanation: There are 5 ways to assign symbols to make the sum of nums be
target 3.
-1 + 1 + 1 + 1 + 1 = 3
+1 - 1 + 1 + 1 + 1 = 3
+1 + 1 - 1 + 1 + 1 = 3
+1 + 1 + 1 - 1 + 1 = 3
+1 + 1 + 1 + 1 - 1 = 3


 Example 2:


Input: nums = [1], target = 1
Output: 1



 Constraints:


 1 <= nums.length <= 20
 0 <= nums[i] <= 1000
 0 <= sum(nums[i]) <= 1000
 -1000 <= target <= 1000


 Related Topics 数组 动态规划 回溯 👍 1882 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
package main

func findTargetSumWays(nums []int, target int) int {
	maxSum := 0
	for _, v := range nums {
		maxSum += v
	}

	leftBoundary := 0
	rightBoundary := maxSum * 2
	dp := make([][]int, len(nums))
	for i := range dp {
		dp[i] = make([]int, rightBoundary+1)
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < rightBoundary+1; j++ {
			if i == 0 {
				if mapIdx(maxSum, j) == nums[0] {
					dp[i][j] += 1
				}

				if mapIdx(maxSum, j) == -nums[0] {
					dp[i][j] += 1
				}
				continue
			}

			if j-nums[i] >= leftBoundary && dp[i-1][j-nums[i]] > 0 {
				dp[i][j] += dp[i-1][j-nums[i]]
			}

			if j+nums[i] <= rightBoundary && dp[i-1][j+nums[i]] > 0 {
				dp[i][j] += dp[i-1][j+nums[i]]
			}
		}
	}

	if mapValue(maxSum, target) < leftBoundary || mapValue(maxSum, target) > rightBoundary {
		return 0
	} else {
		return dp[len(nums)-1][mapValue(maxSum, target)]
	}
}

func mapValue(maxSum int, value int) (idx int) {
	return value - (-maxSum)
}

func mapIdx(maxSum int, idx int) (value int) {
	return idx - maxSum
}

//leetcode submit region end(Prohibit modification and deletion)
