/*
 *     Implements simple synchronous array for Go.
 *     Copyright (C) 2024  Dviih
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU Affero General Public License as published
 *     by the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU Affero General Public License for more details.
 *
 *     You should have received a copy of the GNU Affero General Public License
 *     along with this program.  If not, see <https://www.gnu.org/licenses/>.
 *
 */

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

func (chanArray *ChanArray[T]) Close(name string) bool {
	c := chanArray.Get(name)
	if c == nil {
		return false
	}

	close(c)

	chanArray.channels.Del(name)
	return true
}
