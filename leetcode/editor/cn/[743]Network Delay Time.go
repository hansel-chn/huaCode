/**
You are given a network of n nodes, labeled from 1 to n. You are also given
times, a list of travel times as directed edges times[i] = (ui, vi, wi), where ui
is the source node, vi is the target node, and wi is the time it takes for a
signal to travel from source to target.

 We will send a signal from a given node k. Return the minimum time it takes
for all the n nodes to receive the signal. If it is impossible for all the n nodes
to receive the signal, return -1.


 Example 1:


Input: times = [[2,1,1],[2,3,1],[3,4,1]], n = 4, k = 2
Output: 2


 Example 2:


Input: times = [[1,2,1]], n = 2, k = 1
Output: 1


 Example 3:


Input: times = [[1,2,1]], n = 2, k = 2
Output: -1



 Constraints:


 1 <= k <= n <= 100
 1 <= times.length <= 6000
 times[i].length == 3
 1 <= ui, vi <= n
 ui != vi
 0 <= wi <= 100
 All the pairs (ui, vi) are unique. (i.e., no multiple edges.)


 Related Topics 深度优先搜索 广度优先搜索 图 最短路 堆（优先队列） 👍 741 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
package main

import (
	"math"
)

func networkDelayTime(times [][]int, n int, k int) int {
	graph := make([][][]int, n+1)
	for _, time := range times {
		graph[time[0]] = append(graph[time[0]], []int{time[1], time[2]})
	}

	costTime := minCostTime(graph, k)

	maxTimeAmongAll := -1
	for i, cost := range costTime {
		if i == 0 {
			continue
		}

		if cost == math.MaxInt {
			return -1
		}
		maxTimeAmongAll = max(maxTimeAmongAll, cost)
	}
	return maxTimeAmongAll
}

func minCostTime(graph [][][]int, start int) []int {

	minTime := make([]int, len(graph))
	for i := range minTime {
		minTime[i] = math.MaxInt
	}
	minTime[start] = 0

	queue := make(priorityQueue743, 0)
	queue.Push(node743{
		id:   start,
		time: 0,
	})

	for queue.Len() > 0 {
		pop := queue.Pop().(node743)
		if minTime[pop.id] < pop.time {
			continue
		}

		for _, des := range graph[pop.id] {
			if pop.time+des[1] < minTime[des[0]] {
				queue.Push(node743{
					id:   des[0],
					time: pop.time + des[1],
				})
				minTime[des[0]] = pop.time + des[1]
			}
		}
	}
	return minTime
}

type node743 struct {
	id   int
	time int
}

type priorityQueue743 []node743

func (p priorityQueue743) Less(i, j int) bool {
	return p[i].time <= p[j].time
}

func (p priorityQueue743) Len() int {
	return len(p)
}

func (p *priorityQueue743) Swap(i, j int) {
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

func (p *priorityQueue743) Push(x any) {
	*p = append(*p, x.(node743))
}

func (p *priorityQueue743) Pop() any {
	length := p.Len()
	old := (*p)[length-1]
	*p = (*p)[:length-1]
	return old
}

//leetcode submit region end(Prohibit modification and deletion)
