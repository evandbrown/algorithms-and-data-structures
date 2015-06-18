package ds

import (
	"fmt"
)

func Reverse(l *LinkedList) *LinkedList {
	if l.next == nil {
		return l
	}
	newHead := l.next
	l.next = nil
	reversed := Reverse(newHead)
	newHead.next = l
	return reversed
}

type LinkedList struct {
	data interface{}
	next *LinkedList
}

func (l *LinkedList) String() string {
	var s string
	for ; l != nil; l = l.next {
		s += fmt.Sprintf("%v", l.data)
		if l.next != nil {
			s += fmt.Sprintf("-")
		}
	}
	return s
}
