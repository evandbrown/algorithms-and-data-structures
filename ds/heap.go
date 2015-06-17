package ds

import (
	"bytes"
	"errors"
	"fmt"
	"math"
	"strconv"
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

func (h *Heap) Peek() (int, error) {
	if len(h.data) == 0 {
		return 0, errors.New("Heap is empty")
	}

	return h.data[0], nil
}

func (h *Heap) Insert(val int) {
	h.data = append(h.data, val)
	i := len(h.data) - 1
	h.swapUp(i)
}

func (h *Heap) String() string {
	if len(h.data) == 0 {
		return ""
	} else {
		l := h.levelMap()
		var b bytes.Buffer
		b.WriteString("\n")
		maxval, _ := h.Peek()
		maxchar := len(strconv.Itoa(maxval)) + 2
		maxlevel := h.maxLevel()
		maxrowlen := int(float64(maxchar) * math.Exp2(float64(maxlevel)))
		for i := 0; i <= maxlevel; i++ {
			for _, v := range l[i] {
				pad := (maxrowlen - maxchar*int(math.Exp2(float64(i)))) / int(math.Exp2(float64(i))) / 2
				fmt.Printf("Padding is %v\n", pad)
				b.WriteString(fmt.Sprintf("%"+strconv.Itoa(pad)+"v", " "))
				b.WriteString(fmt.Sprintf("%"+strconv.Itoa(maxchar)+"s", strconv.Itoa(v)))
				b.WriteString(fmt.Sprintf("%"+strconv.Itoa(pad)+"v", " "))
			}
			b.WriteString("\n")
		}
		return b.String()
	}
}

func (h *Heap) levelMap() map[int][]int {
	// Map of int -> []int
	l := make(map[int][]int)

	for i, v := range h.data {
		l[h.level(i)] = append(l[h.level(i)], v)
	}

	return l
}

func (h *Heap) maxLevel() int {
	return h.level(len(h.data) - 1)
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

func (h *Heap) level(i int) int {
	if i == 0 {
		return 0
	}

	return int(math.Log(float64(i+1)) / math.Log(2))
}
