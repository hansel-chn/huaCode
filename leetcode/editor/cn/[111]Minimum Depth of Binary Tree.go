/**
Given a binary tree, find its minimum depth.

 The minimum depth is the number of nodes along the shortest path from the root
node down to the nearest leaf node.

 Note: A leaf is a node with no children.


 Example 1:


Input: root = [3,9,20,null,null,15,7]
Output: 2


 Example 2:


Input: root = [2,null,3,null,4,null,5,null,6]
Output: 5



 Constraints:


 The number of nodes in the tree is in the range [0, 10⁵].
 -1000 <= Node.val <= 1000


 Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 1175 👎 0

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

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var queue []*TreeNode
	queue = append(queue, root)
	depth := 0
	for len(queue) > 0 {
		depth++
		for _, v := range queue {
			node := v
			queue = queue[1:]
			if node.Left == nil && node.Right == nil {
				return depth
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return depth
}

//leetcode submit region end(Prohibit modification and deletion)
