package singlelinked

import (
	"testing"
)

func TestNewList(t *testing.T) {
	list := NewList[int]()
	if list == nil {
		t.Fatal("Expected a new list, got nil")
	}
	if list.Length() != 0 {
		t.Errorf("Expected list length to be 0, got %d", list.Length())
	}
}

func TestInsertAtHead(t *testing.T) {
	list := NewList[int]()
	list.InsertAtHead(10)

	if list.Length() != 1 {
		t.Errorf("Expected list length to be 1, got %d", list.Length())
	}

	value, err := list.GetAt(0)
	if err != nil {
		t.Fatal("Unexpected error while getting value:", err)
	}
	if value != 10 {
		t.Errorf("Expected value at index 0 to be 10, got %v", value)
	}
}

func TestInsertAtTail(t *testing.T) {
	list := NewList[int]()
	list.InsertAtTail(20)
	list.InsertAtTail(30)

	if list.Length() != 2 {
		t.Errorf("Expected list length to be 2, got %d", list.Length())
	}

	value, err := list.GetAt(1)
	if err != nil {
		t.Fatal("Unexpected error while getting value:", err)
	}
	if value != 30 {
		t.Errorf("Expected value at index 1 to be 30, got %v", value)
	}
}

func TestInsertAt(t *testing.T) {
	list := NewList[int]()
	list.InsertAtHead(10)
	list.InsertAtTail(30)

	// Insert at the beginning
	if err := list.InsertAt(20, 1); err != nil {
		t.Fatal("Unexpected error while inserting at index 1:", err)
	}

	if list.Length() != 3 {
		t.Errorf("Expected list length to be 3, got %d", list.Length())
	}

	value, err := list.GetAt(1)
	if err != nil {
		t.Fatal("Unexpected error while getting value:", err)
	}
	if value != 20 {
		t.Errorf("Expected value at index 1 to be 20, got %v", value)
	}

	// Insert at out-of-bounds index
	if err := list.InsertAt(40, 5); err != errCouldNotInsertOutOfIndex {
		t.Errorf("Expected error %v, got %v", errCouldNotInsertOutOfIndex, err)
	}
}

func TestRemoveFromHead(t *testing.T) {
	list := NewList[int]()
	list.InsertAtHead(10)
	list.InsertAtHead(20)

	value, err := list.RemoveFromHead()
	if err != nil {
		t.Fatal("Unexpected error while removing from head:", err)
	}

	if value != 20 {
		t.Errorf("Expected value 20, got %v", value)
	}

	if list.Length() != 1 {
		t.Errorf("Expected list length to be 1, got %d", list.Length())
	}
}

func TestRemoveFromTail(t *testing.T) {
	list := NewList[int]()
	list.InsertAtTail(10)
	list.InsertAtTail(20)

	value, err := list.RemoveFromTail()
	if err != nil {
		t.Fatal("Unexpected error while removing from tail:", err)
	}

	if value != 20 {
		t.Errorf("Expected value 20, got %v", value)
	}

	if list.Length() != 1 {
		t.Errorf("Expected list length to be 1, got %d", list.Length())
	}
}

func TestRemoveAt(t *testing.T) {
	list := NewList[int]()
	list.InsertAtTail(10)
	list.InsertAtTail(20)
	list.InsertAtTail(30)

	value, err := list.RemoveAt(1)
	if err != nil {
		t.Fatal("Unexpected error while removing at index 1:", err)
	}

	if value != 20 {
		t.Errorf("Expected value 20, got %v", value)
	}

	if list.Length() != 2 {
		t.Errorf("Expected list length to be 2, got %d", list.Length())
	}

	// Invalid index
	_, err = list.RemoveAt(5)
	if err != errInvalidIndex {
		t.Errorf("Expected error %v, got %v", errInvalidIndex, err)
	}
}

func TestGetAt(t *testing.T) {
	list := NewList[int]()
	list.InsertAtHead(10)
	list.InsertAtTail(20)

	value, err := list.GetAt(1)
	if err != nil {
		t.Fatal("Unexpected error while getting value:", err)
	}
	if value != 20 {
		t.Errorf("Expected value 20 at index 1, got %v", value)
	}

	// Invalid index
	_, err = list.GetAt(5)
	if err != errInvalidIndex {
		t.Errorf("Expected error %v, got %v", errInvalidIndex, err)
	}
}

func TestFind(t *testing.T) {
	list := NewList[int]()
	list.InsertAtTail(10)
	list.InsertAtTail(20)
	list.InsertAtTail(30)

	index := list.Find(20)
	if index == nil || *index != 1 {
		t.Errorf("Expected index 1, got %v", index)
	}

	index = list.Find(40)
	if index != nil {
		t.Errorf("Expected nil, got %v", index)
	}
}

func TestExists(t *testing.T) {
	list := NewList[int]()
	list.InsertAtTail(10)
	list.InsertAtTail(20)

	if !list.Exists(10) {
		t.Errorf("Expected 10 to exist in the list")
	}

	if list.Exists(30) {
		t.Errorf("Expected 30 to not exist in the list")
	}
}

func TestIterate(t *testing.T) {
	list := NewList[int]()
	list.InsertAtTail(10)
	list.InsertAtTail(20)

	ch := list.Iterate()
	var values []int
	for value := range ch {
		values = append(values, value)
	}

	if len(values) != 2 {
		t.Errorf("Expected 2 values, got %d", len(values))
	}

	if values[0] != 10 || values[1] != 20 {
		t.Errorf("Expected values [10, 20], got %v", values)
	}
}

func TestIsEmpty(t *testing.T) {
	list := NewList[int]()

	if !list.IsEmpty() {
		t.Errorf("Expected list to be empty")
	}

	list.InsertAtHead(10)
	if list.IsEmpty() {
		t.Errorf("Expected list to be non-empty")
	}
}

func TestClear(t *testing.T) {
	list := NewList[int]()
	list.InsertAtTail(10)
	list.InsertAtTail(20)

	list.Clear()
	if list.Length() != 0 {
		t.Errorf("Expected list length to be 0 after clear, got %d", list.Length())
	}
}
