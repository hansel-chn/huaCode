/**
Given an integer array nums, return the length of the longest strictly
increasing subsequence.


 Example 1:


Input: nums = [10,9,2,5,3,7,101,18]
Output: 4
Explanation: The longest increasing subsequence is [2,3,7,101], therefore the
length is 4.


 Example 2:


Input: nums = [0,1,0,3,2,3]
Output: 4


 Example 3:


Input: nums = [7,7,7,7,7,7,7]
Output: 1



 Constraints:


 1 <= nums.length <= 2500
 -10⁴ <= nums[i] <= 10⁴



 Follow up: Can you come up with an algorithm that runs in O(n log(n)) time
complexity?

 Related Topics 数组 二分查找 动态规划 👍 3579 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
package main

import (
	"math"
	"sort"
)

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))

	assist := make([]int, 1)
	assist[0] = math.MinInt
	for i := 0; i < len(nums); i++ {
		idx := sort.SearchInts(assist, nums[i])
		if idx < len(assist) {
			assist[idx] = min(assist[idx], nums[i])
		} else {
			assist = append(assist, nums[i])
		}

		dp[i] = len(assist)-1
	}
	return dp[len(nums)-1]
}

//func lengthOfLIS(nums []int) int {
//	if len(nums) == 0 {
//		return 0
//	}
//
//	dp := make([]int, len(nums))
//	for i := range dp {
//		dp[i] = 1
//	}
//
//	assist := make([]int, 1)
//	assist[0] = math.MinInt
//	for i := 0; i < len(nums); i++ {
//		idx := sort.SearchInts(assist, nums[i])
//		dp[i] = idx
//		if idx < len(assist) {
//			assist[idx] = min(assist[idx], nums[i])
//		} else {
//			assist = append(assist, nums[i])
//		}
//	}
//
//		res := 0
//		for _, v := range dp {
//			res = max(res, v)
//		}
//
//		return res
//}

//func lengthOfLIS(nums []int) int {
//	if len(nums) == 0 {
//		return 0
//	}
//
//	dp := make([]int, len(nums))
//	for i := range dp {
//		dp[i] = 1
//	}
//
//	for i := 0; i < len(dp); i++ {
//		for j := 0; j < i; j++ {
//			if nums[j] < nums[i] {
//				dp[i] = max(dp[i], dp[j]+1)
//			}
//		}
//	}
//
//	res := 0
//	for _, v := range dp {
//		res = max(res, v)
//	}
//
//	return res
//}

//// iterate
//func lengthOfLIS(nums []int) int {
//	memo := make([]int, len(nums))
//
//	var dp func(nums []int, idx int) int
//	dp = func(nums []int, idx int) int {
//		if idx < 0 {
//			return 0
//		}
//
//		if memo[idx] != 0 {
//			return memo[idx]
//		}
//
//		rlt := 1
//		for i := idx - 1; i >= 0; i-- {
//			if nums[i] < nums[idx] {
//				rlt = max(rlt, dp(nums, i))
//			}
//		}
//		memo[idx] = rlt
//		return rlt
//	}
//
//	//This solution is incorrect because the dp function obtains the longest subsequence containing the last index
//	//instead of the longest subsequence in the current dp table
//	return dp(nums, len(nums)-1)
//}

func max(i, j int) int {
	if i >= j {
		return i
	} else {
		return j
	}
}

//leetcode submit region end(Prohibit modification and deletion)
