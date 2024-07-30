/**
Given an integer n, return a list of all possible full binary trees with n
nodes. Each node of each tree in the answer must have Node.val == 0.

 Each element of the answer is the root node of one possible tree. You may
return the final list of trees in any order.

 A full binary tree is a binary tree where each node has exactly 0 or 2
children.


 Example 1:


Input: n = 7
Output: [[0,0,0,null,null,0,0,null,null,0,0],[0,0,0,null,null,0,0,0,0],[0,0,0,0,
0,0,0],[0,0,0,0,0,null,null,null,null,0,0],[0,0,0,0,0,null,null,0,0]]


 Example 2:


Input: n = 3
Output: [[0,0,0]]



 Constraints:


 1 <= n <= 20


 Related Topics 树 递归 记忆化搜索 动态规划 二叉树 👍 410 👎 0

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

func allPossibleFBT(n int) []*TreeNode {
	memo := make(map[int][]*TreeNode)
	var build func(n int) []*TreeNode
	build = func(n int) []*TreeNode {
		if v, ok := memo[n]; ok {
			return v
		}

		if n == 1 {
			return []*TreeNode{&TreeNode{Val: 0}}
		}

		res := make([]*TreeNode, 0)
		for i := 1; i < n-1; i = i + 2 {
			lNodes := build(i)
			rNodes := build(n - i - 1)
			for _, l := range lNodes {
				for _, r := range rNodes {
					cur := &TreeNode{Val: 0}
					cur.Left = l
					cur.Right = r
					res = append(res, cur)
				}
			}
		}

		memo[n] = res
		return res
	}
	return build(n)
}

//leetcode submit region end(Prohibit modification and deletion)
