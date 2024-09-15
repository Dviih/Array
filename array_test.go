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
	"testing"
)

type Test struct {
	Name  string
	Other int
}

var (
	Tests = []*Test{
		{
			Name:  "One",
			Other: 1,
		},
		{
			Name:  "Two",
			Other: 2,
		},
		{
			Name:  "Three",
			Other: 3,
		},
	}
	array Array[*Test]
	_chan Chan[*Test]
)

func TestArrayAppend(t *testing.T) {
	array.Append(Tests[0])
}

func TestArrayIndex(t *testing.T) {
	t.Logf("Found at 0: %+v", array.Index(0))
}

func TestArrayEach(t *testing.T) {
	array.Each(func(test *Test) bool {
		t.Logf("My name is %s and I have the other as %d\n", test.Name, test.Other)
		return true
	})

	t.Log("Each done")
}

func TestArrayRemove(t *testing.T) {
	array.Append(Tests[2])
	array.Remove(0)
}

func TestChan(t *testing.T) {
	go func() {
		for {
			select {
			case data := <-_chan.Receive():
				t.Logf("Received: %+v\n", data)
			}
		}
	}()

	t.Log("Receiver was created, waiting for data")
	_chan.Send(Tests[2])
}

func TestChanClose(t *testing.T) {
	_chan.Close()
}

func TestChanArray(t *testing.T) {
	t.Logf("Returned from index 0: %+v", _chan.Array().Index(0))
}
