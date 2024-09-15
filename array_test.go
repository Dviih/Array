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
	"fmt"
	"testing"
	"time"
)

type Test struct {
	Name  string
	Other int
}

var Tests = []*Test{
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

func (test *Test) Empty() bool {
	if test == nil || (test != nil && test.Name == "" && test.Other == 0) {
		return true
	}

	return false
}

var syncArray = SyncArray[*Test]{}
var chanArray = ChanArray[*Test]{}

const (
	ChannelOne = "channel0"
	ChannelTwo = "channel1"
	Size       = 1024 * 8
)

func TestSyncArrayAppend(t *testing.T) {
	syncArray.Append(Tests[0])
}

func TestSyncArrayIndex(t *testing.T) {
	t.Logf("Found at 0: %+v", syncArray.Index(0))
	t.Logf("Found at 1: %+v", syncArray.Index(1))
	t.Log("Push two")
	syncArray.Append(Tests[1])
	t.Logf("Found at 1: %+v", syncArray.Index(1))
}

func TestSyncArray_Each(t *testing.T) {
	syncArray.Each(func(test *Test) bool {
		t.Logf("My name is %s and I have the other as %d\n", test.Name, test.Other)
		return true
	})
	t.Log("Each done")
}

func TestSyncArray_Remove(t *testing.T) {
	syncArray.Append(Tests[2])
	t.Logf("Found at 0: %+v", syncArray.Index(0))
	syncArray.Remove(0)
	t.Logf("Found at 0: %+v", syncArray.Index(0))
	t.Logf("Found at 1: %+v", syncArray.Index(1))
}

func TestChanArrayInit(t *testing.T) {
	chanArray.Append(Tests[0])
	chanArray.Append(Tests[1])
	t.Log("Initialized channel array with 2 members!")
}

func TestChanArrayCreate(t *testing.T) {
	chanArray.Create(ChannelOne, Size)
	chanArray.Create(ChannelTwo, Size)
}

func TestChanArraySetup(t *testing.T) {
	t.Log("Creating go routines and loops to handle channels!")

	go func() {
		for {
			select {
			case data := <-chanArray.Get(ChannelOne):
				fmt.Printf("Received from %s: %+v\n", ChannelOne, data)
			}
		}
	}()

	go func() {
		for {
			select {
			case data := <-chanArray.Get(ChannelTwo):
				fmt.Printf("Received from %s: %+v\n", ChannelTwo, data)
			}
		}
	}()

	t.Log("Channels are setup, waiting for communication!")
}

func TestChanArrayAppend(t *testing.T) {
	chanArray.Append(Tests[2])
	t.Log("Sent append to channel array, sleeping for 1 seconds!")
	time.Sleep(time.Second * 1)
}

func TestChanArrayClose(t *testing.T) {
	t.Logf("Closing %s", ChannelOne)
	chanArray.Close(ChannelOne)
	t.Logf("Closing %s", ChannelTwo)
	chanArray.Close(ChannelTwo)
}
