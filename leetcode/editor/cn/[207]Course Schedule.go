/**
There are a total of numCourses courses you have to take, labeled from 0 to
numCourses - 1. You are given an array prerequisites where prerequisites[i] = [ai,
bi] indicates that you must take course bi first if you want to take course ai.


 For example, the pair [0, 1], indicates that to take course 0 you have to
first take course 1.


 Return true if you can finish all courses. Otherwise, return false.


 Example 1:


Input: numCourses = 2, prerequisites = [[1,0]]
Output: true
Explanation: There are a total of 2 courses to take.
To take course 1 you should have finished course 0. So it is possible.


 Example 2:


Input: numCourses = 2, prerequisites = [[1,0],[0,1]]
Output: false
Explanation: There are a total of 2 courses to take.
To take course 1 you should have finished course 0, and to take course 0 you
should also have finished course 1. So it is impossible.



 Constraints:


 1 <= numCourses <= 2000
 0 <= prerequisites.length <= 5000
 prerequisites[i].length == 2
 0 <= ai, bi < numCourses
 All the pairs prerequisites[i] are unique.


 Related Topics 深度优先搜索 广度优先搜索 图 拓扑排序 👍 1924 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	for _, v := range prerequisites {
		graph[v[1]] = append(graph[v[1]], v[0])
	}

	visited := make([]bool, numCourses)
	onPath := make([]bool, numCourses)
	hasCycle := false

	var traverse func(node int)
	traverse = func(node int) {
		if hasCycle {
			return
		}

		if onPath[node] {
			hasCycle = true
			return
		}

		if visited[node] {
			return
		}

		onPath[node] = true
		visited[node] = true
		for _, i := range graph[node] {
			traverse(i)
		}

		onPath[node] = false
	}

	for i := 0; i < numCourses; i++ {
		traverse(i)
	}
	return !hasCycle
}

//leetcode submit region end(Prohibit modification and deletion)
