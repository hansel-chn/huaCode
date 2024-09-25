//Given an array of distinct integers nums and a target integer target, return
//the number of possible combinations that add up to target.
//
// The test cases are generated so that the answer can fit in a 32-bit integer.
//
//
//
// Example 1:
//
//
//Input: nums = [1,2,3], target = 4
//Output: 7
//Explanation:
//The possible combination ways are:
//(1, 1, 1, 1)
//(1, 1, 2)
//(1, 2, 1)
//(1, 3)
//(2, 1, 1)
//(2, 2)
//(3, 1)
//Note that different sequences are counted as different combinations.
//
//
// Example 2:
//
//
//Input: nums = [9], target = 3
//Output: 0
//
//
//
// Constraints:
//
//
// 1 <= nums.length <= 200
// 1 <= nums[i] <= 1000
// All the elements of nums are unique.
// 1 <= target <= 1000
//
//
//
// Follow up: What if negative numbers are allowed in the given array? How does
//it change the problem? What limitation we need to add to the question to allow
//negative numbers?
//
// Related Topics 数组 动态规划 👍 1043 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
package main

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i < target+1; i++ {
		for _, num := range nums {
			if i-num >= 0 {
				dp[i] += dp[i-num]
			}
		}
	}

	return dp[target]
}

//func combinationSum4(nums []int, target int) int {
//	memo := make(map[int]int)
//	var dp func(target int) int
//	dp = func(target int) int {
//		rlt := 0
//		for _, num := range nums {
//			if target-num == 0 {
//				rlt += 1
//			} else if target-num > 0 {
//				if v, ok := memo[target-num]; ok {
//					rlt += v
//				} else {
//					memoV := dp(target - num)
//					memo[target-num] = memoV
//					rlt += memoV
//				}
//			}
//		}
//		return rlt
//	}
//	return dp(target)
//}

//leetcode submit region end(Prohibit modification and deletion)
