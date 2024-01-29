package Array

import (
	syncmap "github.com/Dviih/SyncMap"
)

type ChanArray[T Interface[T]] struct {
	SyncArray[T]
	channels syncmap.SyncMap[string, chan T]
}

