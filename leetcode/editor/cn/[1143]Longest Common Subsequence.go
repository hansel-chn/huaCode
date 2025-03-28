//Given two strings text1 and text2, return the length of their longest common
//subsequence. If there is no common subsequence, return 0.
//
// A subsequence of a string is a new string generated from the original string
//with some characters (can be none) deleted without changing the relative order
//of the remaining characters.
//
//
// For example, "ace" is a subsequence of "abcde".
//
//
// A common subsequence of two strings is a subsequence that is common to both
//strings.
//
//
// Example 1:
//
//
//Input: text1 = "abcde", text2 = "ace"
//Output: 3
//Explanation: The longest common subsequence is "ace" and its length is 3.
//
//
// Example 2:
//
//
//Input: text1 = "abc", text2 = "abc"
//Output: 3
//Explanation: The longest common subsequence is "abc" and its length is 3.
//
//
// Example 3:
//
//
//Input: text1 = "abc", text2 = "def"
//Output: 0
//Explanation: There is no such common subsequence, so the result is 0.
//
//
//
// Constraints:
//
//
// 1 <= text1.length, text2.length <= 1000
// text1 and text2 consist of only lowercase English characters.
//
//
// Related Topics 字符串 动态规划 👍 1628 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
package main

func longestCommonSubsequence(text1 string, text2 string) int {
	row := len(text1) + 1
	col := len(text2) + 1
	dp := make([][]int, row)
	for i := range dp {
		dp[i] = make([]int, col)
	}

	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[row-1][col-1]
}

//leetcode submit region end(Prohibit modification and deletion)
