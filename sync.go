package Array

import (
	syncmap "github.com/Dviih/SyncMap"
)

type SyncArray[T Interface[T]] struct {
	m syncmap.SyncMap[int, T]
}

func (syncArray *SyncArray[T]) Index(i int) T {
	return syncArray.m.Get(i)
}

func (syncArray *SyncArray[T]) Append(t T) bool {
	if syncArray == nil {
		return false
	}

	syncArray.m.Add(len(syncArray.m.Data), t)

	return true
}

