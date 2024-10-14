//Given an integer array nums, return the number of longest increasing
//subsequences.
//
// Notice that the sequence has to be strictly increasing.
//
//
// Example 1:
//
//
//Input: nums = [1,3,5,4,7]
//Output: 2
//Explanation: The two longest increasing subsequences are [1, 3, 4, 7] and [1,
//3, 5, 7].
//
//
// Example 2:
//
//
//Input: nums = [2,2,2,2,2]
//Output: 5
//Explanation: The length of the longest increasing subsequence is 1, and there
//are 5 increasing subsequences of length 1, so output 5.
//
//
//
// Constraints:
//
//
// 1 <= nums.length <= 2000
// -10⁶ <= nums[i] <= 10⁶
// The answer is guaranteed to fit inside a 32-bit integer.
//
//
// Related Topics 树状数组 线段树 数组 动态规划 👍 895 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
package main

//func findNumberOfLIS(nums []int) (ans int) {
//	maxLen := 0
//	n := len(nums)
//	dp := make([][2]int, n)
//	for i, x := range nums {
//		dp[i][0] = 1
//		dp[i][1] = 1
//		for j, y := range nums[:i] {
//			if x > y {
//				if dp[j][0]+1 > dp[i][0] {
//					dp[i][0] = dp[j][0] + 1
//					dp[i][1] = dp[j][1] // 重置计数
//				} else if dp[j][0]+1 == dp[i][0] {
//					dp[i][1] += dp[j][1]
//				}
//			}
//		}
//		if dp[i][0] > maxLen {
//			maxLen = dp[i][0]
//			ans = dp[i][1] // 重置计数
//		} else if dp[i][0] == maxLen {
//			ans += dp[i][1]
//		}
//	}
//	return
//}

func findNumberOfLIS(nums []int) int {
	dp := make([][2]int, len(nums)+1)
	maxSL := 0
	for i := 1; i < len(nums)+1; i++ {
		// optimize
		maxL := 0
		count := 1
		for j := 1; j < i; j++ {
			if nums[i-1] > nums[j-1] {
				if dp[j][0] == maxL {
					count += dp[j][1]
				} else if dp[j][0] > maxL {
					maxL = dp[j][0]
					count = dp[j][1]
				}
			}
		}
		maxSL = max(maxSL, maxL+1)
		dp[i] = [2]int{maxL + 1, count}

		//maxn := 0
		//maxIdx := make([]int, 0)
		//for j := 1; j < i; j++ {
		//	if nums[i-1] > nums[j-1] {
		//		if dp[j][0] == maxn {
		//			maxIdx = append(maxIdx, j)
		//		} else if dp[j][0] > maxn {
		//			maxn = dp[j][0]
		//			maxIdx = []int{j}
		//		}
		//	}
		//}
		//
		//if maxn == 0 {
		//	dp[i] = [2]int{1, 1}
		//} else {
		//	sum := 0
		//	for _, idx := range maxIdx {
		//		sum += dp[idx][1]
		//	}
		//	dp[i] = [2]int{maxn + 1, sum}
		//}
	}

	count := 0
	for _, v := range dp[1:] {
		if v[0] == maxSL {
			count += v[1]
		}
	}

	return count
}

//leetcode submit region end(Prohibit modification and deletion)
