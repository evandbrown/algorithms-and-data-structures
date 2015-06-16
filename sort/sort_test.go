package sort

import (
	"math/rand"
	"testing"
	"time"
)

func TestInsertSort(t *testing.T) {
	items := getList(50000, 50000)
	items[100] = 0
	items = InsertSortNew(items)
	if items[0] != 0 {
		t.Error("Expected 0 as first value", items[0])
	}
}

func TestSelectionSort(t *testing.T) {
	items := getList(50000, 50000)
	items[100] = 0
	items = SelectionSort(items)
	if items[0] != 0 {
		t.Error("Expected 0 as first value", items[0])
	}
}

func TestBubbleSort(t *testing.T) {
	items := getList(50000, 50000)
	items[100] = 0
	items = BubbleSort(items)
	if items[0] != 0 {
		t.Error("Expected 0 as first value", items[0])
	}
}

func TestQuickSort(t *testing.T) {
	items := getList(50000, 50000)
	items[40] = 0
	items = QuickSort(items)
	if items[0] != 0 {
		t.Error("Expected 0 as first value", items[0])
	}
}

func TestBinFind(t *testing.T) {
	items := getList(500000, 500000)
	items[100] = 4997
	items = QuickSort(items)
	i, err := BinFind(4997, items)
	t.Logf("Find returned index %v\n", i)
	if err != nil {
		t.Error("Expected to find 4997 and for error to be nil")
	}
}

func getList(length int, max int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	list := make([]int, length)
	for i := 0; i < length; i++ {
		list[i] = r.Intn(max)
	}
	return list
}
