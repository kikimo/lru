package main

import "fmt"

type List struct {
	size int
	head Element // sentinel
	tail Element // sentinel
}

type Element struct {
	next  *Element
	prev  *Element
	Value interface{}
}

func (l *List) PushFront(v interface{}) *Element {
	e := &Element{Value: v}
	l.doPushFront(e)
	l.size++

	return e
}

func (l *List) doPushFront(e *Element) {
	e.prev, e.next = &l.head, l.head.next
	e.prev.next = e
	e.next.prev = e
}

func (l *List) MoveToFront(e *Element) {
	l.Remove(e)
	l.doPushFront(e)
}

func (l *List) Back() *Element {
	if l.size == 0 {
		return nil
	}

	e := l.tail.prev
	return e
}

func (l *List) Remove(e *Element) {
	prev, next := e.prev, e.next
	prev.next = next
	next.prev = prev
	e.next = nil
	e.prev = nil
}

func NewList() *List {
	l := &List{}
	l.head.next = &l.tail
	l.tail.prev = &l.head

	return l
}

func (l *List) Print() {
	for ptr := l.head.next; ptr != &l.tail; ptr = ptr.next {
		val := ptr.Value.(int)
		fmt.Printf("%d ", val)
	}

	fmt.Println()
}

func main() {
	l := NewList()
	e1 := l.PushFront(1)
	e2 := l.PushFront(2)
	l.PushFront(3)
	l.PushFront(4)
	l.PushFront(5)
	l.Print()

	l.Remove(e1)
	l.Print()

	l.MoveToFront(e2)
	l.Print()
}
