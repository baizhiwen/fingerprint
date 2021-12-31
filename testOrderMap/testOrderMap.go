package main

import (
	"fmt"
	"github.com/elliotchance/orderedmap"
)

func main() {

	m := orderedmap.NewOrderedMap()
	m.Set("aaa", 123)

	fmt.Println(m.Get("aaa"))

}
