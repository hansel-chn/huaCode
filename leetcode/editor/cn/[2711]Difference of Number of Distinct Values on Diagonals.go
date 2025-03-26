//Given a 2D grid of size m x n, you should find the matrix answer of size m x
//n.
//
// The cell answer[r][c] is calculated by looking at the diagonal values of the
//cell grid[r][c]:
//
//
// Let leftAbove[r][c] be the number of distinct values on the diagonal to the
//left and above the cell grid[r][c] not including the cell grid[r][c] itself.
// Let rightBelow[r][c] be the number of distinct values on the diagonal to the
//right and below the cell grid[r][c], not including the cell grid[r][c] itself.
// Then answer[r][c] = |leftAbove[r][c] - rightBelow[r][c]|.
//
//
// A matrix diagonal is a diagonal line of cells starting from some cell in
//either the topmost row or leftmost column and going in the bottom-right direction
//until the end of the matrix is reached.
//
//
// For example, in the below diagram the diagonal is highlighted using the cell
//with indices (2, 3) colored gray:
//
//
//
// Red-colored cells are left and above the cell.
// Blue-colored cells are right and below the cell.
//
//
//
//
//
//
// Return the matrix answer.
//
//
// Example 1:
//
//
// Input: grid = [[1,2,3],[3,1,5],[3,2,1]]
//
//
// Output: Output: [[1,1,0],[1,0,1],[0,1,1]]
//
// Explanation:
//
// To calculate the answer cells:
//
//
//
//
// answer
// left-above elements
// leftAbove
// right-below elements
// rightBelow
// |leftAbove - rightBelow|
//
//
//
//
// [0][0]
// []
// 0
// [grid[1][1], grid[2][2]]
// |{1, 1}| = 1
// 1
//
//
// [0][1]
// []
// 0
// [grid[1][2]]
// |{5}| = 1
// 1
//
//
// [0][2]
// []
// 0
// []
// 0
// 0
//
//
// [1][0]
// []
// 0
// [grid[2][1]]
// |{2}| = 1
// 1
//
//
// [1][1]
// [grid[0][0]]
// |{1}| = 1
// [grid[2][2]]
// |{1}| = 1
// 0
//
//
// [1][2]
// [grid[0][1]]
// |{2}| = 1
// []
// 0
// 1
//
//
// [2][0]
// []
// 0
// []
// 0
// 0
//
//
// [2][1]
// [grid[1][0]]
// |{3}| = 1
// []
// 0
// 1
//
//
// [2][2]
// [grid[0][0], grid[1][1]]
// |{1, 1}| = 1
// []
// 0
// 1
//
//
//
//
// Example 2:
//
//
// Input: grid = [[1]]
//
//
// Output: Output: [[0]]
//
//
// Constraints:
//
//
// m == grid.length
// n == grid[i].length
// 1 <= m, n, grid[i][j] <= 50
//
//
// Related Topics 数组 哈希表 矩阵 👍 40 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
package main

import (
	"math/bits"
)

func differenceOfDistinctValues(grid [][]int) [][]int {
	m := len(grid)
	n := len(grid[0])

	rlt := make([][]int, m)
	for i := range rlt {
		rlt[i] = make([]int, n)
	}

	if m == 1 || n == 1 {
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				rlt[i][j] = 0
			}
		}
		return rlt
	}

	lt := make([][]int, m)
	for i := range lt {
		lt[i] = make([]int, n)
	}

	rb := make([][]int, m)
	for i := range rb {
		rb[i] = make([]int, n)
	}

	for i := -m + 1; i < n; i++ {
		var row, col int
		if i <= 0 {
			row, col = -i, 0
		} else {
			row, col = 0, i
		}

		var ltDiag int
		for col < n && row < m {
			ltDiag |= 1 << grid[row][col]
			lt[row][col] = bits.OnesCount(uint(ltDiag))
			row++
			col++
		}
	}

	for i := -m + 1; i < n; i++ {
		var row, col int
		if i <= n-m {
			row = m - 1
			col = row + i
		} else {
			col = n - 1
			row = col - i
		}

		var rbDiag int
		for col >= 0 && row >= 0 {
			rbDiag |= 1 << grid[row][col]
			rb[row][col] = bits.OnesCount(uint(rbDiag))
			row--
			col--
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (i == m-1 && j == 0) || (i == 0 && j == n-1) {
				rlt[i][j] = 0
				continue
			}
			if i == m-1 || j == n-1 {
				rlt[i][j] = lt[i-1][j-1]
				continue
			}
			if i == 0 || j == 0 {
				rlt[i][j] = rb[i+1][j+1]
				continue
			}

			rlt[i][j] = abs(lt[i-1][j-1] - rb[i+1][j+1])
		}
	}

	return rlt
}

func abs(i int) int {
	if i < 0 {
		return -i
	} else {
		return i
	}
}

//leetcode submit region end(Prohibit modification and deletion)
