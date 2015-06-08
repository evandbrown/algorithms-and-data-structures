package sort

import (
	"math/rand"
	"testing"
	"time"
)

func TestInsertSort(t *testing.T) {
	items := getList(5000, 5000)
	items[100] = 0
	items = InsertSort(items)
	if items[0] != 0 {
		t.Error("Expected 0 as first value", items[0])
	}
}

func TestInsertSortNoSwap(t *testing.T) {
	items := getList(50000, 50000)
	items[100] = 0
	items = InsertSortNoSwap(items)
	if items[0] != 0 {
		t.Error("Expected 0 as first value", items[0])
	}
}

func TestBubbleSort(t *testing.T) {
	items := getList(5000, 5000)
	items[100] = 0
	items = BubbleSort(items)
	if items[0] != 0 {
		t.Error("Expected 0 as first value", items[0])
	}
}

func TestQuickSort(t *testing.T) {
	items := getList(50000000, 5000000000)
	items[40] = 0
	items = QuickSort(items)
	t.Logf("%v...%v\n", items[:10], items[len(items)-10:])
	if items[0] != 0 {
		t.Error("Expected 0 as first value", items[0])
	}
}

func TestBinFind(t *testing.T) {
	items := getList(5000, 5000)
	items[100] = 4997
	items = InsertSortNoSwap(items)
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
