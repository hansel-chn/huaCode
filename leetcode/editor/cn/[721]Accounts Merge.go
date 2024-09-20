//Given a list of accounts where each element accounts[i] is a list of strings,
//where the first element accounts[i][0] is a name, and the rest of the elements
//are emails representing emails of the account.
//
// Now, we would like to merge these accounts. Two accounts definitely belong
//to the same person if there is some common email to both accounts. Note that even
//if two accounts have the same name, they may belong to different people as
//people could have the same name. A person can have any number of accounts initially,
//but all of their accounts definitely have the same name.
//
// After merging the accounts, return the accounts in the following format: the
//first element of each account is the name, and the rest of the elements are
//emails in sorted order. The accounts themselves can be returned in any order.
//
//
// Example 1:
//
//
//Input: accounts = [["John","johnsmith@mail.com","john_newyork@mail.com"],[
//"John","johnsmith@mail.com","john00@mail.com"],["Mary","mary@mail.com"],["John",
//"johnnybravo@mail.com"]]
//Output: [["John","john00@mail.com","john_newyork@mail.com","johnsmith@mail.
//com"],["Mary","mary@mail.com"],["John","johnnybravo@mail.com"]]
//Explanation:
//The first and second John's are the same person as they have the common email
//"johnsmith@mail.com".
//The third John and Mary are different people as none of their email addresses
//are used by other accounts.
//We could return these lists in any order, for example the answer [['Mary',
//'mary@mail.com'], ['John', 'johnnybravo@mail.com'],
//['John', 'john00@mail.com', 'john_newyork@mail.com', 'johnsmith@mail.com']]
//would still be accepted.
//
//
// Example 2:
//
//
//Input: accounts = [["Gabe","Gabe0@m.co","Gabe3@m.co","Gabe1@m.co"],["Kevin",
//"Kevin3@m.co","Kevin5@m.co","Kevin0@m.co"],["Ethan","Ethan5@m.co","Ethan4@m.co",
//"Ethan0@m.co"],["Hanzo","Hanzo3@m.co","Hanzo1@m.co","Hanzo0@m.co"],["Fern","Fern5@
//m.co","Fern1@m.co","Fern0@m.co"]]
//Output: [["Ethan","Ethan0@m.co","Ethan4@m.co","Ethan5@m.co"],["Gabe","Gabe0@m.
//co","Gabe1@m.co","Gabe3@m.co"],["Hanzo","Hanzo0@m.co","Hanzo1@m.co","Hanzo3@m.
//co"],["Kevin","Kevin0@m.co","Kevin3@m.co","Kevin5@m.co"],["Fern","Fern0@m.co",
//"Fern1@m.co","Fern5@m.co"]]
//
//
//
// Constraints:
//
//
// 1 <= accounts.length <= 1000
// 2 <= accounts[i].length <= 10
// 1 <= accounts[i][j].length <= 30
// accounts[i][0] consists of English letters.
// accounts[i][j] (for j > 0) is a valid email.
//
//
// Related Topics 深度优先搜索 广度优先搜索 并查集 数组 哈希表 字符串 排序 👍 533 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
package main

import "sort"

func accountsMerge(accounts [][]string) [][]string {
	email2name := make(map[string]string)
	p := make(map[string]string)
	uf := NewUF721(p)

	for _, account := range accounts {
		uName := account[0]
		for _, email := range account[1:] {
			if _, ok := email2name[email]; !ok {
				email2name[email] = uName
			}
			if _, ok := p[email]; !ok {
				p[email] = email
			}

			uf.Connect(account[1], email)
		}
	}

	afterUnion := make(map[string][]string)
	for email, _ := range email2name {
		parent := uf.Find(email)
		if _, ok := afterUnion[parent]; !ok {
			afterUnion[parent] = []string{email}
		} else {
			afterUnion[parent] = append(afterUnion[parent], email)
		}
	}

	rlt := make([][]string, 0)
	for _, emails := range afterUnion {
		sort.Strings(emails)
		userEmail := append([]string{email2name[emails[0]]}, emails...)
		rlt = append(rlt, userEmail)
	}
	return rlt
}

func NewUF721(p map[string]string) UF721 {
	return UF721{p}
}

type UF721 struct {
	Parent map[string]string
}

func (u *UF721) Find(i string) string {
	if i != u.Parent[i] {
		p := u.Find(u.Parent[i])
		u.Parent[i] = p
	}
	return u.Parent[i]
}

func (u *UF721) Connect(i, j string) {
	if u.Find(i) != u.Find(j) {
		u.Parent[u.Find(i)] = j
	}
}

//func accountsMerge(accounts [][]string) [][]string {
//	email2ac := make(map[string][]int)
//	for i, account := range accounts {
//		for _, email := range account[1:] {
//			if _, ok := email2ac[email]; ok {
//				email2ac[email] = append(email2ac[email], i)
//			} else {
//				email2ac[email] = []int{i}
//			}
//		}
//	}
//
//	rlt := make([][]string, 0)
//	visitedEmail := make(map[string]struct{})
//
//	var traverse func(email string, userEmail *[]string)
//	traverse = func(email string, userEmail *[]string) {
//		if _, ok := visitedEmail[email]; ok {
//			return
//		}
//
//		visitedEmail[email] = struct{}{}
//		*userEmail = append(*userEmail, email)
//
//		acs := email2ac[email]
//		for _, ac := range acs {
//			for _, nxtEmail := range accounts[ac][1:] {
//				traverse(nxtEmail, userEmail)
//			}
//		}
//	}
//
//	for email, acs := range email2ac {
//		userEmail := make([]string, 0)
//		traverse(email, &userEmail)
//		sort.Strings(userEmail)
//		if len(userEmail) != 0 {
//			userEmail = append([]string{accounts[acs[0]][0]}, userEmail...)
//			rlt = append(rlt, userEmail)
//		}
//	}
//	return rlt
//}

//leetcode submit region end(Prohibit modification and deletion)
