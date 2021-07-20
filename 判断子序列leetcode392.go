//双指针


package main

import "fmt"

func main () {
	bool := isSubsequence("a","bdac")
	fmt.Println(bool)
}

func isSubsequence(s string, t string) bool {
	m,n := len(s), len(t)
	j,k :=0,0
	for j<m && k<n {
		if s[j] != t[k]{
			k++
		} else {
			j++
			k++
		}
	}
	return j==m
}
