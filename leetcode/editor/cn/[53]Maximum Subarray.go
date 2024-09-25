//Given an integer array nums, find the subarray with the largest sum, and
//return its sum.
//
//
// Example 1:
//
//
//Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
//Output: 6
//Explanation: The subarray [4,-1,2,1] has the largest sum 6.
//
//
// Example 2:
//
//
//Input: nums = [1]
//Output: 1
//Explanation: The subarray [1] has the largest sum 1.
//
//
// Example 3:
//
//
//Input: nums = [5,4,-1,7,8]
//Output: 23
//Explanation: The subarray [5,4,-1,7,8] has the largest sum 23.
//
//
//
// Constraints:
//
//
// 1 <= nums.length <= 10⁵
// -10⁴ <= nums[i] <= 10⁴
//
//
//
// Follow up: If you have figured out the O(n) solution, try coding another
//solution using the divide and conquer approach, which is more .
//
// Related Topics 数组 分治 动态规划 👍 6790 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
package main

import "math"

func maxSubArray(nums []int) int {
	preSums := make([]int, len(nums)+1)
	for i, num := range nums {
		preSums[i+1] = preSums[i] + num
	}

	rlt := math.MinInt

	//minPreSum := 0
	//for i := 0; i < len(nums); i++ {
	//	minPreSum = min(minPreSum, preSums[i])
	//	rlt = max(rlt, preSums[i+1]-minPreSum)
	//}

	minPreSum := 0
	for i := 1; i < len(nums)+1; i++ {
		rlt = max(rlt, preSums[i]-minPreSum)
		minPreSum = min(minPreSum, preSums[i])
	}
	return rlt
}

//func maxSubArray(nums []int) int {
//	dp := make([]int, len(nums))
//	dp[0] = nums[0]
//	for i := 1; i < len(nums); i++ {
//		dp[i] = max(dp[i-1]+nums[i], nums[i])
//	}
//
//	rlt := -int(1e5)
//	for _, v := range dp {
//		rlt = max(rlt, v)
//	}
//	return rlt
//}

//leetcode submit region end(Prohibit modification and deletion)
