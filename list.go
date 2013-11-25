// History: Nov 23 13 tcolar Creation

package algo

import ()

// Linked list simplistic impl ...
type List struct {
	Length int
	Root   *elem
}

func (l *List) Append(val interface{}) {
	cur := l.Root
	if cur == nil {
		l.Root = &elem{Val: val}
		l.Length++
		return
	}
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = &elem{Val: val}
	l.Length++
}

func (l *List) Insert(val interface{}, index int) {
	if index == 0 {
		l.Push(val)
		return
	}
	e := l.elemAt(index - 1)
	e.Next = &elem{Val: val, Next: e.Next}
	l.Length++
}

func (l *List) Get(index int) interface{} {
	return l.elemAt(index).Val
}

func (l *List) Set(val interface{}, index int) {
	l.elemAt(index).Val = val
}

func (l *List) Remove(index int) {
	if index == 0 {
		l.Pop()
		return
	}
	prev := l.elemAt(index - 1)
	prev.Next = prev.Next.Next
	l.Length--
}

func (l *List) Clear() {
	l.Root = nil
	l.Length = 0
}

func (l *List) Index(val interface{}, from int) int {
	elem := l.elemAt(from)
	for elem != nil {
		if elem.Val == val {
			return from
		}
		elem = elem.Next
		from++
	}
	return -1
}

func (l *List) Push(val interface{}) {
	l.Root = &elem{Val: val, Next: l.Root}
	l.Length++
}

func (l *List) Pop() interface{} {
	val := l.Root.Val
	l.Root = l.Root.Next
	l.Length--
	return val
}

func (l *List) Peek() interface{} {
	return l.Root.Val
}

func (l *List) ToSlice() (slice []interface{}) {
	elem := l.Root
	for elem != nil {
		slice = append(slice, elem.Val)
		elem = elem.Next
	}
	return slice
}

// "internal" method to get a list element
func (l *List) elemAt(index int) *elem {
	elem := l.Root
	for i := 0; i < index; i++ {
		elem = elem.Next
	}
	return elem
}

// List element
type elem struct {
	Val  interface{}
	Next *elem
}
