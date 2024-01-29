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

