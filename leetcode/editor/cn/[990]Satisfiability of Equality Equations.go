/**
You are given an array of strings equations that represent relationships
between variables where each string equations[i] is of length 4 and takes one of two
different forms: "xi==yi" or "xi!=yi".Here, xi and yi are lowercase letters (not
necessarily different) that represent one-letter variable names.

 Return true if it is possible to assign integers to variable names so as to
satisfy all the given equations, or false otherwise.


 Example 1:


Input: equations = ["a==b","b!=a"]
Output: false
Explanation: If we assign say, a = 1 and b = 1, then the first equation is
satisfied, but not the second.
There is no way to assign the variables to satisfy both equations.


 Example 2:


Input: equations = ["b==a","a==b"]
Output: true
Explanation: We could assign a = 1 and b = 1 to satisfy both equations.



 Constraints:


 1 <= equations.length <= 500
 equations[i].length == 4
 equations[i][0] is a lowercase letter.
 equations[i][1] is either '=' or '!'.
 equations[i][2] is '='.
 equations[i][3] is a lowercase letter.


 Related Topics 并查集 图 数组 字符串 👍 329 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
package main

func equationsPossible(equations []string) bool {
	uf := NewUF990(26)
	for _, equation := range equations {
		if equation[1] == '=' {
			uf.Union(charToInt990(equation[0]), charToInt990(equation[3]))
		}
	}

	for _, equation := range equations {
		if equation[1] == '!' {
			if uf.IfConnect(charToInt990(equation[0]), charToInt990(equation[3])) {
				return false
			}
		}
	}
	return true
}

func charToInt990(x uint8) int {
	return int(x - 'a')
}

func NewUF990(n int) UF990 {
	parent := make([]int, n)

	for i := range parent {
		parent[i] = i
	}

	return UF990{
		count:  n,
		Parent: parent,
	}
}

type UF990 struct {
	count  int
	Parent []int
}

func (u UF990) RegionCount() int {
	return u.count
}

func (u *UF990) Union(i, j int) {
	iTopP := u.FindCompressTopParent(i)
	jTopP := u.FindCompressTopParent(j)
	if iTopP == jTopP {
		return
	}

	(*u).Parent[iTopP] = jTopP

	u.count--
}

func (u *UF990) FindCompressTopParent(i int) int {
	if u.Parent[i] == i {
		return i
	}

	topP := u.FindCompressTopParent(u.Parent[i])

	u.Parent[i] = topP
	return topP
}

func (u UF990) IfConnect(i, j int) bool {
	if u.FindCompressTopParent(i) == u.FindCompressTopParent(j) {
		return true
	} else {
		return false
	}
}
