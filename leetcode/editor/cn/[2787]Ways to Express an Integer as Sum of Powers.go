//Given two positive integers n and x.
//
// Return the number of ways n can be expressed as the sum of the xᵗʰ power of
//unique positive integers, in other words, the number of sets of unique integers [
//n1, n2, ..., nk] where n = n1ˣ + n2ˣ + ... + nkˣ.
//
// Since the result can be very large, return it modulo 10⁹ + 7.
//
// For example, if n = 160 and x = 3, one way to express n is n = 2³ + 3³ + 5³.
//
//
//
// Example 1:
//
//
//Input: n = 10, x = 2
//Output: 1
//Explanation: We can express n as the following: n = 3² + 1² = 10.
//It can be shown that it is the only way to express 10 as the sum of the 2ⁿᵈ
//power of unique integers.
//
//
// Example 2:
//
//
//Input: n = 4, x = 1
//Output: 2
//Explanation: We can express n in the following ways:
//- n = 4¹ = 4.
//- n = 3¹ + 1¹ = 4.
//
//
//
// Constraints:
//
//
// 1 <= n <= 300
// 1 <= x <= 5
//
//
// Related Topics 动态规划 👍 21 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
package main

import "math"

func numberOfWays(n int, x int) int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i < n+1; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < n+1; j++ {
			dp[i][j] = dp[i][j-1]
			if i-pow2787(j, x) >= 0 {
				dp[i][j] += dp[i-pow2787(j, x)][j-1]
			}
		}
	}

	return dp[n][n] % (1e9 + 7)
}

func pow2787(base, exp int) int {
	return int(math.Pow(float64(base), float64(exp)))
}

//func  pow2787(x, n int) int {
//	res := 1
//	for ; n > 0; n /= 2 {
//		if n%2 > 0 {
//			res = res * x
//		}
//		x = x * x
//	}
//	return res
//}

//leetcode submit region end(Prohibit modification and deletion)
