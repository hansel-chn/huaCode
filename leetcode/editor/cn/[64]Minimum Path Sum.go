/**
Given a m x n grid filled with non-negative numbers, find a path from top left
to bottom right, which minimizes the sum of all numbers along its path.

 Note: You can only move either down or right at any point in time.


 Example 1:


Input: grid = [[1,3,1],[1,5,1],[4,2,1]]
Output: 7
Explanation: Because the path 1 → 3 → 1 → 1 → 1 minimizes the sum.


 Example 2:


Input: grid = [[1,2,3],[4,5,6]]
Output: 12



 Constraints:


 m == grid.length
 n == grid[i].length
 1 <= m, n <= 200
 0 <= grid[i][j] <= 200


 Related Topics 数组 动态规划 矩阵 👍 1649 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
package main

//func minPathSum(grid [][]int) int {
//	row := len(grid)
//	col := len(grid[0])
//
//	dp := make([][]int, row+1)
//	for i := range dp {
//		dp[i] = make([]int, col+1)
//		dp[i][0] = math.MaxInt
//	}
//	for i := 1; i < len(dp[0]); i++ {
//		dp[0][i] = math.MaxInt
//	}
//	dp[1][0] = 0
//	dp[0][1] = 0
//
//	for i := 1; i < row+1; i++ {
//		for j := 1; j < col+1; j++ {
//			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i-1][j-1]
//		}
//	}
//	return dp[row][col]
//}

func minPathSum(grid [][]int) int {
	memo := make(map[[2]int]int)
	row := len(grid)
	col := len(grid[0])
	var dp func(grid [][]int, i, j int) int
	dp = func(grid [][]int, i, j int) int {
		if v, ok := memo[[2]int{i, j}]; ok {
			return v
		}

		if i == 0&&j==0 {
			return  grid[i][j]
		}

		if i == 0 {
			rlt := dp(grid, i, j-1) + grid[i][j]
			memo[[2]int{i, j}] = rlt
			return rlt
		}
		if j == 0 {
			rlt := dp(grid, i-1, j) + grid[i][j]
			memo[[2]int{i, j}] = rlt
			return rlt
		}
		rlt := min(dp(grid, i-1, j), dp(grid, i, j-1)) + grid[i][j]
		memo[[2]int{i, j}] = rlt
		return rlt
	}
	return dp(grid, row-1, col-1)
}

//leetcode submit region end(Prohibit modification and deletion)
