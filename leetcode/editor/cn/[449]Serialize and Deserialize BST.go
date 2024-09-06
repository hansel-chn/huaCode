//Serialization is converting a data structure or object into a sequence of
//bits so that it can be stored in a file or memory buffer, or transmitted across a
//network connection link to be reconstructed later in the same or another computer
//environment.
//
// Design an algorithm to serialize and deserialize a binary search tree. There
//is no restriction on how your serialization/deserialization algorithm should
//work. You need to ensure that a binary search tree can be serialized to a string,
//and this string can be deserialized to the original tree structure.
//
// The encoded string should be as compact as possible.
//
//
// Example 1:
// Input: root = [2,1,3]
//Output: [2,1,3]
//
// Example 2:
// Input: root = []
//Output: []
//
//
// Constraints:
//
//
// The number of nodes in the tree is in the range [0, 10⁴].
// 0 <= Node.val <= 10⁴
// The input tree is guaranteed to be a binary search tree.
//
//
// Related Topics 树 深度优先搜索 广度优先搜索 设计 二叉搜索树 字符串 二叉树 👍 555 👎 0

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
	"strings"
)

const sep = ","
const (
	maxV = 10e5
	minV = -1
)

type Codec struct {
}

//func Constructor() Codec {
//	return Codec{}
//}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	preOrder := make([]string, 0)
	var serializeHelper func(node *TreeNode)
	serializeHelper = func(node *TreeNode) {
		if node == nil {
			return
		}

		preOrder = append(preOrder, strconv.Itoa(node.Val))
		serializeHelper(node.Left)
		serializeHelper(node.Right)
	}
	serializeHelper(root)
	return strings.Join(preOrder, sep)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	preOrder := strings.Split(data, sep)

	var deserializeHelper func(mi, ma int) (node *TreeNode)
	deserializeHelper = func(mi, ma int) (node *TreeNode) {
		if len(preOrder) == 0 {
			return nil
		}

		nvStr := preOrder[0]
		nvInt, _ := strconv.Atoi(nvStr)

		if mi < nvInt && nvInt < ma {
			preOrder = preOrder[1:]
			node = &TreeNode{}
			node.Val = nvInt
			node.Left = deserializeHelper(mi, nvInt)
			node.Right = deserializeHelper(nvInt, ma)
			return node
		} else {
			return nil
		}
	}

	return deserializeHelper(minV, maxV)
}

//// preOrder contain null
//const sep = ","
//
//type Codec struct {
//}
//
//func Constructor() Codec {
//	return Codec{}
//}
//
//// Serializes a tree to a single string.
//func (this *Codec) serialize(root *TreeNode) string {
//	preOrder := make([]string, 0)
//	var serializeHelper func(node *TreeNode)
//	serializeHelper = func(node *TreeNode) {
//		if node == nil {
//			preOrder = append(preOrder, "")
//			return
//		}
//
//		preOrder = append(preOrder, strconv.Itoa(node.Val))
//		serializeHelper(node.Left)
//		serializeHelper(node.Right)
//	}
//	serializeHelper(root)
//	return strings.Join(preOrder, sep)
//}
//
//// Deserializes your encoded data to tree.
//func (this *Codec) deserialize(data string) *TreeNode {
//	preOrder := strings.Split(data, sep)
//
//	var deserializeHelper func() (node *TreeNode)
//	deserializeHelper = func() (node *TreeNode) {
//		if len(preOrder) == 0 {
//			return
//		}
//
//		nvStr := preOrder[0]
//		preOrder=preOrder[1:]
//
//		if nvStr == "" {
//			return nil
//		}
//
//		node = &TreeNode{}
//		nvInt, _ := strconv.Atoi(nvStr)
//		node.Val = nvInt
//
//		node.Left = deserializeHelper()
//		node.Right = deserializeHelper()
//		return node
//	}
//
//	return deserializeHelper()
//}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */
//leetcode submit region end(Prohibit modification and deletion)
