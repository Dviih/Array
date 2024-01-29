package Array

import (
	syncmap "github.com/Dviih/SyncMap"
)

type ChanArray[T Interface[T]] struct {
	SyncArray[T]
	channels syncmap.SyncMap[string, chan T]
}

func (chanArray *ChanArray[T]) Create(name string, size int) chan T {
	c := make(chan T, size)
	chanArray.channels.Add(name, c)
	return c
}

