//给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

//输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
//输出：6
//解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。

//公式 f(i)= max{f(i-1)+num[i],num[i]}    遍历一遍即可

package main

import "fmt"

func main() {
	nums := make([]int,5)
	nums = []int{1,-2,3,4,5}
	v := maxSum(nums)
	fmt.Println(v)
}

func maxSum(nums []int)int {
	max := nums[0]
	for i := 1;i<len(nums);i++ {
		if nums[i] + nums[i-1] > nums[i] {                        //num[i]存到当前的最大子序和
			nums[i] = nums[i-1] + nums[i]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}