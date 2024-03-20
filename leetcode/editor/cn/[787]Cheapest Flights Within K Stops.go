/**
There are n cities connected by some number of flights. You are given an array
flights where flights[i] = [fromi, toi, pricei] indicates that there is a flight
from city fromi to city toi with cost pricei.

 You are also given three integers src, dst, and k, return the cheapest price
from src to dst with at most k stops. If there is no such route, return -1.


 Example 1:


Input: n = 4, flights = [[0,1,100],[1,2,100],[2,0,100],[1,3,600],[2,3,200]],
src = 0, dst = 3, k = 1
Output: 700
Explanation:
The graph is shown above.
The optimal path with at most 1 stop from city 0 to 3 is marked in red and has
cost 100 + 600 = 700.
Note that the path through cities [0,1,2,3] is cheaper but is invalid because
it uses 2 stops.


 Example 2:


Input: n = 3, flights = [[0,1,100],[1,2,100],[0,2,500]], src = 0, dst = 2, k = 1

Output: 200
Explanation:
The graph is shown above.
The optimal path with at most 1 stop from city 0 to 2 is marked in red and has
cost 100 + 100 = 200.


 Example 3:


Input: n = 3, flights = [[0,1,100],[1,2,100],[0,2,500]], src = 0, dst = 2, k = 0

Output: 500
Explanation:
The graph is shown above.
The optimal path with no stops from city 0 to 2 is marked in red and has cost 50
0.



 Constraints:


 1 <= n <= 100
 0 <= flights.length <= (n * (n - 1) / 2)
 flights[i].length == 3
 0 <= fromi, toi < n
 fromi != toi
 1 <= pricei <= 10⁴
 There will not be any multiple flights between two cities.
 0 <= src, dst, k < n
 src != dst


 Related Topics 深度优先搜索 广度优先搜索 图 动态规划 最短路 堆（优先队列） 👍 636 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
package main

//func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
//	dstMap := make(map[int][][2]int)
//	for _, flight := range flights {
//		if _, ok := dstMap[flight[1]]; !ok {
//			dstMap[flight[1]] = [][2]int{{flight[0], flight[2]}}
//		} else {
//			dstMap[flight[1]] = append(dstMap[flight[1]], [2]int{flight[0], flight[2]})
//		}
//	}
//
//	dp := make([][][2]int, n)
//	for i := range dp {
//		dp[i] = make([][2]int, k+1)
//	}
//
//	for j := k; j >= 0; j-- {
//		for i := n - 1; i >= 0; i-- {
//			if j == k {
//				for _, v := range dstMap[dst] {
//					dp[i][j] = [2]int{v[1], 1}
//				}
//				continue
//			}
//
//
//			if dp[i+1][j+1]
//
//		}
//	}
//
//}

//func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
//	srcPath := make(map[int][2]int, 0)
//	dstMap := make(map[int][][2]int)
//	for _, flight := range flights {
//		if _, ok := dstMap[flight[1]]; !ok {
//			dstMap[flight[1]] = [][2]int{{flight[0], flight[2]}}
//		} else {
//			dstMap[flight[1]] = append(dstMap[flight[1]], [2]int{flight[0], flight[2]})
//		}
//
//		if flight[0] == src {
//			if _, ok := srcPath[flight[1]]; !ok {
//				srcPath[flight[1]] = [2]int{flight[1], flight[2]}
//			}
//		}
//	}
//
//	dp := make([][][2]int, n)
//	for i := range dp {
//		dp[i] = make([][2]int, k+1)
//	}
//
//	for j := 0; j <= k; j++ {
//		for i := 0; i < n; i++ {
//			if j == 0 {
//				if v, ok := srcPath[i]; ok {
//					dp[i][j] = [2]int{v[1], 1}
//				}
//				continue
//			}
//
//			minPrice := math.MaxInt32
//			exist := false
//			for _, v := range dstMap[i] {
//				if dp[v[0]][j-1][1] == 1 {
//					exist = true
//					minPrice = min(minPrice, dp[v[0]][j-1][0]+v[1])
//				}
//			}
//			if exist {
//				dp[i][j] = [2]int{minPrice, 1}
//			}
//		}
//	}
//
//	rlt := math.MaxInt32
//	exist := false
//	for i := 0; i <= k; i++ {
//		if dp[dst][i][1] == 1 {
//			exist = true
//			rlt = min(rlt, dp[dst][i][0])
//		}
//	}
//
//	fmt.Println(dp)
//	if exist {
//		return rlt
//	} else {
//		return -1
//	}
//}

//leetcode submit region end(Prohibit modification and deletion)
