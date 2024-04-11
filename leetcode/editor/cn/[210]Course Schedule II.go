/**
There are a total of numCourses courses you have to take, labeled from 0 to
numCourses - 1. You are given an array prerequisites where prerequisites[i] = [ai,
bi] indicates that you must take course bi first if you want to take course ai.


 For example, the pair [0, 1], indicates that to take course 0 you have to
first take course 1.


 Return the ordering of courses you should take to finish all courses. If there
are many valid answers, return any of them. If it is impossible to finish all
courses, return an empty array.


 Example 1:


Input: numCourses = 2, prerequisites = [[1,0]]
Output: [0,1]
Explanation: There are a total of 2 courses to take. To take course 1 you
should have finished course 0. So the correct course order is [0,1].


 Example 2:


Input: numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
Output: [0,2,1,3]
Explanation: There are a total of 4 courses to take. To take course 3 you
should have finished both courses 1 and 2. Both courses 1 and 2 should be taken after
you finished course 0.
So one correct course order is [0,1,2,3]. Another correct ordering is [0,2,1,3].



 Example 3:


Input: numCourses = 1, prerequisites = []
Output: [0]



 Constraints:


 1 <= numCourses <= 2000
 0 <= prerequisites.length <= numCourses * (numCourses - 1)
 prerequisites[i].length == 2
 0 <= ai, bi < numCourses
 ai != bi
 All the pairs [ai, bi] are distinct.


 Related Topics 深度优先搜索 广度优先搜索 图 拓扑排序 👍 935 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
package main

// 考虑四个节点A，B，C，D
// 的确不能前序遍历，之前考虑是在最后对不同起点出发得到的数组进行反转并拼接可以解决某一个节点D依赖于B，C的情况，
// 这种情况（四个节点关系为B，C不依赖于任何节点，D依赖于B，C节点)下是成立的；
// 但是还存在一种情况（四个节点关系为B，C依赖于节点A，D依赖于B，C节点），这种情况下仅仅在最后进行反转并且拼接是不对的，会出现A->B->D->C错误的结果。
// BFS可以天然解决问题
// 问题：有没有可能在每次traverse时得到正常顺序的结果
// 好像还得依赖于后续遍历，或者构建图采用相反的模式
// 思考的时候着重考虑D点，为什么前序遍历和后序遍历有差别，核心描述在于不管什么遍历都是从“左”子树开始，所以造成单纯地取反并不能得到相同的结果，尤其是针对
// 在常规树中不可能出现的，某个节点具有两个不同父节点的情况
//
// 另外再考虑另一种情况多叉树或图，循环内处理和循环外处理的区别(暂时感觉在前处理有点类似，回来试试就是前序遍历和循环内前，好像也不对，append对？因为没有改变顺序？)
// 好像就是前序遍历呀，明天试一试
//
//	func findOrder(numCourses int, prerequisites [][]int) []int {
//		graph := make([][]int, numCourses)
//		for _, v := range prerequisites {
//			graph[v[1]] = append(graph[v[1]], v[0])
//		}
//
//		rlt := make([][]int, 0)
//		visited := make([]bool, numCourses)
//		onPath := make([]bool, numCourses)
//		hasCycle := false
//
//		var traverse func(node int) (seq []int)
//		traverse = func(node int) (seq []int) {
//			if hasCycle {
//				return
//			}
//
//			if onPath[node] {
//				hasCycle = true
//				return
//			}
//
//			if visited[node] {
//				return
//			}
//
//			onPath[node] = true
//			visited[node] = true
//			seq = append(seq, node)
//			for _, i := range graph[node] {
//				seq = append(seq, traverse(i)...)
//			}
//			onPath[node] = false
//			return
//		}
//
//		for i := 0; i < numCourses; i++ {
//			rlt = append(rlt, traverse(i))
//		}
//
//		reverse := make([]int, 0)
//		for i := len(rlt) - 1; i >= 0; i-- {
//			reverse = append(reverse, rlt[i]...)
//		}
//		return reverse
//	}

//func findOrder(numCourses int, prerequisites [][]int) []int {
//	graph := make([][]int, numCourses)
//	for _, v := range prerequisites {
//		graph[v[1]] = append(graph[v[1]], v[0])
//	}
//
//	fmt.Println(graph)
//	rlt := make([]int, 0)
//	visited := make([]bool, numCourses)
//	onPath := make([]bool, numCourses)
//	hasCycle := false
//
//	var traverse func(node int)
//	traverse = func(node int) {
//		if hasCycle {
//			return
//		}
//
//		if onPath[node] {
//			hasCycle = true
//			return
//		}
//
//		if visited[node] {
//			return
//		}
//
//		onPath[node] = true
//		visited[node] = true
//		for _, i := range graph[node] {
//			traverse(i)
//		}
//		rlt = append(rlt, node)
//		//fmt.Println(node)
//		onPath[node] = false
//		return
//	}
//
//	for i := 0; i < numCourses; i++ {
//		traverse(i)
//	}
//
//	if hasCycle {
//		return []int{}
//	}
//
//	for i, j := 0, len(rlt)-1; i <= j; i, j = i+1, j-1 {
//		rlt[i], rlt[j] = rlt[j], rlt[i]
//	}
//	//slices.SortFunc(rlt, func(a, b int) int { return (-1) * cmp.Compare(a, b) })
//
//	return rlt
//}

// BFS要注意的地方
// 思路其实不是从某一个起点开始的。思路是从图整体出发，每次遍历后，取图上新入度为零的点。这样就可以保证不会出现上面描述的情况（某个节点依赖于两个节点导致最后排序错误）
func findOrder(numCourses int, prerequisites [][]int) []int {
	rlt := make([]int, 0)
	graph := make([][]int, numCourses)
	inDegrees := make([]int, numCourses)
	for _, v := range prerequisites {
		graph[v[1]] = append(graph[v[1]], v[0])
		inDegrees[v[0]]++
	}

	queue := make([]int, 0)
	for i, inDegree := range inDegrees {
		if inDegree == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) > 0 {
		node := queue[0]
		rlt = append(rlt, node)
		queue = queue[1:]
		for _, to := range graph[node] {
			inDegrees[to]--
			if inDegrees[to] == 0 {
				queue = append(queue, to)
			}
		}
	}
	if len(rlt) != numCourses {
		return []int{}
	}
	return rlt
}

//leetcode submit region end(Prohibit modification and deletion)
