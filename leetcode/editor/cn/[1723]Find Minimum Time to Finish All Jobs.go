//You are given an integer array jobs, where jobs[i] is the amount of time it
//takes to complete the iᵗʰ job.
//
// There are k workers that you can assign jobs to. Each job should be assigned
//to exactly one worker. The working time of a worker is the sum of the time it
//takes to complete all jobs assigned to them. Your goal is to devise an optimal
//assignment such that the maximum working time of any worker is minimized.
//
// Return the minimum possible maximum working time of any assignment.
//
//
// Example 1:
//
//
//Input: jobs = [3,2,3], k = 3
//Output: 3
//Explanation: By assigning each person one job, the maximum time is 3.
//
//
// Example 2:
//
//
//Input: jobs = [1,2,4,7,8], k = 2
//Output: 11
//Explanation: Assign the jobs the following way:
//Worker 1: 1, 2, 8 (working time = 1 + 2 + 8 = 11)
//Worker 2: 4, 7 (working time = 4 + 7 = 11)
//The maximum working time is 11.
//
//
// Constraints:
//
//
// 1 <= k <= jobs.length <= 12
// 1 <= jobs[i] <= 10⁷
//
//
// Related Topics 位运算 数组 动态规划 回溯 状态压缩 👍 337 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
package main

import "math"

func minimumTimeRequired(jobs []int, k int) int {
	assign := make([]int, k)
	minTime := math.MaxInt64

	var traverse func(nxt int)
	traverse = func(nxt int) {
		if nxt == len(jobs) {
			minTime = min(calMaxTime(assign), minTime)
			return
		}

		workerTimeMap := make(map[int]struct{})
		for i := range k {
			if minTime <= assign[i]+jobs[nxt] {
				continue
			}

			if _, ok := workerTimeMap[assign[i]]; ok {
				continue
			}

			workerTimeMap[assign[i]] = struct{}{}

			assign[i] += jobs[nxt]
			traverse(nxt + 1)
			assign[i] -= jobs[nxt]
		}
	}
	traverse(0)
	return minTime
}

func calMaxTime(assign []int) int {
	maxTime := 0
	for _, time := range assign {
		maxTime = max(maxTime, time)
	}
	return maxTime
}

//leetcode submit region end(Prohibit modification and deletion)
