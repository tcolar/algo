// History: Dec 11 13 tcolar Creation

package algo

import (
	"fmt"
	"log"
	"testing"
)

func TestHanoi(t *testing.T) {
	h := Hanoi{}
	// Initialize Tower 1 with 5 disks (smallest on top)
	h.T1.Push(5)
	h.T1.Push(4)
	h.T1.Push(3)
	h.T1.Push(2)
	h.T1.Push(1)

	// Get started by moving top disk to tower 2 and recurse from there
	h.HanoiMove(5, 1, 2)

	if h.T2.Pop() != 1 || h.T2.Pop() != 2 || h.T2.Pop() != 3 ||
		h.T2.Pop() != 4 || h.T2.Pop() != 5 {
		log.Fatal("Failed")
	}
}

// Towers of Hanoi problem
// Goal is to move all disks from t1 to t2 using only legal moves
// disk can only be set atop a larger one, can only move one at a a time
// can only move disk at top of tower
type Hanoi struct {
	// Towers
	T1 List
	T2 List
	T3 List
}

// Recursively move the disk stack
func (h *Hanoi) HanoiMove(nbdisk int, from int, to int) {
	// Smallest disk can be moved to bottom of tower
	if nbdisk == 1 {
		// only one disk to move, juts mve it
		h.Move(1, from, to)
		return
	}
	// Otherwise we got a stack of disks to move
	other := h.OtherTower(from, to)
	h.HanoiMove(nbdisk-1, from, other)
	h.Move(nbdisk, from, to)
	h.HanoiMove(nbdisk-1, other, to)
}

func (h *Hanoi) Move(disk int, from int, to int) {
	log.Printf("Moving Disk %d from %d to %d", disk, from, to)
	twFrom := h.Tower(from)
	twTo := h.Tower(to)
	d := twFrom.Pop()
	if d != disk {
		log.Fatalf("Expected disk %d, but got %d", disk, d)
	}
	if twTo.Length > 0 && twTo.Peek().(int) < d.(int) {
		log.Fatalf("Illegal move.")
	}
	twTo.Push(d)
	h.Print()
}

// Get a tower list by index(1-3)
func (h *Hanoi) Tower(index int) *List {
	if index == 1 {
		return &h.T1
	} else if index == 2 {
		return &h.T2
	} else {
		return &h.T3
	}
}

// Return the "other" tower than the ones we are moving from, to.
func (h *Hanoi) OtherTower(from int, to int) int {
	if (from == 1 && to == 2) || (from == 2 && to == 1) {
		return 3
	} else if (from == 1 && to == 3) || (from == 3 && to == 1) {
		return 2
	} else {
		return 1
	}
}

// Pretty print the tower
func (h *Hanoi) Print() {
	t1, t2, t3 := "", "", ""
	for i := 0; i != 5; i++ {
		if h.T1.Length > i {
			t1 += fmt.Sprintf("%d ", h.T1.Get(i).(int))
		}
		if h.T2.Length > i {
			t2 += fmt.Sprintf("%d ", h.T2.Get(i).(int))
		}
		if h.T3.Length > i {
			t3 += fmt.Sprintf("%d ", h.T3.Get(i).(int))
		}
	}
	log.Printf("T1[%10s]\tT2[%10s]\tT3[%10s]", t1, t2, t3)
}
