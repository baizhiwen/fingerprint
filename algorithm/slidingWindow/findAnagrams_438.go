package main

import "fmt"

func main() {
	s := "cbaebabacd"
	p := "abc"
	ans := findAnagrams(s,p)
	fmt.Println(ans)
}

func findAnagrams(s, p string) ( []int) {
	ans := make([]int,0)
	slen,plen := len(s), len(p)
	if slen < plen {
		return ans
	}
	var pcount,scount [26]int
	for k,v :=range p {
		pcount[v-'a']++
		scount[s[k]-'a']++
	}
	if pcount == scount {
		ans = append(ans, 0)
	}
	for k,v := range s[:slen-plen] {
		scount[v-'a']--
		scount[s[k+plen]-'a']++
		if pcount == scount {
			ans = append(ans, k+1)
		}
	}
	return ans
}