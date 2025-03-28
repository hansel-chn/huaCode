//You are given an integer array nums and an integer x. In one operation, you
//can either remove the leftmost or the rightmost element from the array nums and
//subtract its value from x. Note that this modifies the array for future
//operations.
//
// Return the minimum number of operations to reduce x to exactly 0 if it is
//possible, otherwise, return -1.
//
//
// Example 1:
//
//
//Input: nums = [1,1,4,2,3], x = 5
//Output: 2
//Explanation: The optimal solution is to remove the last two elements to
//reduce x to zero.
//
//
// Example 2:
//
//
//Input: nums = [5,6,7,8,9], x = 4
//Output: -1
//
//
// Example 3:
//
//
//Input: nums = [3,2,20,1,1,3], x = 10
//Output: 5
//Explanation: The optimal solution is to remove the last three elements and
//the first two elements (5 operations in total) to reduce x to zero.
//
//
//
// Constraints:
//
//
// 1 <= nums.length <= 10⁵
// 1 <= nums[i] <= 10⁴
// 1 <= x <= 10⁹
//
//
// Related Topics 数组 哈希表 二分查找 前缀和 滑动窗口 👍 403 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
package main

func minOperations(nums []int, x int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	res := sum - x
	maxNum := 0
	if res < 0 {
		return -1
	} else if res == 0 {
		return len(nums)
	}

	var l, r, cur int

	for {
		if cur == res {
			maxNum = max(maxNum, r-l)
			if r >= len(nums) {
				break
			}
			cur += nums[r]
			r++
		} else if cur < res {
			if r >= len(nums) {
				break
			}
			cur += nums[r]
			r++
		} else {
			cur -= nums[l]
			l++
		}
	}
	if maxNum == 0 {
		return -1
	} else {
		return len(nums) - maxNum
	}
}

//leetcode submit region end(Prohibit modification and deletion)
