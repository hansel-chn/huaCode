//Given a m x n matrix mat and an integer k, return a matrix answer where each
//answer[i][j] is the sum of all elements mat[r][c] for:
//
//
// i - k <= r <= i + k,
// j - k <= c <= j + k, and
// (r, c) is a valid position in the matrix.
//
//
//
// Example 1:
//
//
//Input: mat = [[1,2,3],[4,5,6],[7,8,9]], k = 1
//Output: [[12,21,16],[27,45,33],[24,39,28]]
//
//
// Example 2:
//
//
//Input: mat = [[1,2,3],[4,5,6],[7,8,9]], k = 2
//Output: [[45,45,45],[45,45,45],[45,45,45]]
//
//
//
// Constraints:
//
//
// m == mat.length
// n == mat[i].length
// 1 <= m, n, k <= 100
// 1 <= mat[i][j] <= 100
//
//
// Related Topics 数组 矩阵 前缀和 👍 207 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
package main

func matrixBlockSum(mat [][]int, k int) [][]int {
	preMat := buildPreMatrix1314(mat)

	row := len(mat)
	col := len(mat[0])

	rlt := make([][]int, row)
	for i := range rlt {
		rlt[i] = make([]int, col)
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			l, r, t, b := getBound1314(i, j, k, row, col)

			if t == 0 && l == 0 {
				rlt[i][j] = preMat[b][r]
				continue
			}
			if t == 0 {
				rlt[i][j] = preMat[b][r] - preMat[b][l-1]
				continue
			}
			if l == 0 {
				rlt[i][j] = preMat[b][r] - preMat[t-1][r]
				continue
			}
			rlt[i][j] = preMat[b][r] - preMat[b][l-1] - preMat[t-1][r] + preMat[t-1][l-1]
		}
	}
	return rlt
}
func getBound1314(x, y, k, row, col int) (l, r, t, b int) {
	if x-k < 0 {
		t = 0
	} else {
		t = x - k
	}

	if x+k >= row {
		b = row - 1
	} else {
		b = x + k
	}

	if y-k < 0 {
		l = 0
	} else {
		l = y - k
	}

	if y+k >= col {
		r = col-1
	} else {
		r = y + k
	}

	return
}

func buildPreMatrix1314(mat [][]int) [][]int {
	row := len(mat)
	col := len(mat[0])

	preMat := make([][]int, row)
	for i := range preMat {
		preMat[i] = make([]int, col)
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if i == 0 && j == 0 {
				preMat[i][j] = mat[i][j]
				continue
			}
			if i == 0 {
				preMat[i][j] = preMat[i][j-1] + mat[i][j]
				continue
			}
			if j == 0 {
				preMat[i][j] = preMat[i-1][j] + mat[i][j]
				continue
			}
			preMat[i][j] = preMat[i-1][j] + preMat[i][j-1] - preMat[i-1][j-1] + mat[i][j]
		}
	}
	return preMat
}

//leetcode submit region end(Prohibit modification and deletion)
