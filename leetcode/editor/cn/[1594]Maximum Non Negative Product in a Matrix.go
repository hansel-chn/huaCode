//You are given a m x n matrix grid. Initially, you are located at the top-left
//corner (0, 0), and in each step, you can only move right or down in the matrix.
//
//
// Among all possible paths starting from the top-left corner (0, 0) and ending
//in the bottom-right corner (m - 1, n - 1), find the path with the maximum non-
//negative product. The product of a path is the product of all integers in the
//grid cells visited along the path.
//
// Return the maximum non-negative product modulo 10⁹ + 7. If the maximum
//product is negative, return -1.
//
// Notice that the modulo is performed after getting the maximum product.
//
//
// Example 1:
//
//
//Input: grid = [[-1,-2,-3],[-2,-3,-3],[-3,-3,-2]]
//Output: -1
//Explanation: It is not possible to get non-negative product in the path from (
//0, 0) to (2, 2), so return -1.
//
//
// Example 2:
//
//
//Input: grid = [[1,-2,1],[1,-2,1],[3,-4,1]]
//Output: 8
//Explanation: Maximum non-negative product is shown (1 * 1 * -2 * -4 * 1 = 8).
//
//
// Example 3:
//
//
//Input: grid = [[1,3],[0,-4]]
//Output: 0
//Explanation: Maximum non-negative product is shown (1 * 0 * -4 = 0).
//
//
//
// Constraints:
//
//
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 15
// -4 <= grid[i][j] <= 4
//
//
// Related Topics 数组 动态规划 矩阵 👍 62 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
package main

import (
	"math"
)

func maxProductPath(grid [][]int) int {
	dp := make([][][2]int, len(grid))
	for i := range dp {
		dp[i] = make([][2]int, len(grid[0]))
	}

	for i := range dp {
		if i == 0 {
			dp[i][0] = [2]int{grid[i][0], grid[i][0]}
			continue
		}
		dp[i][0] = [2]int{dp[i-1][0][0] * grid[i][0], dp[i-1][0][0] * grid[i][0]}
	}

	for i := range dp[0] {
		if i == 0 {
			dp[0][i] = [2]int{grid[0][i], grid[0][i]}
			continue
		}
		dp[0][i] = [2]int{dp[0][i-1][0] * grid[0][i], dp[0][i-1][0] * grid[0][i]}
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			dp[i][j][0] = max1594(dp[i][j-1][0]*grid[i][j], dp[i][j-1][1]*grid[i][j], dp[i-1][j][0]*grid[i][j], dp[i-1][j][1]*grid[i][j])
			dp[i][j][1] = min1594(dp[i][j-1][0]*grid[i][j], dp[i][j-1][1]*grid[i][j], dp[i-1][j][0]*grid[i][j], dp[i-1][j][1]*grid[i][j])
		}
	}

	rlt := max(dp[len(grid)-1][len(grid[0])-1][0], dp[len(grid)-1][len(grid[0])-1][1])
	if rlt < 0 {
		return -1
	} else {
		return handleVal( rlt)
	}
}

func max1594(vals ...int) int {
	rlt := math.MinInt
	for _, val := range vals {
		rlt = max(rlt, val)
	}
	return rlt
}

func min1594(vals ...int) int {
	rlt := math.MaxInt
	for _, val := range vals {
		rlt = min(rlt, val)
	}
	return rlt
}

func handleVal(val int)int  {
	return val%(1e9+7)
}

//leetcode submit region end(Prohibit modification and deletion)
