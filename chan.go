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

func (chanArray *ChanArray[T]) Append(t T) bool {
	if chanArray == nil {
		return false
	}

	chanArray.m.Add(len(chanArray.m.Data), t)
	chanArray.channels.RLock()

	for _, channel := range chanArray.channels.Data {
		channel <- t
	}

	chanArray.channels.RUnlock()
	return true
}

func (chanArray *ChanArray[T]) Get(name string) chan T {
	return chanArray.channels.Get(name)
}

