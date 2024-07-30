/**
Given two integer arrays inorder and postorder where inorder is the inorder
traversal of a binary tree and postorder is the postorder traversal of the same
tree, construct and return the binary tree.


 Example 1:


Input: inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
Output: [3,9,20,null,null,15,7]


 Example 2:


Input: inorder = [-1], postorder = [-1]
Output: [-1]



 Constraints:


 1 <= inorder.length <= 3000
 postorder.length == inorder.length
 -3000 <= inorder[i], postorder[i] <= 3000
 inorder and postorder consist of unique values.
 Each value of postorder also appears in inorder.
 inorder is guaranteed to be the inorder traversal of the tree.
 postorder is guaranteed to be the postorder traversal of the tree.


 Related Topics 树 数组 哈希表 分治 二叉树 👍 1249 👎 0

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

//func buildTree(inorder []int, postorder []int) *TreeNode {
//	inMap := make(map[int]int)
//	for i, v := range inorder {
//		inMap[v] = i
//	}
//
//	var build func(postS, postE, inS, inE int) *TreeNode
//	build = func(postS, postE, inS, inE int) *TreeNode {
//		if postS > postE {
//			return nil
//		}
//
//		rootV :=postorder[postE]
//		root := &TreeNode{Val: rootV}
//		iIdx := inMap[rootV]
//
//		root.Left = build(postS, postS+iIdx-inS-1, inS, iIdx-1)
//		root.Right = build(postS+iIdx-inS, postE-1, iIdx+1, postE)
//		return root
//	}
//	return build(0, len(postorder)-1, 0, len(inorder)-1)
//}

//leetcode submit region end(Prohibit modification and deletion)
