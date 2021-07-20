//遍历一遍，找最小值




package main

import (
	"fmt"
)

func main() {
	a := []int{7,6,5,4,9}
	m := maxProfit2(a)
	fmt.Println(m)

}

//买卖一次
func maxProfit(price []int) int {
	minPrice := price[0]
	maxProfit := 0

	for i := 0;i< len(price);i++ {
		if i > 0 && price[i] < minPrice {
			minPrice = price[i]
		}
		if maxProfit < price[i]- minPrice {
			maxProfit = price[i]- minPrice
		}
	}
	if maxProfit <=0 {
		return  0
	} else {
		return  maxProfit
	}
}

//买卖多次  贪心  只收集正利润
func maxProfit2(price []int ) int {
	maxPrice := 0
	for i:= 1;i < len(price);i++ {
		maxPrice += Max(price[i]-price[i-1],0)
	}
	return maxPrice
}


func Max(a,b int ) int {
	if a > b {
		return  a
	} else {
		return b
	}
}