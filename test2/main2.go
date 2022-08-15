package main

import (
	"flag"
	"fmt"
	"os"
)

var name string

func init() {
	flag.CommandLine = flag.NewFlagSet("", flag.ExitOnError)
	flag.CommandLine.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
		flag.PrintDefaults()
	}
	flag.StringVar(&name, "nameA", "jiaen", "my baobao")
}

func main() {
	//flag.Usage = func() {
	//	fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
	//	flag.PrintDefaults()
	//}
	flag.Parse()
	fmt.Printf("Hello, %s!\n", name)
	println("aaaaa mytest")

}

//参考答案及解析：不会出现死循环，能正常结束。
//
//循环次数在循环开始前就已经确定，循环内改变切片的长度，不影响循环次数。
func main2() {
	v := []int{1, 2, 3}
	for i := range v {
		v = append(v, i)
	}
}