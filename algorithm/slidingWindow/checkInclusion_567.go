package main

import "fmt"

func main() {
	s1 := "ab"
	s2 := "eidbaooo"
	ans := checkInclusion(s1,s2)
	fmt.Println(ans)
}

func checkInclusion(s1 string, s2 string) bool {
	ans := false
	plen,slen := len(s1), len(s2)
	var scount,pcount [26]int
	if plen > slen {
		return false
	}
	for k,v :=range s1 {
		pcount[v-'a']++
		scount[s2[k]-'a']++
	}
	if pcount == scount {
		return true
	}
	for k,v :=range s2[:slen-plen] {
		scount[v-'a']--
		scount[s2[k+plen]-'a']++
		if pcount == scount {
			return true
		}
	}
	return ans
}