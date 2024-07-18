/**
Given two strings s and t, return the number of distinct subsequences of s
which equals t.

 The test cases are generated so that the answer fits on a 32-bit signed
integer.


 Example 1:


Input: s = "rabbbit", t = "rabbit"
Output: 3
Explanation:
As shown below, there are 3 ways you can generate "rabbit" from s.
rabbbit
rabbbit
rabbbit


 Example 2:


Input: s = "babgbag", t = "bag"
Output: 5
Explanation:
As shown below, there are 5 ways you can generate "bag" from s.
babgbag
babgbag
babgbag
babgbag
babgbag


 Constraints:


 1 <= s.length, t.length <= 1000
 s and t consist of English letters.


 Related Topics 字符串 动态规划 👍 1245 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
package main

func numDistinct(s string, t string) int {
	dp := make([][]int, len(s))
	for i := range dp {
		dp[i] = make([]int, len(t))
	}

	for j := 0; j < len(s); j++ {
		if j == 0 {
			if s[j] == t[0] {
				dp[j][0] = 1
			} else {
				dp[j][0] = 0
			}
			continue
		}

		if s[j] == t[0] {
			dp[j][0] = dp[j-1][0] + 1
		} else {
			dp[j][0] = dp[j-1][0]
		}
	}

	for j := 1; j < len(t); j++ {
		for i := j; i < len(s); i++ {
			count := dp[i-1][j]
			if s[i] == t[j] {
				count += dp[i-1][j-1]
			}
			dp[i][j] = count
		}
	}

	return dp[len(s)-1][len(t)-1]
}

//leetcode submit region end(Prohibit modification and deletion)
