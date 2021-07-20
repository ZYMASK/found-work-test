package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func sayhello() {
	time.Sleep(time.Second*2)
	fmt.Println("hello")
	wg.Done()
	wg.Done()  //相当于减一
}


func player(name string ,id chan int) {
	defer wg.Done()

	for {
		ball ,ok := <- id
		if !ok {
			fmt.Println("channel is close\n")
			return
		}
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(100)
		if n%10 ==0 {
			close(id)
			fmt.Println( "miss ",name)
			return
		}

		ball ++
		fmt.Println(name, "receive ", ball)
		id <- ball
	}
}



func main() {
	wg.Add(2)  //总共几个任务  2 （相当于计数器）
	//time.Sleep(time.Second*3)
	// go sayhello()
	 	//go sayhello()
	//time.Sleep(time.Second*3)
	id := make(chan int)
	go player("zy",id)
	go player("xyw",id)

	id <- 11
	wg.Wait()

	fmt.Println("结束")
}
