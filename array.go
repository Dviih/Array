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
	"sync"
)

type Array[T interface{}] struct {
	m     sync.Mutex
	array []T
}

func (array *Array[T]) Index(i int) T {
	return array.array[i]
}

func (array *Array[T]) Len() int {
	return len(array.array)
}

func (array *Array[T]) Cap() int {
	return cap(array.array)
}

func (array *Array[T]) Append(t ...T) {
	defer array.m.Unlock()
	array.m.Lock()

	array.array = append(array.array, t...)
}

func (array *Array[T]) Remove(i int) {
	defer array.m.Unlock()
	array.m.Lock()

	array.array = append(array.array[:i], array.array[i+1:]...)
}

func (array *Array[T]) Each(fn func(T) bool) {
	for _, t := range array.array {
		fn(t)
	}
}
