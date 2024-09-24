//Given a triangle array, return the minimum path sum from top to bottom.
//
// For each step, you may move to an adjacent number of the row below. More
//formally, if you are on index i on the current row, you may move to either index i
//or index i + 1 on the next row.
//
//
// Example 1:
//
//
//Input: triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
//Output: 11
//Explanation: The triangle looks like:
//   2
//  3 4
// 6 5 7
//4 1 8 3
//The minimum path sum from top to bottom is 2 + 3 + 5 + 1 = 11 (underlined
//above).
//
//
// Example 2:
//
//
//Input: triangle = [[-10]]
//Output: -10
//
//
//
// Constraints:
//
//
// 1 <= triangle.length <= 200
// triangle[0].length == 1
// triangle[i].length == triangle[i - 1].length + 1
// -10⁴ <= triangle[i][j] <= 10⁴
//
//
//
//Follow up: Could you do this using only
//O(n) extra space, where
//n is the total number of rows in the triangle?
//
// Related Topics 数组 动态规划 👍 1373 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
package main

func minimumTotal(triangle [][]int) int {
	dp := make([][]int, len(triangle))
	for i := range dp {
		dp[i] = make([]int, len(triangle))
	}

	for i := len(triangle) - 1; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			if i == len(triangle)-1 {
				dp[i][j] = triangle[i][j]
				continue
			}
			dp[i][j] = min(dp[i+1][j], dp[i+1][j+1]) + triangle[i][j]
		}
	}

	return dp[0][0]
}

//leetcode submit region end(Prohibit modification and deletion)
