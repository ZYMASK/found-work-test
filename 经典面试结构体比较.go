//不能比较的 map slice func
//struct内含有以上三个的，比较时编译出错


package main

import "fmt"

type User struct {
	age  int
	name string
}

func main() {
	user := User{
		1,
		"d",
	}
	user2 := User{
		1,
		"d",
	}
	fmt.Println(user == user2)  //打印的结果是true 会去自动对比内部的属性是否相等
	//但是如果结构体内部含有map,slice,Function 使用==比较编译会报错
}



//当含有不可比较的数据时，可以使用reflect包的reflect.DeepEqual()函数
/*
reflect.DeepEqual
DeepEqual函数用来判断两个值是否深度一致。具体比较规则如下：

不同类型的值永远不会深度相等
当两个数组的元素对应深度相等时，两个数组深度相等
当两个相同结构体的所有字段对应深度相等的时候，两个结构体深度相等
当两个函数都为nil时，两个函数深度相等，其他情况不相等（相同函数也不相等）
当两个interface的真实值深度相等时，两个interface深度相等
map的比较需要同时满足以下几个
两个map都为nil或者都不为nil，并且长度要相等
相同的map对象或者所有key要对应相同
map对应的value也要深度相等
指针，满足以下其一即是深度相等
两个指针满足go的==操作符
两个指针指向的值是深度相等的
切片，需要同时满足以下几点才是深度相等
两个切片都为nil或者都不为nil，并且长度要相等
两个切片底层数据指向的第一个位置要相同或者底层的元素要深度相等
注意：空的切片跟nil切片是不深度相等的
其他类型的值（numbers, bools, strings, channels）如果满足go的==操作符，则是深度相等的。要注意不是所有的值都深度相等于自己，例如函数，以及嵌套包含这些值的结构体，数组等
*/

//两个不同的结构体是否可以比较
//可通过强制转换来比较
//内有不可比较内容时编译报错
type student1 struct {
	name string
	age  int
	class int
}
type teacher struct {
	name string
	age  int
	class int
}
func equal(s student1, t teacher) {
	t1 := teacher(s)
	fmt.Println(t1==t)
}
