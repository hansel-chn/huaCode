/**
We want to split a group of n people (labeled from 1 to n) into two groups of
any size. Each person may dislike some other people, and they should not go into
the same group.

 Given the integer n and the array dislikes where dislikes[i] = [ai, bi]
indicates that the person labeled ai does not like the person labeled bi, return true
if it is possible to split everyone into two groups in this way.


 Example 1:


Input: n = 4, dislikes = [[1,2],[1,3],[2,4]]
Output: true
Explanation: The first group has [1,4], and the second group has [2,3].


 Example 2:


Input: n = 3, dislikes = [[1,2],[1,3],[2,3]]
Output: false
Explanation: We need at least 3 groups to divide them. We cannot put them in
two groups.



 Constraints:


 1 <= n <= 2000
 0 <= dislikes.length <= 10⁴
 dislikes[i].length == 2
 1 <= ai < bi <= n
 All the pairs of dislikes are unique.


 Related Topics 深度优先搜索 广度优先搜索 并查集 图 👍 404 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
package main

func possibleBipartitionDup(n int, dislikes [][]int) bool {
	gf := buildGraph886(n, dislikes)

	color := make([]bool, n+1)
	visited := make([]bool, n+1)
	ok := new(bool)
	*ok = true

	for i := 1; i < n; i++ {
		if !visited[i] {
			traverse886(i, gf, visited, color, ok)
			if !*ok {
				return *ok
			}
		}
	}
	return *ok
}

func traverse886(person int, graph [][]int, visited []bool, color []bool, ok *bool) {
	if !*ok {
		return
	}

	visited[person] = true
	for _, dislike := range graph[person] {
		if !visited[dislike] {
			color[dislike] = !color[person]
			traverse886(dislike, graph, visited, color, ok)
		} else {
			if color[person] == color[dislike] {
				*ok = false
				return
			}
		}
	}
	return
}

func buildGraph886(n int, dislikes [][]int) [][]int {
	graph := make([][]int, n+1)
	for i := range graph {
		graph[i] = make([]int, 0)
	}

	for _, dislike := range dislikes {
		graph[dislike[0]] = append(graph[dislike[0]], dislike[1])
		// !!!!!!!!!!
		// undirected graph
		// think why need both directions
		// if single directions:
		// case:  1->4, 2->3->4
		// color: 0->1, 0->1->0 (wrong)
		graph[dislike[1]] = append(graph[dislike[1]], dislike[0])
	}
	return graph
}

//leetcode submit region end(Prohibit modification and deletion)
