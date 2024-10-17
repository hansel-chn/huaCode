//You have n tiles, where each tile has one letter tiles[i] printed on it.
//
// Return the number of possible non-empty sequences of letters you can make
//using the letters printed on those tiles.
//
//
// Example 1:
//
//
//Input: tiles = "AAB"
//Output: 8
//Explanation: The possible sequences are "A", "B", "AA", "AB", "BA", "AAB",
//"ABA", "BAA".
//
//
// Example 2:
//
//
//Input: tiles = "AAABBC"
//Output: 188
//
//
// Example 3:
//
//
//Input: tiles = "V"
//Output: 1
//
//
//
// Constraints:
//
//
// 1 <= tiles.length <= 7
// tiles consists of uppercase English letters.
//
//
// Related Topics 哈希表 字符串 回溯 计数 👍 276 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
package main

import "sort"

//func numTilePossibilities(tiles string) int {
//	tileLen := len(tiles)
//	rltMap := make(map[string]struct{})
//	visited := make([]bool, len(tiles))
//	cur := ""
//
//	var traverse func()
//	traverse = func() {
//		for i := 0; i < tileLen; i++ {
//			if visited[i] == true {
//				continue
//			}
//
//			cur += string(tiles[i])
//			if _, ok := rltMap[cur]; ok {
//				cur = cur[:len(cur)-1]
//				continue
//			}
//
//			rltMap[cur]= struct{}{}
//			visited[i] = true
//			traverse()
//			cur = cur[:len(cur)-1]
//			visited[i] = false
//		}
//	}
//	traverse()
//	return len(rltMap)
//}

func numTilePossibilities(tiles string) int {
	tilesBytes := []byte(tiles)
	sort.Slice(tilesBytes, func(i, j int) bool {
		return tilesBytes[i] < tilesBytes[j]
	})
	tiles = string(tilesBytes)
	tileLen := len(tiles)

	//rltMap := make(map[string]struct{})
	visited := make([]bool, len(tiles))
	rlt := 0

	var traverse func()
	traverse = func() {
		for i := 0; i < tileLen; i++ {
			if visited[i] == true {
				continue
			}

			if i > 0 && tiles[i-1] == tiles[i] && visited[i-1] == false {
				continue
			}

			rlt++
			visited[i] = true
			traverse()
			visited[i] = false

		}
	}
	traverse()
	return rlt
}

//leetcode submit region end(Prohibit modification and deletion)
