# Array
### Yet another way to mess with arrays all in Go!

---

## Sync
### All based in [SyncMap](https://github.com/Dviih/syncmap/) basically an array as a map!

## Channel
### Same as sync but you also have channels here so you can create a channel to send data of your generic T and when `Append` is called it will send for each channel!

---

# Usage
A simple guide for usage, more chan be found in the example!

## Sync
- `Index`: Gets element from index, similar to `array[i]`.
- `Append`: Appends element to array, similar to `append` built-in function.
- `Remove`: Deletes element from array, similar to slicing.
- `Each`: Gets each element and run a function through all of them.

## Channel
All of above plus:

- `Create`: Creates a channel for T, provide name.
- `Append`: Inherits from Sync but also sends to channel the incoming data.
- `Get`: Gets a channel by its name.
- `Close`: Closes a channel by its name.

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

func (test *Test) Empty() bool {
	if test == nil || test.Value == "" {
		return true
	}

	return false
}

func main() {
	m := Array.SyncArray[*Test]{}
	m.Append(&Test{"Test Value"})
	fmt.Println(m.Index(0))
	m.Remove(0)
}
```
###### More can be found at `array_test.go`.

---

#### Made for Gophers by @Dviih
