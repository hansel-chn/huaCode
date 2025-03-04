//Given a positive integer n, generate an n x n matrix filled with elements
//from 1 to n² in spiral order.
//
//
// Example 1:
//
//
//Input: n = 3
//Output: [[1,2,3],[8,9,4],[7,6,5]]
//
//
// Example 2:
//
//
//Input: n = 1
//Output: [[1]]
//
//
//
// Constraints:
//
//
// 1 <= n <= 20
//
//
// Related Topics 数组 矩阵 模拟 👍 1418 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
package main

func generateMatrix(n int) [][]int {
	var idx int = 1
	sum := n * n

	rlt := make([][]int, n)
	for i := 0; i < n; i++ {
		rlt[i] = make([]int, n)
	}

	l := 0
	t := 0
	r := n - 1
	b := n - 1
	for idx <= sum {
		if idx <= sum {
			for j := l; j <= r; j++ {
				rlt[t][j] = idx
				idx++
			}
			t++
		}

		if idx <= sum {
			for i := t; i <= b; i++ {
				rlt[i][r] = idx
				idx++
			}
			r--
		}

		if idx <= sum {
			for j := r; j >= l; j-- {
				rlt[b][j] = idx
				idx++
			}
			b--
		}

		if idx <= sum {
			for i := b; i >= t; i-- {
				rlt[i][l] = idx
				idx++
			}
			l++
		}
	}

	return rlt
}

//leetcode submit region end(Prohibit modification and deletion)
