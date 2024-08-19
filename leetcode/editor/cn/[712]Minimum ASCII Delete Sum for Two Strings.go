//Given two strings s1 and s2, return the lowest ASCII sum of deleted
//characters to make two strings equal.
//
//
// Example 1:
//
//
//Input: s1 = "sea", s2 = "eat"
//Output: 231
//Explanation: Deleting "s" from "sea" adds the ASCII value of "s" (115) to the
//sum.
//Deleting "t" from "eat" adds 116 to the sum.
//At the end, both strings are equal, and 115 + 116 = 231 is the minimum sum
//possible to achieve this.
//
//
// Example 2:
//
//
//Input: s1 = "delete", s2 = "leet"
//Output: 403
//Explanation: Deleting "dee" from "delete" to turn the string into "let",
//adds 100[d] + 101[e] + 101[e] to the sum.
//Deleting "e" from "leet" adds 101[e] to the sum.
//At the end, both strings are equal to "let", and the answer is 100+101+101+101
// = 403.
//If instead we turned both strings into "lee" or "eet", we would get answers
//of 433 or 417, which are higher.
//
//
//
// Constraints:
//
//
// 1 <= s1.length, s2.length <= 1000
// s1 and s2 consist of lowercase English letters.
//
//
// Related Topics 字符串 动态规划 👍 391 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
package main

func minimumDeleteSum(s1 string, s2 string) int {
	dp := make([][]int, len(s1)+1)
	for i := range dp {
		dp[i] = make([]int, len(s2)+1)
	}

	for i := 1; i < len(s1)+1; i++ {
		dp[i][0] = dp[i-1][0] + int(s1[i-1])
	}

	for j := 1; j < len(s2)+1; j++ {
		dp[0][j] = dp[0][j-1] + int(s2[j-1])
	}

	for i := 1; i < len(s1)+1; i++ {
		for j := 1; j < len(s2)+1; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j]+int(s1[i-1]), dp[i][j-1]+int(s2[j-1]))
			}
		}
	}
	return dp[len(s1)][len(s2)]
}

//leetcode submit region end(Prohibit modification and deletion)
