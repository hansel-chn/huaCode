/**
Given an array of intervals intervals where intervals[i] = [starti, endi],
return the minimum number of intervals you need to remove to make the rest of the
intervals non-overlapping.


 Example 1:


Input: intervals = [[1,2],[2,3],[3,4],[1,3]]
Output: 1
Explanation: [1,3] can be removed and the rest of the intervals are non-
overlapping.


 Example 2:


Input: intervals = [[1,2],[1,2],[1,2]]
Output: 2
Explanation: You need to remove two [1,2] to make the rest of the intervals non-
overlapping.


 Example 3:


Input: intervals = [[1,2],[2,3]]
Output: 0
Explanation: You don't need to remove any of the intervals since they're
already non-overlapping.



 Constraints:


 1 <= intervals.length <= 10⁵
 intervals[i].length == 2
 -5 * 10⁴ <= starti < endi <= 5 * 10⁴


 Related Topics 贪心 数组 动态规划 排序 👍 1133 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
package main

import (
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
	itv := sortIntervals(intervals)
	sort.Sort(&itv)

	last := int(-1e5)
	erase := 0

	for _, interval := range itv {
		if interval[0] < last {
			erase++
		} else {
			last = interval[1]
		}
	}
	return erase
}

type sortIntervals [][]int

func (in sortIntervals) Less(i, j int) bool {
	return in[i][1] <= in[j][1]
}

func (in *sortIntervals) Swap(i, j int) {
	(*in)[i], (*in)[j] = (*in)[j], (*in)[i]
}

func (in sortIntervals) Len() int {
	return len(in)
}

//leetcode submit region end(Prohibit modification and deletion)
