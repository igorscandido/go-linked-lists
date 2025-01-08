package doublelinked

import (
	"errors"
	"fmt"
)

var (
	errCouldNotInsertOutOfIndex = errors.New("could not insert at provided index, out of bounds")
	errEmptyList                = errors.New("list is empty")
	errInvalidIndex             = errors.New("invalid index")
)

type node[T any] struct {
	value T
	next  *node[T]
	prev  *node[T]
}

type list[T any] struct {
	head   *node[T]
	tail   *node[T]
	length int
}

func NewList[T any]() *list[T] {
	return &list[T]{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

func (l *list[T]) InsertAtHead(value T) {
	node := &node[T]{value: value, next: l.head}
	if l.head != nil {
		l.head.prev = node
	}
	l.head = node

	if l.tail == nil {
		l.tail = node
	}

	l.length++
}

func (l *list[T]) InsertAtTail(value T) {
	node := &node[T]{value: value, next: nil, prev: l.tail}
	if l.tail != nil {
		l.tail.next = node
	}
	l.tail = node

	if l.head == nil {
		l.head = node
	}

	l.length++
}

func (l *list[T]) InsertAt(value T, index int) error {
	if index < 0 || index > l.length {
		return errCouldNotInsertOutOfIndex
	}

	if index == 0 {
		l.InsertAtHead(value)
		return nil
	}

	if index == l.length {
		l.InsertAtTail(value)
		return nil
	}

	aux := l.head
	for i := 0; i < index-1; i++ {
		aux = aux.next
	}

	node := &node[T]{value: value, next: aux.next, prev: aux}
	if aux.next != nil {
		aux.next.prev = node
	}
	aux.next = node
	l.length++
	return nil
}

func (l *list[T]) RemoveFromHead() (T, error) {
	if l.head == nil {
		var zero T
		return zero, errEmptyList
	}

	value := l.head.value
	l.head = l.head.next
	if l.head != nil {
		l.head.prev = nil
	} else {
		l.tail = nil
	}
	l.length--
	return value, nil
}

func (l *list[T]) RemoveFromTail() (T, error) {
	if l.tail == nil {
		var zero T
		return zero, errEmptyList
	}

	value := l.tail.value
	l.tail = l.tail.prev
	if l.tail != nil {
		l.tail.next = nil
	} else {
		l.head = nil
	}
	l.length--
	return value, nil
}

func (l *list[T]) RemoveAt(index int) (T, error) {
	if index < 0 || index >= l.length {
		var zero T
		return zero, errInvalidIndex
	}

	if index == 0 {
		return l.RemoveFromHead()
	}

	if index == l.length-1 {
		return l.RemoveFromTail()
	}

	aux := l.head
	for i := 0; i < index; i++ {
		aux = aux.next
	}

	value := aux.value
	aux.prev.next = aux.next
	if aux.next != nil {
		aux.next.prev = aux.prev
	}
	l.length--
	return value, nil
}

func (l *list[T]) GetAt(index int) (T, error) {
	if index < 0 || index >= l.length {
		var zero T
		return zero, errInvalidIndex
	}

	aux := l.head
	for i := 0; i < index; i++ {
		aux = aux.next
	}

	return aux.value, nil
}

func (l *list[T]) Find(value T) *int {
	aux := l.head
	index := 0

	for aux != nil {
		if fmt.Sprintf("%v", aux.value) == fmt.Sprintf("%v", value) {
			return &index
		}
		aux = aux.next
		index++
	}

	return nil
}

func (l *list[T]) Exists(value T) bool {
	return l.Find(value) != nil
}

func (l *list[T]) Iterate() <-chan T {
	ch := make(chan T)
	go func() {
		aux := l.head
		for aux != nil {
			ch <- aux.value
			aux = aux.next
		}
		close(ch)
	}()
	return ch
}

func (l *list[T]) Length() int {
	return l.length
}

func (l *list[T]) IsEmpty() bool {
	return l.length == 0
}

func (l *list[T]) Clear() {
	l.head = nil
	l.tail = nil
	l.length = 0
}
