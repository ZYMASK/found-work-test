//一次爬一格或者两格 公式f(n) = f(n-1) + f(n-2)


package main

import "fmt"

func main() {
	m:=climbStairs2(4)
    n:=climbStairs(4)
    fmt.Println(m,n)
}

func climbStairs(n int ) int {
	j, k, l := 0, 0, 1
	for i := 1; i <= n; i++ {
		j = k
		k = l
		l = j + k
	}
	return l
}

func climbStairs2(n int ) int {
	if n <= 2 {
		return n
	} else {
		m:= climbStairs2(n-1)+ climbStairs2(n-2)
		return m
	}
}