package ds

import (
	"fmt"
)

type LinkedList struct {
	data interface{}
	next *LinkedList
}

func (l *LinkedList) String() string {
	var s string
	for ; l != nil; l = l.next {
		s += fmt.Sprintf("%v\n", l.data)
	}
	return s
}
