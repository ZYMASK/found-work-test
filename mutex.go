package main

import (
	"fmt"
	"sync"
)

var myx sync.Mutex
var num int32
var wg1 sync.WaitGroup
var ch =  make(chan int32,1)

func addnum(num *int32) {
	defer wg1.Done()
	for i :=0;i<100000;i++ {
		*num++
	}
}

func addnum1() {
	defer wg1.Done()
	for i :=0;i<100000;i++ {
		num++
	}
}


func addchannel() {
	defer wg1.Done()
	for i:= 0;i<100000;i++ {
		num = <- ch
		num++
		ch <- num
	}
}

func main() {
	wg1.Add(2)


go addchannel()
	go addchannel()
ch <- 0

	wg1.Wait()

	fmt.Println(num)


}
