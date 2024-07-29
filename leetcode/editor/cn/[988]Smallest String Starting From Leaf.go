/**
You are given the root of a binary tree where each node has a value in the
range [0, 25] representing the letters 'a' to 'z'.

 Return the lexicographically smallest string that starts at a leaf of this
tree and ends at the root.

 As a reminder, any shorter prefix of a string is lexicographically smaller.


 For example, "ab" is lexicographically smaller than "aba".


 A leaf of a node is a node that has no children.


 Example 1:


Input: root = [0,1,2,3,4,3,4]
Output: "dba"


 Example 2:


Input: root = [25,1,3,1,3,0,2]
Output: "adz"


 Example 3:


Input: root = [2,2,1,null,1,0,null,0]
Output: "abc"



 Constraints:


 The number of nodes in the tree is in the range [1, 8500].
 0 <= Node.val <= 25


 Related Topics 树 深度优先搜索 字符串 回溯 二叉树 👍 113 👎 0

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

func smallestFromLeaf(root *TreeNode) string {
	minS := string(rune(127))
	cur := ""
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		r := 'a' + node.Val
		cur = string(rune(r)) + cur

		if node.Left == nil && node.Right == nil {
			minS = min(minS, cur)
		}

		dfs(node.Left)
		dfs(node.Right)

		cur = cur[1:]
	}

	if root == nil {
		return ""
	}
	dfs(root)
	return minS
}

//leetcode submit region end(Prohibit modification and deletion)
