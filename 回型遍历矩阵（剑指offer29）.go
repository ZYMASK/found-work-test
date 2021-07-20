
//按层遍历 最外面第一层



package main



func spiralOrder(order [][]int) []int {
	var output []int
	output = make([]int,(len(order)-1)*len(order[0])-1)
	index0 := 0
	left , right , top ,bottom:=0 ,len(order)-1,0, len(order[0])-1
	for left <= right && top <= bottom {
		for index := left ; index<= right; index++ {
			output[index0] = order[left][index]
			index0++
		}
		for index := top  ; index <= bottom ;index++ {
			output[index0] = order[right][index]
			index0++
		}
		for index := right ; index >= left; index--	{
			output[index0] = order[index][bottom]
			index0++
		}
		for index := bottom; index >= top; index-- {
			output[index0]= order[left][index]
			index0++
		}
		left++
		right--
		top++
		bottom--
		}
}