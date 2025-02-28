//A trie (pronounced as "try") or prefix tree is a tree data structure used to
//efficiently store and retrieve keys in a dataset of strings. There are various
//applications of this data structure, such as autocomplete and spellchecker.
//
// Implement the Trie class:
//
//
// Trie() Initializes the trie object.
// void insert(String word) Inserts the string word into the trie.
// boolean search(String word) Returns true if the string word is in the trie (
//i.e., was inserted before), and false otherwise.
// boolean startsWith(String prefix) Returns true if there is a previously
//inserted string word that has the prefix prefix, and false otherwise.
//
//
//
// Example 1:
//
//
//Input
//["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
//[[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
//Output
//[null, null, true, false, true, null, true]
//
//Explanation
//Trie trie = new Trie();
//trie.insert("apple");
//trie.search("apple");   // return True
//trie.search("app");     // return False
//trie.startsWith("app"); // return True
//trie.insert("app");
//trie.search("app");     // return True
//
//
//
// Constraints:
//
//
// 1 <= word.length, prefix.length <= 2000
// word and prefix consist only of lowercase English letters.
// At most 3 * 10⁴ calls in total will be made to insert, search, and
//startsWith.
//
//
// Related Topics 设计 字典树 哈希表 字符串 👍 1760 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
package main

type Trie struct {
	Trie [26]*Trie
	Val  string
}

func Constructor208() Trie {
	return Trie{
		Trie: [26]*Trie{},
	}
}

func (t *Trie) Insert(word string) {
	if len(word) == 0 {
		return
	}
	insert(t, 0, word)
}

func insert(trie *Trie, idx int, word string) {
	if idx == len(word) {
		trie.Val = word
		return
	}

	if trie.Trie[word[idx]-'a'] == nil {
		trie.Trie[word[idx]-'a'] = &Trie{
			Trie: [26]*Trie{},
		}
	}

	insert(trie.Trie[word[idx]-'a'], idx+1, word)
}

func (t *Trie) Search(word string) bool {
	trie := getNode(t, word)

	if trie != nil && trie.Val != "" {
		return true
	} else {
		return false
	}
}

func getNode(trie *Trie, word string) *Trie {
	for i := 0; i < len(word); i++ {
		if trie.Trie[word[i]-'a'] == nil {
			return nil
		} else {
			trie = trie.Trie[word[i]-'a']
		}
	}
	return trie
}

func (t *Trie) StartsWith(prefix string) bool {
	trie := getNode(t, prefix)
	if trie != nil {
		return true
	} else {
		return false
	}
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
//leetcode submit region end(Prohibit modification and deletion)
