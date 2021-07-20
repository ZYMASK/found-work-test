package main

import (
	"fmt"
	_ "strconv"
)

func main() {
	test := make([]int ,2)
	test2 := []int{2,1,8}
	fmt.Printf("test=%p,test2=%p\n",&test,&test2)
	//test = change(test)
	change(test)
	fmt.Println(test)
}


func change(test []int)  {
	test = append(test,181921)
	fmt.Printf("%p",&test)
	fmt.Println("chage内的test",test)
	//return test
}

//传参进去，值拷贝（地址不一样），不影响原来的值,换个写法，传指针可以改