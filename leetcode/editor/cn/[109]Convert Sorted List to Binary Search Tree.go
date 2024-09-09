//Given the head of a singly linked list where elements are sorted in ascending
//order, convert it to a height-balanced binary search tree.
//
//
// Example 1:
//
//
//Input: head = [-10,-3,0,5,9]
//Output: [0,-3,9,-10,null,5]
//Explanation: One possible answer is [0,-3,9,-10,null,5], which represents the
//shown height balanced BST.
//
//
// Example 2:
//
//
//Input: head = []
//Output: []
//
//
//
// Constraints:
//
//
// The number of nodes in head is in the range [0, 2 * 10⁴].
// -10⁵ <= Node.val <= 10⁵
//
//
// Related Topics 树 二叉搜索树 链表 分治 二叉树 👍 913 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package main

func sortedListToBST(head *ListNode) *TreeNode {
	cur := head
	var count int
	for cur != nil {
		cur = cur.Next
		count++
	}

	var build func(l, r int) *TreeNode
	build = func(l, r int) *TreeNode {
		if l > r {
			return nil
		}

		node := &TreeNode{}
		m := l + (r-l)/2
		node.Left = build(l, m-1)
		node.Val = head.Val
		head = head.Next
		node.Right = build(m+1, r)
		return node
	}
	return build(1, count)
}

//leetcode submit region end(Prohibit modification and deletion)
