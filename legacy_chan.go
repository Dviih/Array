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

// Deprecated: Chan
type ChanArray[T interface{}] struct {
	Chan[T]
}

func (chanArray *ChanArray[T]) Create(_ string, _ int) chan T {
	return chanArray.Get("")
}

func (chanArray *ChanArray[T]) Append(t T) bool {
	chanArray.Chan.Send(t)
	return true
}

func (chanArray *ChanArray[T]) Get(_ string) chan T {
	return chanArray.sender
}

func (chanArray *ChanArray[T]) Close(_ string) bool {
	chanArray.Chan.Close()
	return true
}
