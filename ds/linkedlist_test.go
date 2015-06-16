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
	if n1.data != 1 {
		t.Error("Expected 1 to be the first element in list %v\n", n1)
	}
}

func TestLinkedListReverse(t *testing.T) {
	n1 := &LinkedList{data: 1}
	n1.next = &LinkedList{data: 2}
	n1.next.next = &LinkedList{data: 3}
	n1.next.next.next = &LinkedList{data: 4}
	n1.next.next.next.next = &LinkedList{data: 5}
	n1 = Reverse(n1)
	t.Log(n1)
	if n1.data != 5 {
		t.Error("Expected 5 to be the first element in list %v\n", n1)
	}
}
