/**
Given the root of a binary tree, imagine yourself standing on the right side of
it, return the values of the nodes you can see ordered from top to bottom.


 Example 1:


Input: root = [1,2,3,null,5,null,4]
Output: [1,3,4]


 Example 2:


Input: root = [1,null,3]
Output: [1,3]


 Example 3:


Input: root = []
Output: []



 Constraints:


 The number of nodes in the tree is in the range [0, 100].
 -100 <= Node.val <= 100


 Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 1093 👎 0

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

func rightSideView(root *TreeNode) []int {
	view := make([]int, 0)
	depth := -1
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		depth++
		if depth == len(view) {
			view = append(view, node.Val)
		} else {
			view[depth] = node.Val
		}

		dfs(node.Left)
		dfs(node.Right)
		depth--
	}
	dfs(root)
	return view
}

//leetcode submit region end(Prohibit modification and deletion)
