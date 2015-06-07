package ds

import (
	"errors"
)

type Heap struct {
	data []int
}

func (h *Heap) Pop() (int, error) {
	if len(h.data) == 0 {
		return 0, errors.New("Heap is empty")
	}
	// Get the item to pop
	item := h.data[0]

	// Copy the last item to the root
	h.data[0] = h.data[len(h.data)-1]

	// Remove the last item
	h.data = h.data[:len(h.data)-1]

	// Remove the last item
	h.swapDown(0)
	return item, nil
}

func (h *Heap) Insert(val int) {
	h.data = append(h.data, val)
	i := len(h.data) - 1
	h.swapUp(i)
}

func (h *Heap) swapUp(i int) {
	if p := h.parent(i); p >= 0 {
		if h.data[i] > h.data[p] {
			h.swap(i, p)
			h.swapUp(p)
		}
	}
	return
}

func (h *Heap) swapDown(i int) {
	l := h.left(i)
	r := h.right(i)
	// If there's a right child, check both
	if r != -1 {
		if h.data[i] < h.data[l] {
			if h.data[l] > h.data[r] {
				h.swap(i, l)
				h.swapDown(l)
			} else {
				h.swap(i, r)
				h.swapDown(r)
			}
		}
	} else if l != -1 {
		if h.data[i] < h.data[l] {
			h.swap(i, l)
			h.swapDown(l)
		}
	}
	return
}

func (h *Heap) swap(i int, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *Heap) parent(i int) int {
	return (i+1)/2 - 1
}

func (h *Heap) left(i int) int {
	if li := i*2 + 1; li > len(h.data)-1 {
		return -1
	} else {
		return li
	}
}

func (h *Heap) right(i int) int {
	if li := i*2 + 2; li > len(h.data)-1 {
		return -1
	} else {
		return li
	}
}
