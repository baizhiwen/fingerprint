package main

import (
	"fmt"
	"time"
)

//type Flot []float64
func main() {

	//f := new(Flot)
	//a := []float64{1.2}
	//*f = a
	//
	//fmt.Println(a)
	//fmt.Println(f)
	//fmt.Println(*f)

	//num := Inc()
	//fmt.Println(num)
	//t := time.NewTicker(time.Second)
	//a := <-t.C
	//fmt.Println(a)

	//for {
	//	fmt.Println("aaa")
	//	now := time.Now() // 计算下一个零点
	//	next := now.Add(time.Hour * 24)
	//	next = time.Date(next.Year(), next.Month(), next.Day(), 2, 0, 0, 0, next.Location())
	//	t := time.NewTimer(next.Sub(now))
	//	<-t.C
	//	t.Stop()
	//}

	//go time.AfterFunc(time.Second, func() {
	//	fmt.Println("bbbb")
	//})
	//time.Sleep(time.Second)
	//fmt.Println("-------------" )
	ch := make(chan int)
	select {
	case <-ch:
		fmt.Println("hhhhhh")
	case <-time.After(time.Second * 3):
		fmt.Println("cccc")
	}
	//time.Sleep(time.Second)

}

func Incc() *int {
	v := 12

	return &v
}

func Inc() (v int) {
	defer func() { v++ }()
	return 42
}
