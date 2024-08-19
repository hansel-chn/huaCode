//You are given an m x n integer array grid where grid[i][j] could be:
//
//
// 1 representing the starting square. There is exactly one starting square.
// 2 representing the ending square. There is exactly one ending square.
// 0 representing empty squares we can walk over.
// -1 representing obstacles that we cannot walk over.
//
//
// Return the number of 4-directional walks from the starting square to the
//ending square, that walk over every non-obstacle square exactly once.
//
//
// Example 1:
//
//
//Input: grid = [[1,0,0,0],[0,0,0,0],[0,0,2,-1]]
//Output: 2
//Explanation: We have the following two paths:
//1. (0,0),(0,1),(0,2),(0,3),(1,3),(1,2),(1,1),(1,0),(2,0),(2,1),(2,2)
//2. (0,0),(1,0),(2,0),(2,1),(1,1),(0,1),(0,2),(0,3),(1,3),(1,2),(2,2)
//
//
// Example 2:
//
//
//Input: grid = [[1,0,0,0],[0,0,0,0],[0,0,0,2]]
//Output: 4
//Explanation: We have the following four paths:
//1. (0,0),(0,1),(0,2),(0,3),(1,3),(1,2),(1,1),(1,0),(2,0),(2,1),(2,2),(2,3)
//2. (0,0),(0,1),(1,1),(1,0),(2,0),(2,1),(2,2),(1,2),(0,2),(0,3),(1,3),(2,3)
//3. (0,0),(1,0),(2,0),(2,1),(2,2),(1,2),(1,1),(0,1),(0,2),(0,3),(1,3),(2,3)
//4. (0,0),(1,0),(2,0),(2,1),(1,1),(0,1),(0,2),(0,3),(1,3),(1,2),(2,2),(2,3)
//
//
// Example 3:
//
//
//Input: grid = [[0,1],[2,0]]
//Output: 0
//Explanation: There is no path that walks over every empty square exactly once.
//
//Note that the starting and ending square can be anywhere in the grid.
//
//
//
// Constraints:
//
//
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 20
// 1 <= m * n <= 20
// -1 <= grid[i][j] <= 2
// There is exactly one starting cell and one ending cell.
//
//
// Related Topics 位运算 数组 回溯 矩阵 👍 361 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
package main

func uniquePathsIII(grid [][]int) int {
	var dirs = [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	var start [2]int
	var expectCount int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				start = [2]int{i, j}
			}

			if grid[i][j] == 1 || grid[i][j] == 0 {
				expectCount++
			}
		}
	}

	var rlt int
	var count int
	var traverse func(x, y int)
	traverse = func(x, y int) {
		if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) {
			return
		}

		if grid[x][y] == 2 {
			if count == expectCount {
				rlt++
			}
			return
		}

		if grid[x][y] == -1 {
			return
		}

		for _, dir := range dirs {
			grid[x][y] = -1
			count++
			traverse(x+dir[0], y+dir[1])
			grid[x][y] = 1
			count--
		}
	}
	traverse(start[0], start[1])
	return rlt
}

//leetcode submit region end(Prohibit modification and deletion)
