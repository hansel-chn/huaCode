/**
You are a hiker preparing for an upcoming hike. You are given heights, a 2D
array of size rows x columns, where heights[row][col] represents the height of cell
(row, col). You are situated in the top-left cell, (0, 0), and you hope to
travel to the bottom-right cell, (rows-1, columns-1) (i.e., 0-indexed). You can move
up, down, left, or right, and you wish to find a route that requires the
minimum effort.

 A route's effort is the maximum absolute difference in heights between two
consecutive cells of the route.

 Return the minimum effort required to travel from the top-left cell to the
bottom-right cell.


 Example 1:




Input: heights = [[1,2,2],[3,8,2],[5,3,5]]
Output: 2
Explanation: The route of [1,3,5,3,5] has a maximum absolute difference of 2 in
consecutive cells.
This is better than the route of [1,2,2,2,5], where the maximum absolute
difference is 3.


 Example 2:




Input: heights = [[1,2,3],[3,8,4],[5,3,5]]
Output: 1
Explanation: The route of [1,2,3,4,5] has a maximum absolute difference of 1 in
consecutive cells, which is better than route [1,3,5,3,5].


 Example 3:


Input: heights = [[1,2,1,1,1],[1,2,1,2,1],[1,2,1,2,1],[1,2,1,2,1],[1,1,1,2,1]]
Output: 0
Explanation: This route does not require any effort.



 Constraints:


 rows == heights.length
 columns == heights[i].length
 1 <= rows, columns <= 100
 1 <= heights[i][j] <= 10⁶


 Related Topics 深度优先搜索 广度优先搜索 并查集 数组 二分查找 矩阵 堆（优先队列） 👍 511 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
package main

import (
	"container/heap"
	"math"
)

func minimumEffortPath(heights [][]int) int {
	pq := make(priorityQueue1631, 0)
	heap.Push(&pq, node1631{
		x:     0,
		y:     0,
		value: 0,
	})

	minEffort := make([][]int, len(heights))
	for i := range minEffort {
		minEffort[i] = make([]int, len(heights[0]))
		for j := range minEffort[i] {
			minEffort[i][j] = math.MaxInt
		}
	}

	minEffort[0][0] = 0

	for pq.Len() > 0 {
		pop := heap.Pop(&pq).(node1631)

		if pop.x == len(heights)-1 && pop.y == len(heights[0])-1 {
			return pop.value
		}

		if pop.value > minEffort[pop.x][pop.y] {
			continue
		}

		for _, adjNode := range adj1631(pop.x, pop.y, heights) {
			adjX := adjNode[0]
			adjY := adjNode[1]
			adjEffort := max(pop.value, int(math.Abs(float64(heights[adjX][adjY]-heights[pop.x][pop.y]))))
			if adjEffort < minEffort[adjX][adjY] {

				heap.Push(&pq, node1631{
					x:     adjX,
					y:     adjY,
					value: adjEffort,
				})
				minEffort[adjX][adjY] = adjEffort
			}
		}
	}
	return minEffort[len(heights)-1][len(heights[0])-1]
}

type node1631 struct {
	x, y  int
	value int
}

type priorityQueue1631 []node1631

func (p priorityQueue1631) Less(i, j int) bool {
	return p[i].value <= p[j].value
}

func (p *priorityQueue1631) Swap(i, j int) {
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

func (p priorityQueue1631) Len() int {
	return len(p)
}

func (p *priorityQueue1631) Push(x any) {
	*p = append(*p, x.(node1631))
}

func (p *priorityQueue1631) Pop() any {
	length := p.Len()
	pop := (*p)[length-1]
	*p = (*p)[:length-1]
	return pop
}

var dirs = [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

func adj1631(i, j int, heights [][]int) [][2]int {
	adj := make([][2]int, 0)
	rLimit := len(heights)
	cLimit := len(heights[0])
	for _, dir := range dirs {
		ni := i + dir[0]
		nj := j + dir[1]
		if ni >= 0 && ni < rLimit && nj >= 0 && nj < cLimit {
			adj = append(adj, [2]int{ni, nj})
		}
	}
	return adj
}

//leetcode submit region end(Prohibit modification and deletion)
