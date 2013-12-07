// History: Dec 06 13 tcolar Creation

package algo

import (
	"log"
	"testing"
)

//7.4 write methods to implement *, / and - for integers using only '+' ops.
func TestOps(t *testing.T) {
	log.Print(multiply(3, 101))
	log.Print(divide(305, 3))
	log.Print(substract(82, 57))
}

// Lame but any better ways ?? (probably not allowed to use bit manipulation ?)
func substract(a uint32, b uint32) uint32 {
	if a < b {
		b, a = a, b
	}
	dif := 0
	for b+uint32(dif) < a {
		dif++
	}
	return uint32(dif)
}

func divide(base uint32, by uint32) int {
	cpt := 0
	for v := 0; v+int(by) < int(base); {
		cpt++
		v += int(by)
	}
	return cpt
}

func multiply(base uint32, by uint32) int {
	v := 0
	if by > base {
		base, by = by, base // optimize
	}
	for i := 0; i != int(by); i++ {
		v += int(base)
	}
	return v
}

func negate(a uint32) uint32 {
	return a
}
