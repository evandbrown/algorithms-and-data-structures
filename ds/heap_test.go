package ds

import (
	"testing"
)

func TestEmpty(t *testing.T) {
	h := new(Heap)
	if _, err := h.Pop(); err == nil {
		t.Error("Expected error when trying to pop from empty heap")
	}
}

func TestInsert(t *testing.T) {
	h := new(Heap)
	var err error
	var item int

	h.Insert(400)
	h.Insert(100)
	h.Insert(5)
	h.Insert(500)
	h.Insert(200)
	h.Insert(9)
	item, _ = h.Pop()
	if item != 500 {
		t.Error("Expected 500 to be popped from heap", item)
	}
	item, _ = h.Pop()
	if item != 400 {
		t.Error("Expected 400 to be popped from heap", item)
	}
	item, _ = h.Pop()
	if item != 200 {
		t.Error("Expected 200 to be popped from heap", item)
	}
	item, _ = h.Pop()
	if item != 100 {
		t.Error("Expected 100 to be popped from heap", item, h.data)
	}
	h.Insert(10)
	item, _ = h.Pop()
	if item != 10 {
		t.Error("expected 9 to be popped from heap", item)
	}
	item, _ = h.Pop()
	if item != 9 {
		t.Error("expected 9 to be popped from heap", item)
	}
	h.Insert(10)
	item, _ = h.Pop()
	if item != 10 {
		t.Error("expected 9 to be popped from heap", item)
	}
	item, _ = h.Pop()
	if item != 5 {
		t.Error("expected 5 to be popped from heap", item)
	}
	h.Insert(10)
	item, _ = h.Pop()
	if item != 10 {
		t.Error("expected 9 to be popped from heap", item)
	}
	if _, err = h.Pop(); err == nil {
		t.Error("Expected error when trying to pop from empty heap")
	}
}
