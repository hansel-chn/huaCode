/**
Given a string s and an integer k, return the length of the longest substring
of s such that the frequency of each character in this substring is greater than
or equal to k.

 if no such substring exists, return 0.


 Example 1:


Input: s = "aaabb", k = 3
Output: 3
Explanation: The longest substring is "aaa", as 'a' is repeated 3 times.


 Example 2:


Input: s = "ababbc", k = 2
Output: 5
Explanation: The longest substring is "ababb", as 'a' is repeated 2 times and
'b' is repeated 3 times.



 Constraints:


 1 <= s.length <= 10⁴
 s consists of only lowercase English letters.
 1 <= k <= 10⁵


 Related Topics 哈希表 字符串 分治 滑动窗口 👍 896 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
package main

func longestSubstring(s string, k int) int {
	var rlt int
	sMap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		sMap[s[i]]++
	}

	split := make(map[byte]struct{}, 0)
	for c, num := range sMap {
		if num < k {
			split[c] = struct{}{}
		}
	}
	if len(split) == 0 {
		return len(s)
	}

	splitStr := make([]string, 0)
	var start int
	for i := 0; i < len(s); i++ {
		if _, ok := split[s[i]]; ok {
			if s[start:i] != "" {
				splitStr = append(splitStr, s[start:i])
			}
				start = i + 1
		}
	}

	if s[start:] != "" {
		splitStr = append(splitStr, s[start:])
	}

	for _, v := range splitStr {
		rlt = max(rlt, longestSubstring(v, k))
	}
	return rlt
}

//leetcode submit region end(Prohibit modification and deletion)
