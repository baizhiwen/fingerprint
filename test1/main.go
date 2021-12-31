package main

import (
	"fmt"
	"time"
)

func main() {
	for n := range gen() {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}

	time.Sleep(time.Second * 20)
}

func gen() <-chan int {
	ch := make(chan int)
	go func() {
		var n int
		for {
			fmt.Println("in ", n)
			ch <- n
			n++
			fmt.Println("aa ", n)
			//time.Sleep(time.Second)
		}
		for {
			fmt.Println(n)
			//time.Sleep(time.Second)
		}
	}()
	return ch
}

func worker(taskCh <-chan int) {
	const N = 5
	// 启动 5 个工作协程
	//for i := 0; i < N; i++ {
	go func() {
		for {
			task := <-taskCh
			fmt.Printf("finish task: %d by worker %d\n", task, 1)
			time.Sleep(time.Second)
		}
	}()
	//}
}
