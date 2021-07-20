//选择排序 时间O(n^2)  每遍历一遍选择一个最小的放在前面
//插入排序 时间O（n^2)
//冒泡  时间O（n^2） 相邻比较 ，遍历一次比较出个最大的
//归并
//快排   以某个元素为标准，小于此元素的放左边，大于此元素的放右边


package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	test := []int{4,32,100,5,56,1,3,2}
	//res := selectSort(test)
	//res := insertSort(test)
	//res := mergeSort(test)
	res := quickSort(test)
	fmt.Println(res)
	fmt.Println(test)
}
//怎么感觉选择写出成冒泡了
func selectSort (num []int ) []int {
	//res := make([]int,len(num))
	//min := num[0]
	for i := 0;i < len(num)-1;i++ {
		for j := i+1;j < len(num);j++ {
			if num[i] > num[j] {
				num[i] , num[j] = num[j], num[i]
			}
		}
	}
	 return num
}


func insertSort(num []int) []int {
	for i := 1;i< len(num); i++ {
		for j :=i ;j > 0 ;j -- {
			if num[j] < num[j-1]{
				num[j], num[j-1] = num[j-1],num[j]
			}
		}
	}
	return num
}


func mergeSort(num []int) []int {
	length := len(num)
	if length < 2 {
		return num
	} else {
		i := length/2
		left := mergeSort(num[:i])
		right := mergeSort(num[i:])
		res := merge(left, right)
		return res
	}
}

func merge(a,b []int) []int {
	result := make([]int,0)   //注意是0，不能为1
	l,r := 0 ,0
	len1,len2 := len(a), len(b)
	for l < len1 && r < len2 {
		if a[l] < b [r] {
			result = append(result,a[l])
			l++
		} else {
			result = append(result, b[r])
			r++
		}
	}
	//如果有一个数组比完了，另一个数组家在后面
	result = append(result,a[l:]...)
	result = append(result, b[r:]...)
	return result
}



//快排 https://blog.csdn.net/zhizhengguan/article/details/108163606     一个使用多线程排序的方法


func quickSort(num []int) []int{
	var wg sync.WaitGroup
	var length = len(num)
	if length <= 1 {
		return num
	}
	low := make([]int,0)
	mid := make([]int,0)
	high := make([]int,0)
	//获取一个随机数
	rand.Seed(time.Now().UnixNano())
	base := num[rand.Intn(length)]

	for i := 0;i < length; i ++ {
		if num[i] < base {
			low = append(low, num[i])
		} else if num[i] == base {
			mid = append(mid, num[i])
		} else if num[i] > base {
			high = append(high, num[i])
		}
	}
	wg.Add(2)
	go func() {
		low = quickSort(low)
		wg.Done()
	}()
	go func() {
		high = quickSort(high)
		wg.Done()
	} ()
	wg.Wait()
    return append(append(low,mid...),high...)
}