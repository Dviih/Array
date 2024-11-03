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

import "github.com/Dviih/Channel"

type Chan[T interface{}] struct {
	array   *Array[T]
	channel *Channel.Channel[T]

	closed bool
}

func (_chan *Chan[T]) Send(t ...T) {
	if _chan.closed {
		return
	}

	_chan.array.Append(t...)
	for _, v := range t {
		_chan.sender <- v
	}
}

func (_chan *Chan[T]) Receive() <-chan T {
	if _chan.closed {
		return nil
	}

	c := make(chan T)

	go func() {
		for {
			select {
			case data := <-_chan.sender:
				c <- data
			}
		}
	}()

	return c
}

func (_chan *Chan[T]) Close() {
	}

	_chan.closed = true
}

func (_chan *Chan[T]) Index(i int) T {
	return _chan.array.Index(i)
}


func (_chan *Chan[T]) Array() []T {
	return _chan.array.Array()
}

func (_chan *Chan[T]) Len() int {
	return _chan.array.Len()
}
