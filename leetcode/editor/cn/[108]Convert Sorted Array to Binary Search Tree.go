//Given an integer array nums where the elements are sorted in ascending order,
//convert it to a height-balanced binary search tree.
//
//
// Example 1:
//
//
//Input: nums = [-10,-3,0,5,9]
//Output: [0,-3,9,-10,null,5]
//Explanation: [0,-10,5,null,-3,null,9] is also accepted:
//
//
//
// Example 2:
//
//
//Input: nums = [1,3]
//Output: [3,1]
//Explanation: [1,null,3] and [3,1] are both height-balanced BSTs.
//
//
//
// Constraints:
//
//
// 1 <= nums.length <= 10⁴
// -10⁴ <= nums[i] <= 10⁴
// nums is sorted in a strictly increasing order.
//
//
// Related Topics 树 二叉搜索树 数组 分治 二叉树 👍 1551 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package main

func sortedArrayToBST(nums []int) *TreeNode {
	var build func(nums []int) *TreeNode
	build = func(nums []int) *TreeNode {
		if len(nums) == 0 {
			return nil
		}
		node := &TreeNode{}
		idx := len(nums) >> 1
		node.Val = nums[idx]
		node.Left = build(nums[:idx])
		node.Right = build(nums[idx+1:])
		return node
	}
	return build(nums)
}

//leetcode submit region end(Prohibit modification and deletion)
