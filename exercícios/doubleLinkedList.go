package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
	info interface{}
	prev *Node
	next *Node
}

type List struct {
	head *Node
	tail *Node
}

func (l *List) Insert(d interface{}) {
	list := &Node{info: d, prev: nil, next: nil}
	if l.head == nil {
		l.head = list
		l.tail = list
	} else {
		p := l.head
		for p.next != nil {
			p = p.next
		}
		list.prev = p
		p.next = list
		l.tail = list
	}
}

func Show(l *List) {
	p := l.head
	for p != nil {
		fmt.Printf("-> %v ", p.info)
		p = p.next
	}
}

func ReverseShow(l *List) {
	r := l.tail
	for r != nil {
		fmt.Printf("-> %v ", r.info)
		r = r.prev
	}
}

func main() {
	sl := List{}
	seed := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(seed)
	for i := 0; i < 10; i++ {
		sl.Insert(rnd.Intn(100))
	}
	Show(&sl)
	fmt.Println("\n----------------------------")
	ReverseShow(&sl)
}
