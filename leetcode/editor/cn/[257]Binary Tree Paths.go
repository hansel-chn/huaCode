/**
Given the root of a binary tree, return all root-to-leaf paths in any order.

 A leaf is a node with no children.


 Example 1:


Input: root = [1,2,3,null,5]
Output: ["1->2->5","1->3"]


 Example 2:


Input: root = [1]
Output: ["1"]



 Constraints:


 The number of nodes in the tree is in the range [1, 100].
 -100 <= Node.val <= 100


 Related Topics 树 深度优先搜索 字符串 回溯 二叉树 👍 1142 👎 0

*/

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

import (
	"strconv"
)

func binaryTreePaths(root *TreeNode) []string {
	res := make([]string, 0)

	type rec struct {
		node  *TreeNode
		trail string
	}

	stack := make([]rec, 0)
	stack = append(stack, rec{
		node:  root,
		trail: strconv.Itoa(root.Val),
	})

	for len(stack) > 0 {
		r := stack[0]
		stack = stack[1:]

		if r.node.Left == nil && r.node.Right == nil {
			res = append(res, r.trail)
		}

		if r.node.Left != nil {
			stack = append(stack, rec{
				node:  r.node.Left,
				trail: r.trail + "->" + strconv.Itoa(r.node.Left.Val),
			})
		}

		if r.node.Right != nil {
			stack = append(stack, rec{
				node:  r.node.Right ,
				trail: r.trail + "->" + strconv.Itoa(r.node.Right.Val),
			})
		}
	}

	return res
}

//func binaryTreePaths(root *TreeNode) []string {
//	rlt := make([]string, 0)
//	trail := make([]string, 0)
//	var traverse func(node *TreeNode)
//	traverse = func(node *TreeNode) {
//		if node == nil {
//			return
//		}
//
//		trail = append(trail, strconv.Itoa(node.Val))
//		if node.Left == nil && node.Right == nil {
//			rlt = append(rlt, strings.Join(trail, "->"))
//		}
//
//		traverse(node.Left)
//		traverse(node.Right)
//		trail = trail[:len(trail)-1]
//	}
//	traverse(root)
//	return rlt
//}

//leetcode submit region end(Prohibit modification and deletion)
