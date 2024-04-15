/**
Given a circular integer array nums (i.e., the next element of nums[nums.length
- 1] is nums[0]), return the next greater number for every element in nums.

 The next greater number of a number x is the first greater number to its
traversing-order next in the array, which means you could search circularly to find
its next greater number. If it doesn't exist, return -1 for this number.


 Example 1:


Input: nums = [1,2,1]
Output: [2,-1,2]
Explanation: The first 1's next greater number is 2;
The number 2 can't find next greater number.
The second 1's next greater number needs to search circularly, which is also 2.


 Example 2:


Input: nums = [1,2,3,4,3]
Output: [2,3,4,-1,4]



 Constraints:


 1 <= nums.length <= 10⁴
 -10⁹ <= nums[i] <= 10⁹


 Related Topics 栈 数组 单调栈 👍 926 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
package main

func nextGreaterElements(nums []int) []int {
	rlt := make([]int, len(nums))
	monoStack := make([]int, 0)

	for i := len(nums)*2 - 1; i >= 0; i-- {
		for len(monoStack) > 0 && nums[i%len(nums)] >= monoStack[len(monoStack)-1] {
			monoStack = monoStack[:len(monoStack)-1]
		}

		if i < len(nums) {
			if len(monoStack) > 0 {
				rlt[i] = monoStack[len(monoStack)-1]
			} else {
				rlt[i] = -1
			}
		}
		monoStack = append(monoStack, nums[i%len(nums)])
	}
	return rlt
}

//leetcode submit region end(Prohibit modification and deletion)
