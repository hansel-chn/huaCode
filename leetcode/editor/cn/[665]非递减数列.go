/**
给你一个长度为 n 的整数数组
 nums ，请你判断在 最多 改变 1 个元素的情况下，该数组能否变成一个非递减数列。

 我们是这样定义一个非递减数列的： 对于数组中任意的 i (0 <= i <= n-2)，总满足 nums[i] <= nums[i + 1]。



 示例 1:


输入: nums = [4,2,3]
输出: true
解释: 你可以通过把第一个 4 变成 1 来使得它成为一个非递减数列。


 示例 2:


输入: nums = [4,2,1]
输出: false
解释: 你不能在只改变一个元素的情况下将其变为非递减数列。




 提示：



 n == nums.length
 1 <= n <= 10⁴
 -10⁵ <= nums[i] <= 10⁵


 Related Topics 数组 👍 768 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
package main

func checkPossibility(nums []int) bool {
	//left := int(-2e9)
	//right := int(2e9)
	checkV := int(-2e9)
	count := 0

	for i, num := range nums {
		if checkV <= num {
			checkV = num
		} else {
			if count > 0 {
				return false
			} else {
				var left int
				var right int
				if i-2 < 0 {
					left = -2e9
				} else {
					left = nums[i-2]
				}

				if i+1 > len(nums)-1 {
					right = 2e9
				} else {
					right = nums[i+1]
				}

				if (left <= checkV && checkV <= right) || (left <= num && num <= right) {
					count++
					checkV = num
				} else {
					return false
				}
			}
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
