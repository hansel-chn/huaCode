/**
给你一个长度为 n 的整数数组，每次操作将会使 n - 1 个元素增加 1 。返回让数组所有元素相等的最小操作次数。



 示例 1：


输入：nums = [1,2,3]
输出：3
解释：
只需要3次操作（注意每次操作会增加两个元素的值）：
[1,2,3]  =>  [2,3,3]  =>  [3,4,3]  =>  [4,4,4]


 示例 2：


输入：nums = [1,1,1]
输出：0




 提示：


 n == nums.length
 1 <= nums.length <= 10⁵
 -10⁹ <= nums[i] <= 10⁹
 答案保证符合 32-bit 整数


 Related Topics 数组 数学 👍 544 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
package main

//另一种思路，较小数需要几次追平最大数；
//随着最大数个数的增多，较小数需要额外几次追平较大数；
//不过性能没啥优势，但是想想挺有意思的
//func minMoves(nums []int) int {
//	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
//	residual := make([]int, 0)
//	for i := 0; i < len(nums)-1; i++ {
//		residual = append(residual, nums[i]-nums[i+1])
//	}
//	count := 0
//	for i := 0; i < len(residual); i++ {
//		count += residual[i] * (i + 1)
//	}
//	return count
//}

func minMoves(nums []int) int {
	var minNum int = 2e9
	for _, num := range nums {
		if num < minNum {
			minNum = num
		}
	}
	var ans int
	for _, num := range nums {
		ans+=num-minNum
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
