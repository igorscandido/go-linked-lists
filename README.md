
# Go Linked Lists

`go-linked-lists` is a Go library providing a generic implementation of single and double linked lists. It offers a simple interface and functionality to manipulate linked lists with ease, leveraging Go's generics for type safety.

---

## Features

- **Generic Support**: Use any type as the element type for your lists.
- **Single and Double Linked Lists**: Includes both single and double linked list implementations.
- **Rich Operations**:
  - Insertions: Head, Tail, or at a specific index.
  - Removals: Head, Tail, or at a specific index.
  - Access: Get element by index, check existence, find index.
  - Utilities: Iterate, clear, check if empty, get length.
- **Concurrency-Friendly Iteration**: Supports iterating over elements via channels.

---

## Installation

Install the library using `go get`:

```bash
go get github.com/igorscandido/go-linked-lists
```

Import it into your Go project:

```go
import (
    golinkedlists "github.com/igorscandido/go-linked-lists"
    "github.com/igorscandido/go-linked-lists/singlelinked"
    "github.com/igorscandido/go-linked-lists/doublelinked"
)
```

---

## Usage

### Single Linked List

```go
package main

import (
    golinkedlists "github.com/igorscandido/go-linked-lists"
    "github.com/igorscandido/go-linked-lists/singlelinked"
)

func main() {
    var list golinkedlists.LinkedList[int]
    list = singlelinked.NewList[int]()

    list.InsertAtHead(10)
    list.InsertAtTail(20)
    list.InsertAt(15, 1)

    ch := list.Iterate()
    for value := range ch {
        println(value) // Outputs: 10, 15, 20
    }
}
```

### Double Linked List

```go
package main

import (
    golinkedlists "github.com/igorscandido/go-linked-lists"
    "github.com/igorscandido/go-linked-lists/doublelinked"
)

func main() {
    var list golinkedlists.LinkedList[string]
    list = doublelinked.NewList[string]()

    list.InsertAtHead("first")
    list.InsertAtTail("second")
    list.InsertAt("middle", 1)

    value, _ := list.RemoveAt(1)
    println(value) // Outputs: "middle"
}
```

---

## Interface

The library defines a generic interface `LinkedList` with the following methods:

```go
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
```

---

## Error Handling

Common errors include:

- `errCouldNotInsertOutOfIndex`: Attempted to insert at an invalid index.
- `errEmptyList`: Attempted to remove from an empty list.
- `errInvalidIndex`: Provided an invalid index for access or removal.
