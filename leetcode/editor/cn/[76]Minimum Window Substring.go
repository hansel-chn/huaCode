/**
Given two strings s and t of lengths m and n respectively, return the minimum
window substring of s such that every character in t (including duplicates) is
included in the window. If there is no such substring, return the empty string "".


 The testcases will be generated such that the answer is unique.


 Example 1:


Input: s = "ADOBECODEBANC", t = "ABC"
Output: "BANC"
Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C'
from string t.


 Example 2:


Input: s = "a", t = "a"
Output: "a"
Explanation: The entire string s is the minimum window.


 Example 3:


Input: s = "a", t = "aa"
Output: ""
Explanation: Both 'a's from t must be included in the window.
Since the largest window of s only has one 'a', return empty string.



 Constraints:


 m == s.length
 n == t.length
 1 <= m, n <= 10⁵
 s and t consist of uppercase and lowercase English letters.



 Follow up: Could you find an algorithm that runs in O(m + n) time?

 Related Topics 哈希表 字符串 滑动窗口 👍 2902 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
package main

import (
	"math"
)

func minWindow(s string, t string) string {
	tMap := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		tMap[t[i]] += 1
	}

	winMap := make(map[byte]int)
	var left, right, valid, start int
	length := math.MaxInt

	for right < len(s) {
		addV := s[right]
		right++
		if _, ok := tMap[addV]; ok {
			winMap[addV]++
			if winMap[addV] == tMap[addV] {
				valid++
			}
		}

		for valid == len(tMap) {
			if right-left < length {
				start = left
				length =  right-left
			}
			if _, ok := tMap[s[left]]; ok {
				if winMap[s[left]] == tMap[s[left]] {
					valid--
				}
				winMap[s[left]]--
			}
			left++
		}
	}
	if length == math.MaxInt {
		return ""
	} else {
		return s[start : start+length]
	}
}

//leetcode submit region end(Prohibit modification and deletion)
