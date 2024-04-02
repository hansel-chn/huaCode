/**
Given an m x n grid of characters board and a string word, return true if word
exists in the grid.

 The word can be constructed from letters of sequentially adjacent cells, where
adjacent cells are horizontally or vertically neighboring. The same letter cell
may not be used more than once.


 Example 1:


Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word =
"ABCCED"
Output: true


 Example 2:


Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word =
"SEE"
Output: true


 Example 3:


Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word =
"ABCB"
Output: false



 Constraints:


 m == board.length
 n = board[i].length
 1 <= m, n <= 6
 1 <= word.length <= 15
 board and word consists of only lowercase and uppercase English letters.



 Follow up: Could you use search pruning to make your solution faster with a
larger board?

 Related Topics 数组 字符串 回溯 矩阵 👍 1797 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
package main

func exist(board [][]byte, word string) bool {
	current := make([]byte, 0)
	used := make([][]bool, len(board))
	for i := range used {
		used[i] = make([]bool, len(board[0]))
	}

	var backTrack func(row, col int) bool
	backTrack = func(row, col int) bool {
		if string(current) == word {
			return true
		}

		if string(current) != word[:len(current)] {
			return false
		}

		var moveTo func(row, col int) bool
		moveTo = func(row, col int) bool {
			if row >= 0 && col >= 0 && row < len(board) && col < len(board[0]) && used[row][col] != true {
				current = append(current, board[row][col])
				used[row][col] = true
				if backTrack(row, col) {
					return true
				}
				used[row][col] = false
				current = current[:len(current)-1]
			}
			return false
		}

		if moveTo(row-1, col) {
			return true
		}
		if moveTo(row, col-1) {
			return true
		}
		if moveTo(row+1, col) {
			return true
		}
		if moveTo(row, col+1) {
			return true
		}

		//if row-1 >= 0 && used[row-1][col] != true {
		//	current = append(current, board[row-1][col])
		//	used[row-1][col] = true
		//	if backTrack(row-1, col) {
		//		return true
		//	}
		//	used[row-1][col] = false
		//	current = current[:len(current)-1]
		//}
		//if col-1 >= 0 && used[row][col-1] != true {
		//	current = append(current, board[row][col-1])
		//	used[row][col-1] = true
		//	if backTrack(row, col-1) {
		//		return true
		//	}
		//	current = current[:len(current)-1]
		//	used[row][col-1] = false
		//}
		//if row+1 < len(board) && used[row+1][col] != true {
		//	current = append(current, board[row+1][col])
		//	used[row+1][col] = true
		//	if backTrack(row+1, col) {
		//		return true
		//	}
		//	current = current[:len(current)-1]
		//	used[row+1][col] = false
		//}
		//if col+1 < len(board[0]) && used[row][col+1] != true {
		//	current = append(current, board[row][col+1])
		//	used[row][col+1] = true
		//	if backTrack(row, col+1) {
		//		return true
		//	}
		//	current = current[:len(current)-1]
		//	used[row][col+1] = false
		//}
		return false
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			current = append(current, board[i][j])
			used[i][j] = true
			if backTrack(i, j) {
				return true
			}
			current = current[:len(current)-1]
			used[i][j] = false
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
