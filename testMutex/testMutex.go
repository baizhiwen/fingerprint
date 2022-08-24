package main

import (
	"fmt"
	"sync"
)

//var total struct {
//	sync.Mutex
//	value int
//}
//
//func worker(wg *sync.WaitGroup) {
//	go func() {
//		defer wg.Done()
//		for i := 0; i <= 100; i++ {
//			total.Lock()
//			total.value += i
//			total.Unlock()
//		}
//	}()
//}
//
//func main() {
//	var wg sync.WaitGroup
//	wg.Add(2)
//	worker(&wg)
//	worker(&wg)
//	wg.Wait()
//
//	fmt.Println(total.value)
//}

//
//func main() {
//	// 互斥锁保护计数器
//	var mu sync.Mutex
//	// 计数器的值
//	var count int64
//
//	// 辅助变量，用来确认所有的goroutine都完成
//	var wg sync.WaitGroup
//	wg.Add(10)
//
//	// 启动10个gourontine
//	for i := 0; i < 10; i++ {
//		go func() {
//			defer wg.Done()
//			// 累加10万次
//			for j := 0; j < 100000; j++ {
//				mu.Lock()
//				count++
//				mu.Unlock()
//				//atomic.AddInt64(&count,1)
//			}
//		}()
//	}
//	wg.Wait()
//	fmt.Println(count)
//}

type Counter struct {
	sync.Mutex
	Count int
}

func main() {
	var c Counter
	c.Lock()
	defer c.Unlock()
	c.Count++
	foo(c) // 复制锁
}

// 这里Counter的参数是通过复制的方式传入的
func foo(c Counter) {
	c.Lock()
	defer c.Unlock()
	fmt.Println("in foo")
}
