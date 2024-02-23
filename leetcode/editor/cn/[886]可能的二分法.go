/**
给定一组 n 人（编号为 1, 2, ..., n）， 我们想把每个人分进任意大小的两组。每个人都可能不喜欢其他人，那么他们不应该属于同一组。

 给定整数 n 和数组 dislikes ，其中 dislikes[i] = [ai, bi] ，表示不允许将编号为 ai 和 bi的人归入同一组。当可以用这种
方法将所有人分进两组时，返回 true；否则返回 false。






 示例 1：


输入：n = 4, dislikes = [[1,2],[1,3],[2,4]]
输出：true
解释：group1 [1,4], group2 [2,3]


 示例 2：


输入：n = 3, dislikes = [[1,2],[1,3],[2,3]]
输出：false


 示例 3：


输入：n = 5, dislikes = [[1,2],[2,3],[3,4],[4,5],[1,5]]
输出：false




 提示：


 1 <= n <= 2000
 0 <= dislikes.length <= 10⁴
 dislikes[i].length == 2
 1 <= dislikes[i][j] <= n
 ai < bi
 dislikes 中每一组都 不同




 Related Topics 深度优先搜索 广度优先搜索 并查集 图 👍 395 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
package main

func possibleBipartition(n int, dislikes [][]int) bool {
	dislikeMap := make(map[int][]int)
	for i := 1; i <= n; i++ {
		dislikeMap[i] = make([]int, 0)
	}

	for _, dislike := range dislikes {
		dislikeMap[dislike[0]] = append(dislikeMap[dislike[0]], dislike[1])
		dislikeMap[dislike[1]] = append(dislikeMap[dislike[1]], dislike[0])
	}

	group1 := make(map[int]struct{})
	group2 := make(map[int]struct{})
	var dfs func(person int, group2 map[int]struct{}, group1 map[int]struct{}) bool
	dfs = func(person int, group2 map[int]struct{}, group1 map[int]struct{}) bool {
		nextIterPersons := make([]int, 0)
		for _, dislikePerson := range dislikeMap[person] {
			if _, ok := group1[dislikePerson]; ok {
				return false
			}

			if _, ok := group2[dislikePerson]; !ok {
				group2[dislikePerson] = struct{}{}
				nextIterPersons = append(nextIterPersons, dislikePerson)
			}
		}

		for _, iterPerson := range nextIterPersons {
			status := dfs(iterPerson, group1, group2)
			if !status {
				return false
			}
		}
		return true
	}

	for i := 1; i <= n; i++ {
		_, ok1 := group1[i]
		_, ok2 := group2[i]
		if !ok1 && !ok2 {
			group1[i] = struct{}{}
			status := dfs(i, group2, group1)
			if !status {
				return false
			}
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
