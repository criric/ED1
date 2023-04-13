package main

import (
	"fmt"

	"github.com/eduardolfalcao/edi/src/go/list"
)

func main() {
	//var linkedlist list.ArrayList
	linkedlist := list.ArrayList{}
	limit := 100
	linkedlist.Init(limit)

	for i := 0; i < limit; i++ {
		linkedlist.Add(i)
	}
	for i := 0; i < limit; i++ {
		val, err := linkedlist.Get(i)
		if err == nil {
			fmt.Println(val)
		}
	}
	for i := 0; i < limit; i++ {
		if i%2 == 0 {
			linkedlist.Set(i*2, i)
		}
	}
	fmt.Println("######################")
	for i := 0; i < limit; i++ {
		val, err := linkedlist.Get(i)
		if err == nil {
			fmt.Println(val)
		}
	}
}
