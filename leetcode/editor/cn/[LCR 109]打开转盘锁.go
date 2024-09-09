//English description is not available for the problem. Please switch to
//Chinese.
//
// Related Topics 广度优先搜索 数组 哈希表 字符串 👍 47 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
package main

func openLock(deadends []string, target string) int {
	checkDead := make(map[string]struct{}, len(deadends))
	for _, dead := range deadends {
		checkDead[dead] = struct{}{}
	}

	used := make(map[string]struct{})
	rotation := 0

	queue := make([]string, 0)
	queue = append(queue, "0000")

	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			cur := queue[0]
			queue = queue[1:]
			if _, ok := checkDead[cur]; ok {
				continue
			}

			if _, ok := used[cur]; ok {
				continue
			}
			used[cur] = struct{}{}

			if target == cur {
				return rotation
			}

			for _, adjacent := range optimizeDirs109(cur) {
				queue = append(queue, adjacent)
			}

			//for j := 0; j < 4; j++ {
			//	up := plusOne(cur, j)
			//	queue = append(queue, up)
			//	down := minusOne(cur, j)
			//	queue = append(queue, down)
			//}
		}
		rotation++
	}
	return -1
}

func plusOne(s string, j int) string {
	ch := []rune(s)
	if ch[j] == '9' {
		ch[j] = '0'
	} else {
		ch[j] += 1
	}
	return string(ch)
}

// 将 s[i] 向下拨动一次
func minusOne(s string, j int) string {
	ch := []rune(s)
	if ch[j] == '0' {
		ch[j] = '9'
	} else {
		ch[j] -= 1
	}
	return string(ch)
}

func optimizeDirs109(src string) (rlt []string) {
	rlt = make([]string, 0)
	for i := 0; i < 4; i++ {
		srcRune := []rune(src)
		if srcRune[i] == '0' {
			srcRune[i] = '9'
			rlt = append(rlt, string(srcRune))
			srcRune[i] = '1'
			rlt = append(rlt, string(srcRune))
			continue
		}

		if srcRune[i] == '9' {
			srcRune[i] = '0'
			rlt = append(rlt, string(srcRune))
			srcRune[i] = '8'
			rlt = append(rlt, string(srcRune))
			continue
		}

		srcRune[i] += 1
		rlt = append(rlt, string(srcRune))
		srcRune[i] -= 2
		rlt = append(rlt, string(srcRune))
	}
	return
}

func dirs109(src string) (rlt []string) {
	rlt = make([]string, 0)
	for i := 0; i < 4; i++ {
		pos := src[i]
		if pos == '0' {
			rlt = append(rlt, src[:i]+"9"+src[i+1:])
			rlt = append(rlt, src[:i]+"1"+src[i+1:])
			continue
		}

		if pos == '9' {
			rlt = append(rlt, src[:i]+"8"+src[i+1:])
			rlt = append(rlt, src[:i]+"0"+src[i+1:])
			continue
		}

		rlt = append(rlt, src[:i]+string([]byte{pos + 1})+src[i+1:])
		rlt = append(rlt, src[:i]+string([]byte{pos - 1})+src[i+1:])
	}

	return rlt
}

//leetcode submit region end(Prohibit modification and deletion)
