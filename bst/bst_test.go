package bst

import (
	"fmt"
	"testing"
)

var bst BST

func fillTree(bst *BST) {
	bst.Insert(8)
	bst.Insert(4)
	bst.Insert(10)
	bst.Insert(2)
	bst.Insert(6)
	bst.Insert(1)
	bst.Insert(3)
	bst.Insert(5)
	bst.Insert(7)
	bst.Insert(9)
}

func TestInsert(t *testing.T) {
	fillTree(&bst)
	bst.String()

	bst.Insert(11)
	bst.String()
}

// isSameSlice returns true if the 2 slices are identical
func isSameSlice(a, b []string) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestInOrderTraverse(t *testing.T) {
	var result []string
	bst.InOrderTraverse(func(i int) {
		result = append(result, fmt.Sprintf("%d", i))
	})
	if !isSameSlice(result, []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11"}) {
		t.Errorf("Traversal order incorrect, got %v", result)
	}
}

func TestPreOrderTraverse(t *testing.T) {
	var result []string
	bst.PreOrderTraverse(func(i int) {
		result = append(result, fmt.Sprintf("%d", i))
	})
	if !isSameSlice(result, []string{"8", "4", "2", "1", "3", "6", "5", "7", "10", "9", "11"}) {
		t.Errorf("Traversal order incorrect, got %v instead of %v", result, []string{"8", "4", "2", "1", "3", "6", "5", "7", "10", "9", "11"})
	}
}

func TestPostOrderTraverse(t *testing.T) {
	var result []string
	bst.PostOrderTraverse(func(i int) {
		result = append(result, fmt.Sprintf("%d", i))
	})
	if !isSameSlice(result, []string{"1", "3", "2", "5", "7", "6", "4", "9", "11", "10", "8"}) {
		t.Errorf("Traversal order incorrect, got %v instead of %v", result, []string{"1", "3", "2", "5", "7", "6", "4", "9", "11", "10", "8"})
	}
}

func TestMin(t *testing.T) {
	v, existed := bst.Min()

	if !existed || v != 1 {
		t.Errorf("min should be 1")
	}
}

func TestMax(t *testing.T) {
	v, existed := bst.Max()
	if !existed || v != 11 {
		t.Errorf("max should be 11")
	}
}

func TestSearch(t *testing.T) {
	if !bst.Search(1) || !bst.Search(8) || !bst.Search(11) {
		t.Errorf("search not working")
	}
}

func TestRemove(t *testing.T) {
	bst.Remove(1)

	v, existed := bst.Min()
	if !existed || v != 2 {
		t.Errorf("min should be 2")
	}
}
