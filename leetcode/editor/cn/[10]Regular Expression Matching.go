/**
Given an input string s and a pattern p, implement regular expression matching
with support for '.' and '*' where:


 '.' Matches any single character.
 '*' Matches zero or more of the preceding element.


 The matching should cover the entire input string (not partial).


 Example 1:


Input: s = "aa", p = "a"
Output: false
Explanation: "a" does not match the entire string "aa".


 Example 2:


Input: s = "aa", p = "a*"
Output: true
Explanation: '*' means zero or more of the preceding element, 'a'. Therefore,
by repeating 'a' once, it becomes "aa".


 Example 3:


Input: s = "ab", p = ".*"
Output: true
Explanation: ".*" means "zero or more (*) of any character (.)".



 Constraints:


 1 <= s.length <= 20
 1 <= p.length <= 20
 s contains only lowercase English letters.
 p contains only lowercase English letters, '.', and '*'.
 It is guaranteed for each appearance of the character '*', there will be a
previous valid character to match.


 Related Topics 递归 字符串 动态规划 👍 3854 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
package main

func isMatch(s string, p string) bool {
	dp := make([][]bool, len(s)+1)
	for i := range dp {
		dp[i] = make([]bool, len(p)+1)
	}

	for i := 0; i < len(s)+1; i++ {
		dp[i][0] = false
	}

	dp[0][0] = true

	for i := 0; i < len(s)+1; i++ {
		for j := 1; j < len(p)+1; j++ {
			if match(i, j, s, p) {
				dp[i][j] = dp[i][j]||dp[i-1][j-1]
			} else if p[j-1] == '*' {
				dp[i][j] = dp[i][j] || dp[i][j-2]
				if match(i, j-1, s, p) {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			}
		}
	}

	return dp[len(s)][len(p)]
}

func match(i, j int, s, p string) bool {
	if i == 0 {
		return false
	}

	return s[i-1] == p[j-1] || p[j-1] == '.'

}

//leetcode submit region end(Prohibit modification and deletion)
