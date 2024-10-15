//Given a string s that contains parentheses and letters, remove the minimum
//number of invalid parentheses to make the input string valid.
//
// Return a list of unique strings that are valid with the minimum number of
//removals. You may return the answer in any order.
//
//
// Example 1:
//
//
//Input: s = "()())()"
//Output: ["(())()","()()()"]
//
//
// Example 2:
//
//
//Input: s = "(a)())()"
//Output: ["(a())()","(a)()()"]
//
//
// Example 3:
//
//
//Input: s = ")("
//Output: [""]
//
//
//
// Constraints:
//
//
// 1 <= s.length <= 25
// s consists of lowercase English letters and parentheses '(' and ')'.
// There will be at most 20 parentheses in s.
//
//
// Related Topics 广度优先搜索 字符串 回溯 👍 940 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
package main

func removeInvalidParentheses(s string) []string {
	numRmR := 0
	numRmL := 0

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			numRmL++
		} else if s[i] == ')' {
			if numRmL > 0 {
				numRmL--
			} else {
				numRmR++
			}
		}
	}

	var traverse func(pos int)
	sum := make([]string, 0)
	trimS := make([]byte, 0)
	traverse = func(pos int) {
		if numRmL == 0 && numRmR == 0 {
			cur := string(trimS) + s[pos:]
			if checkValidP301(cur) {
				sum = append(sum, cur)
			}
			return
		}

		storeLen := len(trimS)
		for i := pos; i < len(s); i++ {
			if s[i] == '(' && numRmL > 0 {
				numRmL--
				traverse(i + 1)
				numRmL++
			}

			if s[i] == ')' && numRmR > 0 {
				numRmR--
				traverse(i + 1)
				numRmR++
			}

			trimS = append(trimS, s[i])
		}
		trimS = trimS[:storeLen]
	}

	traverse(0)

	rlt := make([]string, 0)
	rltMap := make(map[string]struct{})

	for _, v := range sum {
		if _, ok := rltMap[v]; ok {
			continue
		}
		rltMap[v]= struct{}{}
		rlt = append(rlt, v)
	}

	return rlt
}

func checkValidP301(s string) bool {
	numL := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			numL++
		} else if s[i] == ')' {
			if numL > 0 {
				numL--
			} else {
				return false
			}
		}
	}

	if numL == 0 {
		return true
	} else {
		return false
	}
}

//leetcode submit region end(Prohibit modification and deletion)
