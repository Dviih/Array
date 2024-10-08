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

// Deprecated: Array or Chan
type Interface[T interface{}] interface {
	Empty() bool
}

// Deprecated: Array
type SyncArray[T interface{}] struct {
	Array[T]
}

func (syncArray *SyncArray[T]) Append(t T) bool {
	syncArray.Append(t)
	return true
}

func (syncArray *SyncArray[T]) Remove(i int) bool {
	syncArray.Remove(i)
	return true
}

func (syncArray *SyncArray[T]) Each(fn func(T) bool) bool {
	syncArray.Each(fn)
	return true
}
