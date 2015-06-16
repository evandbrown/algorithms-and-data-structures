package ds

import (
	"testing"
)

func TestLinkedList(t *testing.T) {
	n1 := &LinkedList{data: 1}
	n1.next = &LinkedList{data: 2}
	n1.next.next = &LinkedList{data: 3}
	n1.next.next.next = &LinkedList{data: 4}
	n1.next.next.next.next = &LinkedList{data: 5}
	t.Log(n1)
}
