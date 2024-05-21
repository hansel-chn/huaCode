/**
Given a binary array nums, return the maximum length of a contiguous subarray
with an equal number of 0 and 1.


 Example 1:


Input: nums = [0,1]
Output: 2
Explanation: [0, 1] is the longest contiguous subarray with an equal number of 0
 and 1.


 Example 2:


Input: nums = [0,1,0]
Output: 2
Explanation: [0, 1] (or [1, 0]) is a longest contiguous subarray with equal
number of 0 and 1.



 Constraints:


 1 <= nums.length <= 10⁵
 nums[i] is either 0 or 1.


 Related Topics 数组 哈希表 前缀和 👍 735 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
package main

func findMaxLength(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[i] = -1
		}
	}

	//preSum := make([]int, len(nums)+1)
	sumMap := make(map[int]int)
	//sumMap[nums[0]] = 0
	rlt := 0

	//// solution 1
	//for i := 0; i < len(nums); i++ {
	//	preSum[i+1] = preSum[i] + nums[i]
	//	if v, ok := sumMap[preSum[i+1]]; ok {
	//		rlt = max(rlt, i+1-v)
	//	} else {
	//		sumMap[preSum[i+1]] = i+1
	//	}
	//}

	//// solution 1
	//for i := 1; i < len(nums)+1; i++ {
	//	preSum[i] = preSum[i-1] + nums[i-1]
	//	if v, ok := sumMap[preSum[i]]; ok {
	//		rlt = max(rlt, i-v)
	//	} else {
	//		sumMap[preSum[i]] = i
	//	}
	//
	//}

	sumMap[0] = -1
	preSum := 0
	for i := 0; i < len(nums); i++ {
		preSum = preSum + nums[i]
		if v, ok := sumMap[preSum]; ok {
			rlt = max(rlt, i-v)
		} else {
			sumMap[preSum] = i
		}
	}
	return rlt
}

//leetcode submit region end(Prohibit modification and deletion)
