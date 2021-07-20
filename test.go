package main

import (
	"fmt"
	_ "fmt"
)

type student struct {
	Name string
	Age  int
}

func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = &stu
	}
	for k, v := range m {
		println(k, "=>", v.Name)
	}
}

/*
func main () {

//pase_student()

	var a map[string]string
	a = make(map[string]string)
	/*a["a"] = "b"
	a["b"] = "a"
	a["zy"] = "xyw"
	for k ,v := range a {
		println(k,v)
	}

fmt.Println(a)
}
*/


func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c
	fmt.Println(x, y, x+y)
}