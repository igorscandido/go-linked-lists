package golinkedlists

type LinkedList[T any] interface {
	InsertAtHead(value T)
	InsertAtTail(value T)
	InsertAt(value T, index int) error

	RemoveFromHead() (T, error)
	RemoveFromTail() (T, error)
	RemoveAt(index int) (T, error)

	GetAt(index int) (T, error)
	Find(value T) *int
	Exists(value T) bool

	Iterate() <-chan T
	Length() int
	IsEmpty() bool
	Clear()
}
