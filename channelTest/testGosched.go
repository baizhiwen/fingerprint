package main

import (
	"fmt"
	"runtime"
	"sync"
)

const N = 1

func main() {
	var wg sync.WaitGroup

	wg.Add(N)
	for i := 0; i < N; i++ {
		go start(&wg)
		fmt.Println(i, "aaaaa")
	}
	fmt.Println("aa")

	wg.Wait()
}

func start(wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		runtime.Gosched()
		fmt.Println(i)
	}
	wg.Done()
}
