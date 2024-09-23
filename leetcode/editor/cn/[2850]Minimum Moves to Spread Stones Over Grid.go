//You are given a 0-indexed 2D integer matrix grid of size 3 * 3, representing
//the number of stones in each cell. The grid contains exactly 9 stones, and there
//can be multiple stones in a single cell.
//
// In one move, you can move a single stone from its current cell to any other
//cell if the two cells share a side.
//
// Return the minimum number of moves required to place one stone in each cell.
//
//
//
// Example 1:
//
//
//Input: grid = [[1,1,0],[1,1,1],[1,2,1]]
//Output: 3
//Explanation: One possible sequence of moves to place one stone in each cell
//is:
//1- Move one stone from cell (2,1) to cell (2,2).
//2- Move one stone from cell (2,2) to cell (1,2).
//3- Move one stone from cell (1,2) to cell (0,2).
//In total, it takes 3 moves to place one stone in each cell of the grid.
//It can be shown that 3 is the minimum number of moves required to place one
//stone in each cell.
//
//
// Example 2:
//
//
//Input: grid = [[1,3,0],[1,0,0],[1,0,3]]
//Output: 4
//Explanation: One possible sequence of moves to place one stone in each cell
//is:
//1- Move one stone from cell (0,1) to cell (0,2).
//2- Move one stone from cell (0,1) to cell (1,1).
//3- Move one stone from cell (2,2) to cell (1,2).
//4- Move one stone from cell (2,2) to cell (2,1).
//In total, it takes 4 moves to place one stone in each cell of the grid.
//It can be shown that 4 is the minimum number of moves required to place one
//stone in each cell.
//
//
//
// Constraints:
//
//
// grid.length == grid[i].length == 3
// 0 <= grid[i][j] <= 9
// Sum of grid is equal to 9.
//
//
// Related Topics 广度优先搜索 数组 动态规划 矩阵 👍 70 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
package main

import (
	"fmt"
	"math"
)

func minimumMoves(grid [][]int) int {
	more := make([][]int, 0)
	less := make([][]int, 0)

	for i, rows := range grid {
		for j, v := range rows {
			if v == 0 {
				less = append(less, []int{i, j})
			}

			if v > 1 {
				for k := 0; k < v-1; k++ {
					more = append(more, []int{i, j})
				}
			}
		}
	}

	fmt.Println(more)
	fmt.Println(less)

	usedMore := make([]int, len(more))
	usedLess := make([]int, len(less))
	rlt := math.MaxInt
	count := 0
	cur := 0

	var traverse func()
	traverse = func() {
		if count == len(more) {
			rlt = min(cur, rlt)
			return
		}

		for i := 0; i < len(more); i++ {
			//fmt.Println(usedMore)
			//fmt.Println(usedLess)
			//fmt.Println("==============")
			if usedMore[i] == 1 {
				continue
			}

			if i > 0 && slicesEqual(more[i-1], more[i]) && usedMore[i-1] == 0 {
				continue
			}

			usedMore[i] = 1
			for j := 0; j < len(less); j++ {
				if usedLess[j] == 1 {
					continue
				}
				usedLess[j] = 1

				cur += calDistance2850(more[i], less[j])
				count++
				traverse()
				cur -= calDistance2850(more[i], less[j])
				count--

				usedLess[j] = 0
			}
			usedMore[i] = 0
		}
	}
	traverse()
	return rlt
}

func slicesEqual(slice1, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i, v := range slice1 {
		if v != slice2[i] {
			return false
		}
	}
	return true
}

func calDistance2850(i, j []int) int {
	return abs2850(i[0]-j[0]) + abs2850(i[1]-j[1])
}

func abs2850(num int) int {
	if num < 0 {
		return -num
	} else {
		return num
	}
}

//leetcode submit region end(Prohibit modification and deletion)
