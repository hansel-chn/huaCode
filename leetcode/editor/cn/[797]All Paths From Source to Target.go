/**
Given a directed acyclic graph (DAG) of n nodes labeled from 0 to n - 1, find
all possible paths from node 0 to node n - 1 and return them in any order.

 The graph is given as follows: graph[i] is a list of all nodes you can visit
from node i (i.e., there is a directed edge from node i to node graph[i][j]).


 Example 1:


Input: graph = [[1,2],[3],[3],[]]
Output: [[0,1,3],[0,2,3]]
Explanation: There are two paths: 0 -> 1 -> 3 and 0 -> 2 -> 3.


 Example 2:


Input: graph = [[4,3,1],[3,2,4],[3],[4],[]]
Output: [[0,4],[0,3,4],[0,1,3,4],[0,1,2,3,4],[0,1,4]]



 Constraints:


 n == graph.length
 2 <= n <= 15
 0 <= graph[i][j] < n
 graph[i][j] != i (i.e., there will be no self-loops).
 All the elements of graph[i] are unique.
 The input graph is guaranteed to be a DAG.


 Related Topics 深度优先搜索 广度优先搜索 图 回溯 👍 472 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
package main

func allPathsSourceTarget(graph [][]int) [][]int {
	rlt := make([][]int, 0)
	current := make([]int, 0)
	end := len(graph) - 1
	// 遍历由于加上visited数组获取结果错误，比如末尾节点，第一次遍历后标记为true，之后每次遍历都忽略了
	// 注意为什么会产生这个问题，遍历和防止循环目的在哪里
	//
	// 如何在有环的情况下遍历所有路径
	//
	// visited作用感觉是在从一幅图的多个节点遍历时，且每个节点只需要遍历一次时，需要的东西，
	// 她并不期望重复遍历之前已经遍历过的路径，像本题，其实与visited数组的设置相违背
	// 本题从单一结点出发，其实判断环onPath足够了，目的是重复遍历而不是图上所有点遍历一次
	// https://labuladong.online/algo/data-structure/graph-traverse/#%E9%A2%98%E7%9B%AE%E5%AE%9E%E8%B7%B5说:
	// 如果无visited数组有可能会进入无限递归，题目见的少存疑，目前的感觉是会浪费，重复遍历之前的路径，但是有onPath感觉并不会进入死循环

	//visited := make([]bool, len(graph))
	onPath := make([]bool, len(graph))

	var traverse func(node int)
	traverse = func(node int) {
		//if visited[node]  {
		//	return
		//}
		//
		//visited[node] = true
		if onPath[node] {
			return
		}

		current = append(current, node)
		onPath[node] = true

		if node == end {
			temp := make([]int, len(current))
			copy(temp, current)
			rlt = append(rlt, temp)
			current = current[:len(current)-1]
			onPath[node] = false
			return
		}

		for _, nextJump := range graph[node] {
			traverse(nextJump)
		}
		current = current[:len(current)-1]
		onPath[node] = false
	}

	traverse(0)
	return rlt
}

//leetcode submit region end(Prohibit modification and deletion)
