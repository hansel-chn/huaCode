//You are given a string s that consists of only digits.
//
// Check if we can split s into two or more non-empty substrings such that the
//numerical values of the substrings are in descending order and the difference
//between numerical values of every two adjacent substrings is equal to 1.
//
//
// For example, the string s = "0090089" can be split into ["0090", "089"] with
//numerical values [90,89]. The values are in descending order and adjacent
//values differ by 1, so this way is valid.
// Another example, the string s = "001" can be split into ["0", "01"], ["00",
//"1"], or ["0", "0", "1"]. However all the ways are invalid because they have
//numerical values [0,1], [0,1], and [0,0,1] respectively, all of which are not in
//descending order.
//
//
// Return true if it is possible to split s as described above, or false
//otherwise.
//
// A substring is a contiguous sequence of characters in a string.
//
//
// Example 1:
//
//
//Input: s = "1234"
//Output: false
//Explanation: There is no valid way to split s.
//
//
// Example 2:
//
//
//Input: s = "050043"
//Output: true
//Explanation: s can be split into ["05", "004", "3"] with numerical values [5,4
//,3].
//The values are in descending order with adjacent values differing by 1.
//
//
// Example 3:
//
//
//Input: s = "9080701"
//Output: false
//Explanation: There is no valid way to split s.
//
//
//
// Constraints:
//
//
// 1 <= s.length <= 20
// s only consists of digits.
//
//
// Related Topics 字符串 回溯 👍 46 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
package main

func splitString(s string) bool {
	sLen := len(s)
	rlt := false

	var traverse func(pos int, preNum int)
	traverse = func(pos int, preNum int) {
		if pos == sLen {
			rlt = true
			return
		}

		cur := 0
		for i := pos; i < sLen; i++ {
			cur = cur*10 + int(s[i]-'0')
			//optimize
			if cur > preNum {
				return
			}

			if preNum-cur == 1 {
				traverse(i+1, cur)
			}

			if rlt == true {
				return
			}
		}
	}

	firstNum := 0
	for i := 0; i < sLen-1; i++ {
		firstNum = firstNum*10 + int(s[i]-'0')
		if rlt == true {
			return true
		} else {
			traverse(i+1, firstNum)
		}
	}
	return rlt
}

//leetcode submit region end(Prohibit modification and deletion)
