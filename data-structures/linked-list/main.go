package main

import (
	"fmt"
	"sync"
)

type node struct {
	data interface{}
	next *node
}

type linkedList struct {
	head  *node
	count int
	mu    sync.Mutex // thread safe for concurrent access
}

func (l *linkedList) prepend(n *node) {
	second := l.head
	l.mu.Lock()
	defer l.mu.Unlock()
	l.head = n
	n.next = second
	l.count++
}

func (l linkedList) printListData() {
	toPrint := l.head
	for ; l.count != 0; l.count-- {
		fmt.Printf("%v ", toPrint.data)
		toPrint = toPrint.next
	}
	fmt.Printf("\n")
}

func (l *linkedList) deleteWithValue(value interface{}) {
	if l.count == 0 {
		return
	}

	if l.head.data == value {
		l.mu.Lock()
		defer l.mu.Unlock()
		l.head = l.head.next
		l.count--
		return
	}

	previousToDelete := l.head
	for previousToDelete.next.data != value {
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	previousToDelete.next = previousToDelete.next.next
	l.count--
}

func main() {
	mylist := linkedList{}
	node1 := &node{data: 100}
	node2 := &node{data: 200}
	node3 := &node{data: 300}
	node4 := &node{data: 400}
	node5 := &node{data: 500}

	mylist.prepend(node1)
	mylist.prepend(node2)
	mylist.prepend(node3)
	mylist.prepend(node4)
	mylist.prepend(node5)
	mylist.printListData()
	mylist.deleteWithValue(300)
	mylist.deleteWithValue(0)
	mylist.deleteWithValue(100)
	mylist.deleteWithValue(200)
	mylist.deleteWithValue(400)
	mylist.deleteWithValue(500)
	mylist.printListData()
}
