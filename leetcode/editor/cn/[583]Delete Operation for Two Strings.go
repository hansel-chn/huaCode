//Given two strings word1 and word2, return the minimum number of steps
//required to make word1 and word2 the same.
//
// In one step, you can delete exactly one character in either string.
//
//
// Example 1:
//
//
//Input: word1 = "sea", word2 = "eat"
//Output: 2
//Explanation: You need one step to make "sea" to "ea" and another step to make
//"eat" to "ea".
//
//
// Example 2:
//
//
//Input: word1 = "leetcode", word2 = "etco"
//Output: 4
//
//
//
// Constraints:
//
//
// 1 <= word1.length, word2.length <= 500
// word1 and word2 consist of only lowercase English letters.
//
//
// Related Topics 字符串 动态规划 👍 685 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
package main

// func minDistance(word1 string, word2 string) int {
func minDistance583(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
	}

	for i := 0; i < len(word1)+1; i++ {
		dp[i][0] = i
	}

	for j := 0; j < len(word2)+1; j++ {
		dp[0][j] = j
	}

	for i := 1; i < len(word1)+1; i++ {
		for j := 1; j < len(word2)+1; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + 1
			}
		}
	}
	return dp[len(word1)][len(word2)]
}

//leetcode submit region end(Prohibit modification and deletion)
