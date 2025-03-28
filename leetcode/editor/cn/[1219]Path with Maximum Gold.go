//In a gold mine grid of size m x n, each cell in this mine has an integer
//representing the amount of gold in that cell, 0 if it is empty.
//
// Return the maximum amount of gold you can collect under the conditions:
//
//
// Every time you are located in a cell you will collect all the gold in that
//cell.
// From your position, you can walk one step to the left, right, up, or down.
// You can't visit the same cell more than once.
// Never visit a cell with 0 gold.
// You can start and stop collecting gold from any position in the grid that
//has some gold.
//
//
//
// Example 1:
//
//
//Input: grid = [[0,6,0],[5,8,7],[0,9,0]]
//Output: 24
//Explanation:
//[[0,6,0],
// [5,8,7],
// [0,9,0]]
//Path to get the maximum gold, 9 -> 8 -> 7.
//
//
// Example 2:
//
//
//Input: grid = [[1,0,7],[2,0,6],[3,4,5],[0,3,0],[9,0,20]]
//Output: 28
//Explanation:
//[[1,0,7],
// [2,0,6],
// [3,4,5],
// [0,3,0],
// [9,0,20]]
//Path to get the maximum gold, 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7.
//
//
//
// Constraints:
//
//
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 15
// 0 <= grid[i][j] <= 100
// There are at most 25 cells containing gold.
//
//
// Related Topics 数组 回溯 矩阵 👍 256 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
package main

func getMaximumGold(grid [][]int) int {
	maxGold := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] > 0 {
				maxGold = max(maxGold, dfs1219(i, j, grid))
			}
		}
	}
	return maxGold
}

func dfs1219(i, j int, grid [][]int) int {
	maxX, maxY := len(grid)-1, len(grid[0])-1

	cur := 0
	maxGold := 0

	var dfs func(x, y int)
	dfs = func(x, y int) {
		cur += grid[x][y]
		store := grid[x][y]
		grid[x][y] = 0
		defer func() {
			grid[x][y] = store
			cur -= grid[x][y]
		}()

		dirs := direct1219(x, y, maxX, maxY, grid)
		if len(dirs) == 0 {
			maxGold = max(maxGold, cur)
			return
		}

		for _, dir := range dirs {
			dfs(dir[0], dir[1])
		}
		return
	}
	dfs(i, j)
	return maxGold
}

var dirs1219 = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func direct1219(i, j int, maxX, maxY int, grid [][]int) [][]int {
	rlt := make([][]int, 0)
	for _, dir := range dirs1219 {
		curX := i + dir[0]
		curY := j + dir[1]
		if curX >= 0 && curX <= maxX && curY >= 0 && curY <= maxY {
			if grid[curX][curY] > 0 {
				rlt = append(rlt, []int{curX, curY})
			}
		}
	}
	return rlt
}

//func getMaximumGold(grid [][]int) int {
//	maxR, maxC := len(grid), len(grid[0])
//
//	maxGold := 0
//
//	var dfs func(i, j, gold int)
//	dfs = func(i, j, gold int) {
//		store := grid[i][j]
//		curGold := gold + store
//		grid[i][j] = 0
//		maxGold = max(maxGold, curGold)
//
//		for _, dir := range dirs1219 {
//			curI := i + dir[0]
//			curJ := j + dir[1]
//			if curI >= 0 && curI < maxR && curJ >= 0 && curJ < maxC && grid[curI][curJ] > 0 {
//				dfs(curI, curJ, curGold)
//			}
//		}
//		grid[i][j] = store
//		curGold -= store
//	}
//
//	for i := 0; i < len(grid); i++ {
//		for j := 0; j < len(grid[0]); j++ {
//			if grid[i][j] > 0 {
//				dfs(i, j, 0)
//			}
//		}
//	}
//	return maxGold
//}

//leetcode submit region end(Prohibit modification and deletion)
