/**
You are given an m x n matrix board containing letters 'X' and 'O', capture
regions that are surrounded:


 Connect: A cell is connected to adjacent cells horizontally or vertically.
 Region: To form a region connect every 'O' cell.
 Surround: The region is surrounded with 'X' cells if you can connect the
region with 'X' cells and none of the region cells are on the edge of the board.


 A surrounded region is captured by replacing all 'O's with 'X's in the input
matrix board.


 Example 1:


 Input: board = [["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O",
"X","X"]]


 Output: [["X","X","X","X"],["X","X","X","X"],["X","X","X","X"],["X","O","X",
"X"]]

 Explanation:

 In the above diagram, the bottom region is not captured because it is on the
edge of the board and cannot be surrounded.

 Example 2:


 Input: board = [["X"]]


 Output: [["X"]]


 Constraints:


 m == board.length
 n == board[i].length
 1 <= m, n <= 200
 board[i][j] is 'X' or 'O'.


 Related Topics 深度优先搜索 广度优先搜索 并查集 数组 矩阵 👍 1129 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
package main

func solve(board [][]byte) {
	for i := 0; i < len(board); i++ {
		if board[i][0] == 'O' {
			dfs(i, 0, board)
		}
		if board[i][len(board[0])-1] == 'O' {
			dfs(i, len(board[0])-1, board)
		}
	}

	for j := 0; j < len(board[0]); j++ {
		if board[0][j] == 'O' {
			dfs(0, j, board)
		}
		if board[len(board)-1][j] == 'O' {
			dfs(len(board)-1, j, board)
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == '#' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
	return
}

var dirs130 = [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func dfs(i, j int, board [][]byte) {
	if board[i][j] == 'O' {
		board[i][j] = '#'
	} else {
		return
	}

	for _, dir := range dirs130 {
		newI := i + dir[0]
		newJ := j + dir[1]
		if newI >= 0 && newI < len(board) && newJ >= 0 && newJ < len(board[0]) {
			dfs(newI, newJ, board)
		}
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
