//You are given an integer array nums and a positive integer k.
//
// The value of a sequence seq of size 2 * x is defined as:
//
//
// (seq[0] OR seq[1] OR ... OR seq[x - 1]) XOR (seq[x] OR seq[x + 1] OR ... OR
//seq[2 * x - 1]).
//
//
// Return the maximum value of any subsequence of nums having size 2 * k.
//
//
// Example 1:
//
//
// Input: nums = [2,6,7], k = 1
//
//
// Output: 5
//
// Explanation:
//
// The subsequence [2, 7] has the maximum value of 2 XOR 7 = 5.
//
// Example 2:
//
//
// Input: nums = [4,2,5,6,7], k = 2
//
//
// Output: 2
//
// Explanation:
//
// The subsequence [4, 5, 6, 7] has the maximum value of (4 OR 5) XOR (6 OR 7) =
// 2.
//
//
// Constraints:
//
//
// 2 <= nums.length <= 400
// 1 <= nums[i] < 2⁷
// 1 <= k <= nums.length / 2
//
//
// Related Topics 位运算 数组 动态规划 👍 9 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
package main

import (
	"math"
)

func maxValue(nums []int, k int) int {
	// 注意边界条件，为什么这么设置。
	// 什么时候需要初始化为true
	// 在i个数字间选择j个数字，能够构成的结果，存在1<<7大小的切片里,true/false标识是否能构成。
	//
	//for i := 0; i < len(dpPrefix); i++ {
	//	dpPrefix[i][0][0] = true
	//}
	// 理解：
	// 1： i个数字中挑选0个数字是否能构成0（可以）,
	//     i个数字中挑选0个数字是否能构成除0外的数字（不可以）,
	// 2： 与此相对0个数字中挑选j个数字始终构不成任何数字（dpPrefix[0][0][0]除外）
	//     后者若不遵守会出现如，
	//     dpPrefix[1][1][0]=dpPrefix[0][1][0]=true，dpPrefix[2][2][z]会出现基于dpPrefix[1][1][0]=true做出的错误的判断
	// 3： 实际有j <= i;的限制,受斜线上侧数字的影响
	// 4： 思考具体情况，dpPrefix[5][1]可以构成哪些数字，思考可知应基于dpPrefix[4][0]做判断，且理应只基于dpPrefix[4][0][0],
	// 就是前四个数字不选择，从第五个数字进行第一个选择，基于之前的状态。即前四个数字选择0个数字组成0是可以真实出现的。
	// if dpPrefix[i-1][j-1][z] {
	// 	dpPrefix[i][j][z|nums[i-1]] = true
	// }
	//	有且只有dpPrefix[i][0][0] = true必须成立
	// 5： 但是从0个数字里面选择j个数字，不可以组成有意义的数字。（反向证明）。若上述描述有意义，可以推出从大于0个数字选择j个数字，都可以组成上述组成的有意义的数字

	//  初始状态有种有点道理又冥冥之中的感觉

	dpPrefix := make([][][]bool, len(nums)-k+1)
	for i := range dpPrefix {
		dpPrefix[i] = make([][]bool, k+1)
		for j := range dpPrefix[i] {
			dpPrefix[i][j] = make([]bool, 1<<7)
		}
	}

	for i := 0; i < len(dpPrefix); i++ {
		dpPrefix[i][0][0] = true
	}

	for i := 1; i < len(dpPrefix); i++ {
		for j := 1; j < len(dpPrefix[0]) && j <= i; j++ {
			copy(dpPrefix[i][j], dpPrefix[i-1][j])
			for z := 0; z < 1<<7; z++ {
				if dpPrefix[i-1][j-1][z] {
					dpPrefix[i][j][z|nums[i-1]] = true
				}
			}
		}
	}

	dpSuf := make([][][]bool, len(nums)-k+1)
	for i := range dpSuf {
		dpSuf[i] = make([][]bool, k+1)
		for j := range dpSuf[i] {
			dpSuf[i][j] = make([]bool, 1<<7)
		}
	}

	for i := 0; i < len(dpSuf); i++ {
		dpSuf[i][0][0] = true
	}

	for i := 1; i < len(dpSuf); i++ {
		for j := 1; j < len(dpSuf[0]) && j <= i; j++ {
			copy(dpSuf[i][j], dpSuf[i-1][j])
			for z := 0; z < 1<<7; z++ {
				if dpSuf[i-1][j-1][z] {
					dpSuf[i][j][z|nums[len(nums)-i]] = true
				}
			}
		}
	}

	rlt := math.MinInt64
	for i := k; i < len(nums)-k+1; i++ {
		for v1, exist1 := range dpPrefix[i][k] {
			if exist1 {
				for v2, exist12 := range dpSuf[len(nums)-i][k] {
					if exist12 {
						rlt = max(rlt, v1^v2)
					}
				}
			}
		}
	}
	return rlt
}

//leetcode submit region end(Prohibit modification and deletion)
