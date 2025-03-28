//There is a car with capacity empty seats. The vehicle only drives east (i.e.,
//it cannot turn around and drive west).
//
// You are given the integer capacity and an array trips where trips[i] = [
//numPassengersi, fromi, toi] indicates that the iᵗʰ trip has numPassengersi
//passengers and the locations to pick them up and drop them off are fromi and toi
//respectively. The locations are given as the number of kilometers due east from the
//car's initial location.
//
// Return true if it is possible to pick up and drop off all passengers for all
//the given trips, or false otherwise.
//
//
// Example 1:
//
//
//Input: trips = [[2,1,5],[3,3,7]], capacity = 4
//Output: false
//
//
// Example 2:
//
//
//Input: trips = [[2,1,5],[3,3,7]], capacity = 5
//Output: true
//
//
//
// Constraints:
//
//
// 1 <= trips.length <= 1000
// trips[i].length == 3
// 1 <= numPassengersi <= 100
// 0 <= fromi < toi <= 1000
// 1 <= capacity <= 10⁵
//
//
// Related Topics 数组 前缀和 排序 模拟 堆（优先队列） 👍 411 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
package main

func carPooling(trips [][]int, capacity int) bool {
	diffs := make([]int, 1001)
	for _, trip := range trips {
		from := trip[1]
		to := trip[2]
		diffs[from] += trip[0]
		if to < 1000 {
			diffs[to] -= trip[0]
		}
	}

	cur := 0
	for _, diff := range diffs {
		cur += diff
		if cur > capacity {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
