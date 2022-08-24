package main

import (
	"encoding/json"
	"fmt"
	"github.com/elliotchance/orderedmap"
	"time"
)

func main1() {

	m := orderedmap.NewOrderedMap()
	m.Set("aaa", 123)

	fmt.Println(m.Get("aaa"))

}
func hello(num [3]int) {
	num[0] = 18
}
func hello1(num []int) {
	num[0] = 18
}

func main2() {
	var i [3]int
	i[0] = 5
	i[1] = 6
	i[2] = 7
	hello(i)
	fmt.Println(i) //[5 6 7]

	a := make([]int, 0)
	a = append(a, 5)
	a = append(a, 6)
	a = append(a, 7)
	hello1(a)
	fmt.Println(a) //[18 6 7]
}

//==============================
func hello2(num ...int) {
	num[0] = 18
}

//func main() {
//	i := []int{5, 6, 7}
//	hello2(i...)
//	fmt.Println(i[0])
//}

type AutoGenerated struct {
	Age   int    `json:"age"`
	Name  string `json:"name"`
	Child []int  `json:"child"`
}

func main() {
	t := struct {
		N int
		time.Time
	}{
		5,
		time.Date(2020, 12, 20, 0, 0, 0, 0, time.UTC),
	}

	m, _ := json.Marshal(t)
	fmt.Printf("%s", m)
}
