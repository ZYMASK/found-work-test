package main

import (
	"fmt"
	"sync"
	"time"
)

var once sync.Once



func test() {
	t := time.NewTicker(20 * time.Second)
	for {
		select {
		case <-t.C:
			fmt.Println(1)

		}
	}
}

func main () {

	test1()
	time.Sleep(time.Second*100)
}


func test1() {
	once.Do(func() {
		go test()
	})

}