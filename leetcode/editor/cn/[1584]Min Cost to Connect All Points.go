/**
You are given an array points representing integer coordinates of some points
on a 2D-plane, where points[i] = [xi, yi].

 The cost of connecting two points [xi, yi] and [xj, yj] is the manhattan
distance between them: |xi - xj| + |yi - yj|, where |val| denotes the absolute value
of val.

 Return the minimum cost to make all points connected. All points are connected
if there is exactly one simple path between any two points.


 Example 1:


Input: points = [[0,0],[2,2],[3,10],[5,2],[7,0]]
Output: 20
Explanation:

We can connect the points as shown above to get the minimum cost of 20.
Notice that there is a unique path between every pair of points.


 Example 2:


Input: points = [[3,12],[-2,5],[-4,1]]
Output: 18



 Constraints:


 1 <= points.length <= 1000
 -10⁶ <= xi, yi <= 10⁶
 All pairs (xi, yi) are distinct.


 Related Topics 并查集 图 数组 最小生成树 👍 318 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
package main

import (
	"math"
	"sort"
)

func minCostConnectPoints(points [][]int) int {
	var cost int
	uf := newUF1584(len(points))

	weights := make([][3]int, 0)
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			weights = append(weights, [3]int{i, j, distance1584(points, i, j)})
		}
	}

	sort.Slice(weights, func(i, j int) bool {
		return weights[i][2] <= weights[j][2]
	})

	for _, weight := range weights {
		x := weight[0]
		y := weight[1]

		if uf.Count() == 1 {
			return cost
		}

		if uf.IfConnect(x, y) {
			continue
		} else {
			uf.Union(x, y)
			cost += weight[2]
		}

	}

	if uf.Count() == 1 {
		return cost
	} else {
		return -1
	}
}

func distance1584(points [][]int, x, y int) int {
	return int(math.Abs(float64(points[x][0]-points[y][0])) + math.Abs(float64(points[x][1]-points[y][1])))
}

func newUF1584(n int) UF1584 {
	p := make([]int, n)

	for i := range p {
		p[i] = i
	}

	return UF1584{
		count:  n,
		parent: p,
	}
}

type UF1584 struct {
	count  int
	parent []int
}

func (u UF1584) Count() int {
	return u.count
}

func (u *UF1584) findTopP(x int) int {
	if x == u.parent[x] {
		return x
	}

	topP := u.findTopP(u.parent[x])
	u.parent[x] = topP
	return topP
}

func (u *UF1584) IfConnect(x, y int) bool {
	if u.findTopP(x) == u.findTopP(y) {
		return true
	} else {
		return false
	}
}

func (u *UF1584) Union(x, y int) {
	xTopP := u.findTopP(x)
	yTopP := u.findTopP(y)
	if xTopP == yTopP {
		return
	}

	u.parent[xTopP] = yTopP
	u.count--
	return
}

//leetcode submit region end(Prohibit modification and deletion)
