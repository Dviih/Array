# Array
### Since there's no `sync.Array` here you Go!

---

# Deprecation notice
### This project was originally created to serve as an addon to `sync` with arrays with a very weird approach. This project was merged under [sync](https://github.com/Dviih/sync) as `sync.Slice[T]` and uses a slice with a mutex for its operations.

---

## Array
### This implements an array backed by a `sync.Mutex`.

## Channel
### This implements an `Array` with channel logic included.

---

# Usage

## Array
- `Index`: Returns an element from array.
- `Len`: Returns the length of the array.
- `Cap`: Returns the capacity of the array.
- `Append`: Appends an element to the array.
- `Remove`: Deletes an element from array.
- `Each`: Ranges through the elements of the array and executes a function.
- `Range` - Ranges through the elements of the array including index and executes a function.
- `Array` - Returns a shallow copy of the array.

## Channel
- `Send`: Sends T to a channel array.
- `Receive`: Receives T from a channel from a channel array.
- `Close`: Closes a channel array.
- `Index`: Returns an element from array.
- `Remove`: Removes an element from array.
- `Array`: Returns a shallow copy of the array.
- `Len`: Returns the length of the array.

# Example

```go
package main

import (
	"fmt"
	"github.com/Dviih/Array"
)

type Test struct {
	Value string
}

func main() {
	m := Array.New[*Test]()
	
	m.Append(&Test{Value: "Test Value"})
	
	fmt.Println(m.Index(0))
	m.Remove(0)
}
```
###### More can be found at `array_test.go`.

---

#### Made for Gophers by @Dviih
