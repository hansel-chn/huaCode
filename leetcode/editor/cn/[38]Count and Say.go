/**
The count-and-say sequence is a sequence of digit strings defined by the
recursive formula:


 countAndSay(1) = "1"
 countAndSay(n) is the way you would "say" the digit string from countAndSay(n-1
), which is then converted into a different digit string.


 To determine how you "say" a digit string, split it into the minimal number of
substrings such that each substring contains exactly one unique digit. Then for
each substring, say the number of digits, then say the digit. Finally,
concatenate every said digit.

 For example, the saying and conversion for digit string "3322251":

 Given a positive integer n, return the nᵗʰ term of the count-and-say sequence.



 Example 1:


Input: n = 1
Output: "1"
Explanation: This is the base case.


 Example 2:


Input: n = 4
Output: "1211"
Explanation:
countAndSay(1) = "1"
countAndSay(2) = say "1" = one 1 = "11"
countAndSay(3) = say "11" = two 1's = "21"
countAndSay(4) = say "21" = one 2 + one 1 = "12" + "11" = "1211"



 Constraints:


 1 <= n <= 30


 Related Topics 字符串 👍 1072 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
package main

import (
	"strconv"
	"strings"
)

func countAndSay(n int) string {
	prev := "1"
	for i := 2; i <= n; i++ {
		str := strings.Builder{}
		start, j := 0, 0
		for ; j < len(prev); j++ {
			// 现在写的不好
			// 感觉放到内层两层循环更优雅，不会重复写代码
			// 通过两个循环的内层循环把对应字符不相等和到达边界两种情况统一起来处理，代码更优雅
			if prev[start] != prev[j] {
				str.WriteString(strconv.Itoa(j - start))
				str.WriteByte(prev[start])
				start = j
			}
		}
		str.WriteString(strconv.Itoa(j - start))
		str.WriteByte(prev[start])
		prev = str.String()
	}
	return prev
}

//leetcode submit region end(Prohibit modification and deletion)
