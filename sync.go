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

func (syncArray *SyncArray[T]) Remove(i int) bool {
	if syncArray == nil {
		return false
	}

	syncArray.m.Del(i)

	for ; i < len(syncArray.m.Data); i++ {
		syncArray.m.Push(i-1, syncArray.Index(i))
	}

	return true
}

