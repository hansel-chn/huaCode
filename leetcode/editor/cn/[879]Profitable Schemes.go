//There is a group of n members, and a list of various crimes they could commit.
// The iᵗʰ crime generates a profit[i] and requires group[i] members to
//participate in it. If a member participates in one crime, that member can't participate
//in another crime.
//
// Let's call a profitable scheme any subset of these crimes that generates at
//least minProfit profit, and the total number of members participating in that
//subset of crimes is at most n.
//
// Return the number of schemes that can be chosen. Since the answer may be
//very large, return it modulo 10⁹ + 7.
//
//
// Example 1:
//
//
//Input: n = 5, minProfit = 3, group = [2,2], profit = [2,3]
//Output: 2
//Explanation: To make a profit of at least 3, the group could either commit
//crimes 0 and 1, or just crime 1.
//In total, there are 2 schemes.
//
// Example 2:
//
//
//Input: n = 10, minProfit = 5, group = [2,3,5], profit = [6,7,8]
//Output: 7
//Explanation: To make a profit of at least 5, the group could commit any
//crimes, as long as they commit one.
//There are 7 possible schemes: (0), (1), (2), (0,1), (0,2), (1,2), and (0,1,2).
//
//
//
// Constraints:
//
//
// 1 <= n <= 100
// 0 <= minProfit <= 100
// 1 <= group.length <= 100
// 1 <= group[i] <= 100
// profit.length == group.length
// 0 <= profit[i] <= 100
//
//
// Related Topics 数组 动态规划 👍 314 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
package main

import "fmt"

func profitableSchemes(n int, minProfit int, group []int, profit []int) int {
	dp := make([][][]int, len(group)+1)
	for i := range dp {
		dp[i] = make([][]int, n+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, minProfit+1)
		}
	}

	for i := range dp[0] {
		dp[0][i][0] = 1
	}

	for i := 1; i < len(group)+1; i++ {
		for j := 0; j < n+1; j++ {
			for k := 0; k < minProfit+1; k++ {
				dp[i][j][k] = dp[i-1][j][k]
				if j-group[i-1] >= 0 {
					dp[i][j][k] += dp[i-1][j-group[i-1]][max(0, k-profit[i-1])]
				}
				dp[i][j][k] = dp[i][j][k] % (1e9 + 7)
			}
		}
	}
	fmt.Println(dp)
	return dp[len(group)][n][minProfit]
}

//leetcode submit region end(Prohibit modification and deletion)
