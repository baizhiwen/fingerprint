package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

func main() {
	f, err := os.Create("./testTrace/trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	fmt.Println("hhhhh")
	trace.Stop()

	//go tool trace ./testTrace/trace.out
}
