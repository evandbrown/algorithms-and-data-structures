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

func TestMaxLevel(t *testing.T) {
	h := new(Heap)

	h.Insert(400)
	if h.MaxLevel() != 0 {
		t.Error("Expected MaxLevel of heap to be 0", h.MaxLevel())
	}
	h.Insert(100)
	if h.MaxLevel() != 1 {
		t.Error("Expected MaxLevel of heap to be 1", h.MaxLevel())
	}
	h.Insert(5)
	if h.MaxLevel() != 1 {
		t.Error("Expected MaxLevel of heap to be 1", h.MaxLevel())
	}
	h.Insert(500)
	h.Insert(200)
	h.Insert(9)

	if h.MaxLevel() != 2 {
		t.Error("Expected MaxLevel of heap to be 2", h.MaxLevel())
	}

	h.Insert(400)
	h.Insert(100)
	if h.MaxLevel() != 3 {
		t.Error("Expected MaxLevel of heap to be 3", h.MaxLevel())
	}
}
