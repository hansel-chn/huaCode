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

//// Test where to exclude illegitimate choices
//func solveNQueens(n int) [][]string {
//	res := [][]string{}
//	board := make([]string, n)
//	for i := 0; i < n; i++ {
//		board[i] = strings.Repeat(".", n)
//	}
//
//	testBoard  := make([]string, n)
//	copy(testBoard, board)
//	backtrack(board, 0, &res, 0, 0, testBoard)
//	return res
//}
//
//// 路径：board 中小于 row 的那些行都已经成功放置了皇后
//// 选择列表：第 row 行的所有列都是放置皇后的选择
//// 结束条件：row 超过 board 的最后一行
//func backtrack(board []string, row int, res *[][]string, testRow, testCol int, testBoard []string) {
//	if !isValid(testBoard, testRow, testCol) {
//		return
//	}
//
//	// 触发结束条件
//	if row == len(board) {
//		newRow := make([]string, len(board))
//		copy(newRow, board)
//		*res = append(*res, newRow)
//		return
//	}
//
//	n := len(board[row])
//	for col := 0; col < n; col++ {
//		//// 排除不合法选择
//		//if !isValid(board, row, col) {
//		//	continue
//		//}
//		// 做选择
//		testRow = row
//		testCol = col
//		testBoard  = make([]string, n)
//		copy(testBoard, board)
//
//		newLine := []byte(board[row])
//		newLine[col] = 'Q'
//		board[row] = string(newLine)
//		// 进入下一行决策
//		backtrack(board, row+1, res, testRow, testCol, testBoard)
//		// 撤销选择
//		newLine[col] = '.'
//		board[row] = string(newLine)
//	}
//}
//
//func isValid(board []string, row, col int) bool {
//	n := len(board)
//	// 检查列是否有皇后冲突
//	for i := 0; i < n; i++ {
//		if board[i][col] == 'Q' {
//			return false
//		}
//	}
//	// 检查右上方是否有皇后冲突
//	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
//		if board[i][j] == 'Q' {
//			return false
//		}
//	}
//	// 检查左上方是否有皇后冲突
//	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
//		if board[i][j] == 'Q' {
//			return false
//		}
//	}
//	return true
//}

//leetcode submit region end(Prohibit modification and deletion)
