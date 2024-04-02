/**
The n-queens puzzle is the problem of placing n queens on an n x n chessboard
such that no two queens attack each other.

 Given an integer n, return all distinct solutions to the n-queens puzzle. You
may return the answer in any order.

 Each solution contains a distinct board configuration of the n-queens'
placement, where 'Q' and '.' both indicate a queen and an empty space, respectively.


 Example 1:


Input: n = 4
Output: [[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
Explanation: There exist two distinct solutions to the 4-queens puzzle as shown
above


 Example 2:


Input: n = 1
Output: [["Q"]]



 Constraints:


 1 <= n <= 9


 Related Topics 数组 回溯 👍 2049 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
package main

import "strings"

func solveNQueens(n int) [][]string {
	board := make([]string, n)
	for i := 0; i < len(board); i++ {
		board[i] = strings.Repeat(".", n)
	}

	var checkValid func(row, col int) bool
	checkValid = func(row, col int) bool {
		for i := 0; i < row; i++ {
			if board[i][col] != '.' {
				return false
			}
		}

		for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
			if board[i][j] != '.' {
				return false
			}
		}

		for i, j := row-1, col+1; i >= 0 && j < len(board[0]); i, j = i-1, j+1 {
			if board[i][j] != '.' {
				return false
			}
		}
		return true
	}

	rlt := make([][]string, 0)
	var backTrack func(row int)
	backTrack = func(row int) {
		if len(board) == row {
			temp := make([]string, len(board))
			copy(temp, board)
			rlt = append(rlt, temp)
			return
		}

		for col := 0; col < len(board[0]); col++ {
			if checkValid(row, col) == false {
				continue
			}

			newLine := []byte(board[row])
			newLine[col] = 'Q'
			board[row] = string(newLine)

			backTrack(row + 1)

			newLine[col] = '.'
			board[row] = string(newLine)
		}
	}

	backTrack(0)
	return rlt
}

//leetcode submit region end(Prohibit modification and deletion)
