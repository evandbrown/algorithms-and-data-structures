package sort

import (
	"errors"
)

func InsertSort(items []int) []int {
	for i := 1; i < len(items); i++ {
		t, j := items[i], 0
		for j = i - 1; j >= 0 && t < items[j]; j-- {
			items[j+1] = items[j]
		}
		items[j+1] = t
	}
	return items
}

func InsertSortNew(items []int) []int {
	for i := 1; i < len(items); i++ {
		t, j := items[i], i-1
		for ; j >= 0 && t < items[j]; j-- {
			items[j+1] = items[j]
		}
		items[j+1] = t
	}
	return items
}

func SelectionSort(items []int) []int {
	for i := 0; i < len(items)-1; i++ {
		li := i
		for j := i + 1; j < len(items); j++ {
			if items[j] < items[li] {
				li = j
			}
		}
		items[i], items[li] = items[li], items[i]
	}
	return items
}

func BubbleSort(items []int) []int {
	f := false

	for i := 0; i < len(items)-1; i++ {
		if items[i] > items[i+1] {
			items[i], items[i+1] = items[i+1], items[i]
			f = true
		}
		if i == len(items)-2 && f == true {
			i = -1
			f = false
		}
	}
	return items
}

func QuickSort(items []int) []int {
	if len(items) <= 1 {
		return items
	}
	pi := len(items) - 1
	for i := 0; i < pi; i++ {
		if items[i] > items[pi] {
			items[pi-1], items[i] = items[i], items[pi-1]
			items[pi-1], items[pi] = items[pi], items[pi-1]
			pi--
			i--
		}
	}
	QuickSort(items[:pi])
	QuickSort(items[pi:])
	return items
}

func BinFind(key int, items []int) (int, error) {
	low, high := 0, len(items)-1
	var mid int
	for low <= high {
		mid = low + (high-low)/2
		if key == items[mid] {
			return mid, nil
		} else if key < items[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1, errors.New("Item not found")
}
