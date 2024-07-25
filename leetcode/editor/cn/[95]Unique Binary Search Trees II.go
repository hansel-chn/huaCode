/**
Given an integer n, return all the structurally unique BST's (binary search
trees), which has exactly n nodes of unique values from 1 to n. Return the answer
in any order.


 Example 1:


Input: n = 3
Output: [[1,null,2,null,3],[1,null,3,2],[2,1,3],[3,1,null,null,2],[3,2,null,1]]


 Example 2:


Input: n = 1
Output: [[1]]



 Constraints:


 1 <= n <= 8


 Related Topics 树 二叉搜索树 动态规划 回溯 二叉树 👍 1562 👎 0

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

func generateTrees(n int) []*TreeNode {
	memo := make(map[[2]int][]*TreeNode)
	var getTree func(start, end int) []*TreeNode
	getTree = func(start, end int) []*TreeNode {
		if v, ok := memo[[2]int{start, end}]; ok {
			return v
		}

		if start > end {
			return []*TreeNode{nil}
		}

		res := make([]*TreeNode, 0)
		for i := start; i <= end; i++ {
			res = append(res, crossCombine(i, getTree(start, i-1), getTree(i+1, end))...)
		}

		memo[[2]int{start, end}] = res
		return res
	}
	return getTree(1, n)
}

func crossCombine(v int, lefts []*TreeNode, rights []*TreeNode) []*TreeNode {
	res := make([]*TreeNode, 0)
	for _, left := range lefts {
		for _, right := range rights {
			res = append(res, &TreeNode{
				Val:   v,
				Left:  left,
				Right: right,
			})
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
