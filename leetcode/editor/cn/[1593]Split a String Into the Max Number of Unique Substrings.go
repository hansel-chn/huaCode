//Given a string s, return the maximum number of unique substrings that the
//given string can be split into.
//
// You can split string s into any list of non-empty substrings, where the
//concatenation of the substrings forms the original string. However, you must split
//the substrings such that all of them are unique.
//
// A substring is a contiguous sequence of characters within a string.
//
//
// Example 1:
//
//
//Input: s = "ababccc"
//Output: 5
//Explanation: One way to split maximally is ['a', 'b', 'ab', 'c', 'cc'].
//Splitting like ['a', 'b', 'a', 'b', 'c', 'cc'] is not valid as you have 'a' and 'b'
//multiple times.
//
//
// Example 2:
//
//
//Input: s = "aba"
//Output: 2
//Explanation: One way to split maximally is ['a', 'ba'].
//
//
// Example 3:
//
//
//Input: s = "aa"
//Output: 1
//Explanation: It is impossible to split the string any further.
//
//
//
// Constraints:
//
//
// 1 <= s.length <= 16
// s contains only lower case English letters.
//
//
// Related Topics 哈希表 字符串 回溯 👍 62 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
package main

func maxUniqueSplit(s string) int {
	sLen := len(s)
	checkMap := make(map[string]struct{})
	maxSplit := -1
	curSplit := 0

	var traverse func(pos int)
	traverse = func(pos int) {
		if pos == sLen {
			maxSplit = max(maxSplit, curSplit)
		}

		for i := pos; i < sLen; i++ {
			cur := s[pos : i+1]
			if _, ok := checkMap[cur]; !ok {
				checkMap[cur] = struct{}{}
				curSplit++
				traverse(i + 1)
				delete(checkMap, cur)
				curSplit--
			}
		}
	}

	traverse(0)
	return maxSplit
}

//leetcode submit region end(Prohibit modification and deletion)
