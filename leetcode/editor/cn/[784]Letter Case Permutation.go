//Given a string s, you can transform every letter individually to be lowercase
//or uppercase to create another string.
//
// Return a list of all possible strings we could create. Return the output in
//any order.
//
//
// Example 1:
//
//
//Input: s = "a1b2"
//Output: ["a1b2","a1B2","A1b2","A1B2"]
//
//
// Example 2:
//
//
//Input: s = "3z4"
//Output: ["3z4","3Z4"]
//
//
//
// Constraints:
//
//
// 1 <= s.length <= 12
// s consists of lowercase English letters, uppercase English letters, and
//digits.
//
//
// Related Topics 位运算 字符串 回溯 👍 586 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
package main

const diff784 = 'a' - 'A'

func letterCasePermutation(s string) []string {
	sLen := len(s)
	bs := []byte(s)

	resMap := make(map[string]struct{})
	var traverse func(pos int)
	traverse = func(pos int) {
		if pos == sLen {
			resMap[string(bs)] = struct{}{}
			return
		}

		if bs[pos] >= '0' && bs[pos] <= '9' {
			traverse(pos + 1)
		} else {
			traverse(pos + 1)

			bs[pos] = transform784(bs[pos])
			traverse(pos + 1)
			bs[pos] = transform784(bs[pos])
		}
	}
	traverse(0)

	res := make([]string, 0)
	for v := range resMap {
		res = append(res, v)
	}
	return res
}

func transform784(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b + diff784
	} else {
		return b - diff784
	}
}

//leetcode submit region end(Prohibit modification and deletion)
