//Given an m x n matrix, return all elements of the matrix in spiral order.
//
//
// Example 1:
//
//
//Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
//Output: [1,2,3,6,9,8,7,4,5]
//
//
// Example 2:
//
//
//Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
//Output: [1,2,3,4,8,12,11,10,9,5,6,7]
//
//
//
// Constraints:
//
//
// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 10
// -100 <= matrix[i][j] <= 100
//
//
// Related Topics 数组 矩阵 模拟 👍 1874 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
package main

func spiralOrder(matrix [][]int) []int {
	var l, r, t, b = 0, len(matrix[0]) - 1, 0, len(matrix) - 1

	rlt := make([]int, 0)
	num:=len(matrix[0])*len(matrix)
	for len(rlt) <num {
		if len(rlt) <num {
			for j := l; j <= r; j++ {
				rlt = append(rlt, matrix[t][j])
			}
			t++
		}

		if len(rlt) <num {
			for i := t; i <= b; i++ {
				rlt = append(rlt, matrix[i][r])
			}
			r--
		}

		if len(rlt) <num {
			for j := r; j >= l; j-- {
				rlt = append(rlt, matrix[b][j])
			}
			b--
		}

		if len(rlt) <num {
			for i := b; i >= t; i-- {
				rlt = append(rlt, matrix[i][l])
			}
			l++
		}
	}

	return rlt
}

//leetcode submit region end(Prohibit modification and deletion)
