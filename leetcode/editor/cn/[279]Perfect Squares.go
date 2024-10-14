//Given an integer n, return the least number of perfect square numbers that
//sum to n.
//
// A perfect square is an integer that is the square of an integer; in other
//words, it is the product of some integer with itself. For example, 1, 4, 9, and 16
//are perfect squares while 3 and 11 are not.
//
//
// Example 1:
//
//
//Input: n = 12
//Output: 3
//Explanation: 12 = 4 + 4 + 4.
//
//
// Example 2:
//
//
//Input: n = 13
//Output: 2
//Explanation: 13 = 4 + 9.
//
//
//
// Constraints:
//
//
// 1 <= n <= 10⁴
//
//
// Related Topics 广度优先搜索 数学 动态规划 👍 2047 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
package main

import "math"

func numSquares(n int) int {
	dp := make([]int, n+1)

	for i := 1; i < len(dp); i++ {
		minn := 100000
		for j := 1; j*j <= i; j++ {
			minn = min(minn, dp[i-j*j])
		}
		dp[i] = minn + 1

		//dp[i] = 100000
		//for j := 1; j*j <= i; j++ {
		//	dp[i]= min(dp[i], dp[i-j*j]+1)
		//}
	}

	return dp[n]
}

//func numSquares(n int) int {
//	memo := make(map[int]int)
//	var dp func(num int) int
//	dp = func(num int) int {
//		if v, ok := memo[num]; ok {
//			return v
//		}
//
//		if num == 0 {
//			return 0
//		}
//
//		var rlt int = 100000
//		for i := 1; i*i <= num; i++ {
//			rlt = min(rlt, dp(num-i*i)+1)
//		}
//		memo[num] = rlt
//		return rlt
//	}
//	return dp(n)
//}

func sqrt279(num int) int {
	return int(math.Sqrt(float64(num)))
}

//func numSquares(n int) int {
//	maxRow := sqrt279(n)
//	dp := make([][]int, maxRow+1)
//	for i := 0; i < len(dp); i++ {
//		dp[i] = make([]int, n+1)
//	}
//
//	for j := 1; j < len(dp[0]); j++ {
//		dp[0][j] = 100000
//	}
//
//	for i := 1; i < len(dp); i++ {
//		for j := 1; j < len(dp[0]); j++ {
//			minn := dp[i-1][j]
//			if j-i*i >= 0 {
//				minn = min(minn, dp[i][j-i*i]+1)
//			}
//			dp[i][j] = minn
//
//			//if j < i*i {
//			//	dp[i][j] = dp[i-1][j] // 只能不选
//			//} else {
//			//	dp[i][j] = min(dp[i-1][j], dp[i][j-i*i]+1) // 不选 vs 选
//			//}
//
//		}
//	}
//
//	return dp[maxRow][n]
//}

//leetcode submit region end(Prohibit modification and deletion)
