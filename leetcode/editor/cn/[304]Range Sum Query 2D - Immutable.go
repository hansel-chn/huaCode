/**
Given a 2D matrix matrix, handle multiple queries of the following type:


 Calculate the sum of the elements of matrix inside the rectangle defined by
its upper left corner (row1, col1) and lower right corner (row2, col2).


 Implement the NumMatrix class:


 NumMatrix(int[][] matrix) Initializes the object with the integer matrix
matrix.
 int sumRegion(int row1, int col1, int row2, int col2) Returns the sum of the
elements of matrix inside the rectangle defined by its upper left corner (row1,
col1) and lower right corner (row2, col2).


 You must design an algorithm where sumRegion works on O(1) time complexity.


 Example 1:


Input
["NumMatrix", "sumRegion", "sumRegion", "sumRegion"]
[[[[3, 0, 1, 4, 2], [5, 6, 3, 2, 1], [1, 2, 0, 1, 5], [4, 1, 0, 1, 7], [1, 0, 3,
 0, 5]]], [2, 1, 4, 3], [1, 1, 2, 2], [1, 2, 2, 4]]
Output
[null, 8, 11, 12]


Explanation
NumMatrix numMatrix = new NumMatrix([[3, 0, 1, 4, 2], [5, 6, 3, 2, 1], [1, 2, 0,
 1, 5], [4, 1, 0, 1, 7], [1, 0, 3, 0, 5]]);
numMatrix.sumRegion(2, 1, 4, 3); // return 8 (i.e sum of the red rectangle)
numMatrix.sumRegion(1, 1, 2, 2); // return 11 (i.e sum of the green rectangle)
numMatrix.sumRegion(1, 2, 2, 4); // return 12 (i.e sum of the blue rectangle)



 Constraints:


 m == matrix.length
 n == matrix[i].length
 1 <= m, n <= 200
 -10⁴ <= matrix[i][j] <= 10⁴
 0 <= row1 <= row2 < m
 0 <= col1 <= col2 < n
 At most 10⁴ calls will be made to sumRegion.


 Related Topics 设计 数组 矩阵 前缀和 👍 563 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
package main

type NumMatrix struct {
	preMatrix [][]int
}

//func Constructor(matrix [][]int) NumMatrix {
//	preMatrix := make([][]int, len(matrix)+1)
//	for i := range preMatrix {
//		preMatrix[i] = make([]int, len(matrix[0])+1)
//	}
//
//	for i := 0; i < len(matrix); i++ {
//		for j := 0; j < len(matrix[i]); j++ {
//			preMatrix[i+1][j+1] = preMatrix[i][j+1] + preMatrix[i+1][j] - preMatrix[i][j] + matrix[i][j]
//		}
//	}
//
//	return NumMatrix{preMatrix}
//}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.preMatrix[row2+1][col2+1] - this.preMatrix[row2+1][col1] - this.preMatrix[row1][col2+1] + this.preMatrix[row1][col1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
//leetcode submit region end(Prohibit modification and deletion)
