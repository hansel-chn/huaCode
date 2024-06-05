/**
Given the root of a binary tree, return all duplicate subtrees.

 For each kind of duplicate subtrees, you only need to return the root node of
any one of them.

 Two trees are duplicate if they have the same structure with the same node
values.


 Example 1:


Input: root = [1,2,3,4,null,2,4,null,null,4]
Output: [[2,4],[4]]


 Example 2:


Input: root = [2,1,1]
Output: [[1]]


 Example 3:


Input: root = [2,2,2,3,null,3,null]
Output: [[2,3],[3]]



 Constraints:


 The number of the nodes in the tree will be in the range [1, 5000]
 -200 <= Node.val <= 200


 Related Topics 树 深度优先搜索 哈希表 二叉树 👍 747 👎 0

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

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	rlt := make([]*TreeNode, 0)
	repeatMap := make(map[string]int)

	var traverse func(node *TreeNode) string
	traverse = func(node *TreeNode) string {

		if node == nil {
			return "#"
		}

		left := traverse(node.Left)
		right := traverse(node.Right)

		nodeStr := left + "," + right + "," + strconv.Itoa(node.Val)
		if v, ok := repeatMap[nodeStr]; ok && v == 1 {
			rlt = append(rlt, node)
		}
		repeatMap[nodeStr]++

		return nodeStr
	}
	traverse(root)
	return rlt
}

//leetcode submit region end(Prohibit modification and deletion)
// https://cs.stackexchange.com/questions/116655/which-tree-traversal-string-is-unique
// #,0,#,0,#
// 中序遍历会出错，可以表示两种结构，思考为什么
// 证明没看明白，但是大概这么想，满二叉树(full binary tree with distinguished leaf nodes)，preorder traverse和postorder traverse可以根据连续的两个# 确定已经到达根节点，数量需要与父节点有数量关系。
// 但是中序遍历不可以
