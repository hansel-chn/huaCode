/**
Given a string expression of numbers and operators, return all possible results
from computing all the different possible ways to group numbers and operators.
You may return the answer in any order.

 The test cases are generated such that the output values fit in a 32-bit
integer and the number of different results does not exceed 10⁴.


 Example 1:


Input: expression = "2-1-1"
Output: [0,2]
Explanation:
((2-1)-1) = 0
(2-(1-1)) = 2


 Example 2:


Input: expression = "2*3-4*5"
Output: [-34,-14,-10,-10,10]
Explanation:
(2*(3-(4*5))) = -34
((2*3)-(4*5)) = -14
((2*(3-4))*5) = -10
(2*((3-4)*5)) = -10
(((2*3)-4)*5) = 10



 Constraints:


 1 <= expression.length <= 20
 expression consists of digits and the operator '+', '-', and '*'.
 All the integer values in the input expression are in the range [0, 99].


 Related Topics 递归 记忆化搜索 数学 字符串 动态规划 👍 882 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
package main

import "strconv"

func diffWaysToCompute(expression string) []int {
	memo := make(map[string][]int)
	var divideAndConquer func(expression string) []int
	divideAndConquer = func(expression string) []int {
		rlt := make([]int, 0)

		if v, ok := memo[expression]; ok {
			return v
		}

		for i := 0; i < len(expression); i++ {
			c := expression[i]
			if c == '+' || c == '-' || c == '*' {
				left := divideAndConquer(expression[:i])
				right := divideAndConquer(expression[i+1:])

				switch c {
				case '+':
					for _, l := range left {
						for _, r := range right {
							rlt = append(rlt, l+r)
						}
					}
				case '-':
					for _, l := range left {
						for _, r := range right {
							rlt = append(rlt, l-r)
						}
					}
				case '*':
					for _, l := range left {
						for _, r := range right {
							rlt = append(rlt, l*r)
						}
					}
				}
			}
		}

		if len(rlt) == 0 {
			num, _ := strconv.Atoi(expression)
			rlt = append(rlt, num)
		}

		memo[expression] = rlt
		return rlt
	}
	return divideAndConquer(expression)
}

//leetcode submit region end(Prohibit modification and deletion)
