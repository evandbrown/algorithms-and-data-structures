package sort

import (
	"fmt"
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
	items := getList(500, 500)
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
	items := getList(150, 150)
	toFind := -1
	items[100] = toFind
	items = QuickSort(items)
	i, err := BinFind(toFind, items)
	t.Logf("Find returned index %v and error %v\n", i, err)
	if err != nil {
		t.Error(fmt.Sprintf("Expected to find %v and for error to be nil", toFind))
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
