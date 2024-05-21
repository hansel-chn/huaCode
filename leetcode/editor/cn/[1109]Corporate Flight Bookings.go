/**
There are n flights that are labeled from 1 to n.

 You are given an array of flight bookings bookings, where bookings[i] = [
firsti, lasti, seatsi] represents a booking for flights firsti through lasti (
inclusive) with seatsi seats reserved for each flight in the range.

 Return an array answer of length n, where answer[i] is the total number of
seats reserved for flight i.


 Example 1:


Input: bookings = [[1,2,10],[2,3,20],[2,5,25]], n = 5
Output: [10,55,45,25,25]
Explanation:
Flight labels:        1   2   3   4   5
Booking 1 reserved:  10  10
Booking 2 reserved:      20  20
Booking 3 reserved:      25  25  25  25
Total seats:         10  55  45  25  25
Hence, answer = [10,55,45,25,25]


 Example 2:


Input: bookings = [[1,2,10],[2,2,15]], n = 2
Output: [10,25]
Explanation:
Flight labels:        1   2
Booking 1 reserved:  10  10
Booking 2 reserved:      15
Total seats:         10  25
Hence, answer = [10,25]




 Constraints:


 1 <= n <= 2 * 10⁴
 1 <= bookings.length <= 2 * 10⁴
 bookings[i].length == 3
 1 <= firsti <= lasti <= n
 1 <= seatsi <= 10⁴


 Related Topics 数组 前缀和 👍 515 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
package main

func corpFlightBookings(bookings [][]int, n int) []int {
	diff := make([]int, n)
	for _, booking := range bookings {
		diff[booking[0]-1] += booking[2]
		if booking[1] < n {
			diff[booking[1]] -= booking[2]
		}
	}

	rlt := make([]int, n)
	rlt[0] = diff[0]
	for i := 1; i < len(diff); i++ {
		rlt[i] = rlt[i-1] + diff[i]
	}
	return rlt
}

//func corpFlightBookings(bookings [][]int, n int) []int {
//	diff := make([]int, n)
//	for _, booking := range bookings {
//		diff[booking[1]-1] += booking[2]
//		if booking[0] > 1 {
//			diff[booking[0]-2] -= booking[2]
//		}
//	}
//
//	rlt := make([]int, n)
//	for i := n - 1; i >= 0; i-- {
//		if i == n-1 {
//			rlt[i] = diff[i]
//		} else {
//			rlt[i] = rlt[i+1] + diff[i]
//		}
//	}
//	return rlt
//}

//leetcode submit region end(Prohibit modification and deletion)
