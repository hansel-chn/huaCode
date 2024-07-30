/**
Given the root of a binary tree, each node in the tree has a distinct value.

 After deleting all nodes with a value in to_delete, we are left with a forest (
a disjoint union of trees).

 Return the roots of the trees in the remaining forest. You may return the
result in any order.


 Example 1:


Input: root = [1,2,3,4,5,6,7], to_delete = [3,5]
Output: [[1,2,null,4],[6],[7]]


 Example 2:


Input: root = [1,2,4,null,3], to_delete = [3]
Output: [[1,2,4]]



 Constraints:


 The number of nodes in the given tree is at most 1000.
 Each node has a distinct value between 1 and 1000.
 to_delete.length <= 1000
 to_delete contains distinct values between 1 and 1000.


 Related Topics 树 深度优先搜索 数组 哈希表 二叉树 👍 336 👎 0

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

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	delMap := make(map[int]struct{})
	for _, v := range to_delete {
		delMap[v] = struct{}{}
	}

	res := make([]*TreeNode, 0)
	if _, ok := delMap[root.Val]; !ok {
		res = append(res, root)
	}

	var traverse func(node *TreeNode) *TreeNode
	traverse = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}

		node.Left = traverse(node.Left)
		node.Right = traverse(node.Right)

		if _, ok := delMap[node.Val]; ok {
			if node.Left != nil {
				res = append(res, node.Left)
			}

			if node.Right != nil {
				res = append(res, node.Right)
			}
			return nil
		}

		return node
	}

	traverse(root)
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
