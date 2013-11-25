// History: Nov 23 13 tcolar Creation

package algo

import (
	"log"
	"testing"
)

// Test various list aspects
func TestList(t *testing.T) {
	l := List{}
	check(l.Length, 0)
	l.Push(1)
	check(l.Length, 1)
	l.Clear()
	check(l.Length, 0)
	l.Append(1)
	check(l.Length, 1)
	l.Append(2)
	check(l.Length, 2)
	l.Push(3)
	check(l.Length, 3)
	check(l.Get(0), 3)
	check(l.Get(1), 1)
	check(l.Get(2), 2)
	check(l.Peek(), 3)
	check(l.Peek(), 3)
	check(l.Length, 3)
	l.Insert(9, 2)
	check(l.Length, 4)
	check(l.Get(0), 3)
	check(l.Get(1), 1)
	check(l.Get(2), 9)
	check(l.Get(3), 2)
	l.Set(12, 2)
	check(l.Length, 4)
	check(l.Get(2), 12)
	check(l.Pop(), 3)
	check(l.Length, 3)
	l.Remove(1)
	check(l.Length, 2)
	check(l.Get(0), 1)
	check(l.Get(1), 2)
	l.Remove(1)
	check(l.Length, 1)
	check(l.Pop(), 1)
	check(l.Length, 0)
}

// Test reversng lists (in place)
func TestListReverse(t *testing.T) {
	l := List{}
	l.Reverse()
	check(l, List{})
	l.Append(1)
	l.Reverse()
	check(l.Length, 1)
	check(l.ToSlice()[0], 1)
	l.Clear()
	l.Append(1)
	l.Append(2)
	l.Reverse()
	check(l.Length, 2)
	check(l.ToSlice()[0], 2)
	check(l.ToSlice()[1], 1)
	l.Clear()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Reverse()
	check(l.Length, 3)
	check(l.ToSlice()[0], 3)
	check(l.ToSlice()[1], 2)
	check(l.ToSlice()[2], 1)
}

func check(v1 interface{}, v2 interface{}) {
	if v1 != v2 {
		log.Print(v1)
		log.Print(v2)
		log.Fatal("Failed")
	}
}
