/**
We have n jobs, where every job is scheduled to be done from startTime[i] to
endTime[i], obtaining a profit of profit[i].

 You're given the startTime, endTime and profit arrays, return the maximum
profit you can take such that there are no two jobs in the subset with overlapping
time range.

 If you choose a job that ends at time X you will be able to start another job
that starts at time X.


 Example 1:




Input: startTime = [1,2,3,3], endTime = [3,4,5,6], profit = [50,10,40,70]
Output: 120
Explanation: The subset chosen is the first and fourth job.
Time range [1-3]+[3-6] , we get profit of 120 = 50 + 70.


 Example 2:




Input: startTime = [1,2,3,4,6], endTime = [3,5,10,6,9], profit = [20,20,100,70,6
0]
Output: 150
Explanation: The subset chosen is the first, fourth and fifth job.
Profit obtained 150 = 20 + 70 + 60.


 Example 3:




Input: startTime = [1,1,1], endTime = [2,3,4], profit = [5,6,4]
Output: 6



 Constraints:


 1 <= startTime.length == endTime.length == profit.length <= 5 * 10⁴
 1 <= startTime[i] < endTime[i] <= 10⁹
 1 <= profit[i] <= 10⁴


 Related Topics 数组 二分查找 动态规划 排序 👍 457 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
package main

import (
	"sort"
)

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	jobs := make([][3]int, 0)
	for i := 0; i < len(startTime); i++ {
		jobs = append(jobs, [3]int{startTime[i], endTime[i], profit[i]})
	}

	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i][1] <= jobs[j][1]
	})

	dp := make([]int, len(jobs)+1)
	for i := range dp {
		if i == 0 {
			dp[i] = 0
			continue
		}

		idx := sort.Search(len(jobs), func(k int) bool {
			return jobs[k][1] > jobs[i-1][0]
		})
		dp[i] = max(dp[i-1], dp[idx]+jobs[i-1][2])
	}
	return dp[len(jobs)]
}

//leetcode submit region end(Prohibit modification and deletion)
