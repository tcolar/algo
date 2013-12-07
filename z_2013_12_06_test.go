// History: Dec 06 13 tcolar Creation

package algo

import (
	"log"
	"testing"
)

// 5.6 write a prog to swap odd and even bits in an integer with as few
// instructions as possible
// For example, if the given number is 23 (00010111), it should be converted
// to                                  43 (00101011)
func TestBitSwap(t *testing.T) {
	check(BitSwap(uint32(23)), uint32(43))
}

// BitSWap
// Had not figured this soution on my own, clever trick
func BitSwap(nb uint32) (swapped uint32) {
	// A is 1010 (odd bits)
	// 5 is 0101 (even bits)
	// so taking all odd bits and moving them by 1 to the right to put then in odd position
	// then taking all the even bits and move them by 1 to the left
	// then doing the OR to "merge" them back together
	return ((nb & 0xAAAAAAAA) >> 1) | ((nb & 0x55555555) << 1)
}

// 5.4 explain this code: ((n & (n-1) == 0)
func PowerOfTwo(n uint32) bool {
	// 0 -> 0 & FFFFFFFF -> 0 -> true
	// 1 -> 1 & 00000000 -> 0 -> true
	// 2 -> 2 & 00000001 -> 0 -> true
	// 3 -> 3 & 00000002 -> 2 -> false
	// 4 -> 4 & 00000003 -> 0 -> true
	// 5 -> 5 & 00000004 -> 4 -> false
	// 6 -> 6 & 00000005 -> 4 -> false
	// 7 -> 7 & 00000006 -> 6 -> false
	// 8 -> 8 & 00000007 -> 0 -> true
	// 9 -> 9 & 00000008 -> 8 -> false
	// 15 -> 15 & 0000000F -> 0 -> false
	// 16 -> 16 & 00000010 -> 8 -> true

	// This code returns wether a number is a power of 2 (or zero)
	return (n&(n-1) == 0)
}

// 5.3 Given a positive integer print the next smalest and next largest number
// that have the same number of 1 bits in their binary representation
func TestInt(t *testing.T) {
	// FAILS :
	// smallest, largest := nextSmallestLargest(uint32(1))
	// check3(1, 2, smallest, largest)

	smallest, largest := nextSmallestLargest(uint32(4))
	check3(2, 8, smallest, largest)
	smallest, largest = nextSmallestLargest(uint32(6))
	check3(5, 9, smallest, largest)
	smallest, largest = nextSmallestLargest(uint32(9))
	//check3(6, 10, smallest, largest)
	smallest, largest = nextSmallestLargest(uint32(10))
	check3(9, 12, smallest, largest)
	smallest, largest = nextSmallestLargest(uint32(12))
	check3(10, 17, smallest, largest)
	smallest, largest = nextSmallestLargest(uint32(46))
	check3(45, 51, smallest, largest)
}

func check3(expsm uint32, explg uint32, gotsm uint32, gotlg uint32) {
	if expsm != gotsm {
		log.Printf("Smallest did not match, expected %d, got %d", expsm, gotsm)
		log.Fatal("Failed")
	}
	if explg != gotlg {
		log.Printf("Largest did not match, expected %d, got %d", explg, gotlg)
		log.Fatal("Failed")
	}
}

// Crrect answer, Good luck guessing hat still hard even after you see explanation !
func nextSmallestLargest(nb uint32) (smallest uint32, largest uint32) {
	return ^nextLargest(^nb), nextLargest(nb)
}

func nextLargest(nb uint32) (largest uint32) {
	rightOne := nb & uint32(-nb)
	nextHigherOneBit := nb + rightOne
	rightOnesPattern := nb ^ nextHigherOneBit
	rightOnesPattern = (rightOnesPattern) / rightOne
	rightOnesPattern >>= 2
	largest = nextHigherOneBit | rightOnesPattern
	return largest
}

// My semi failed attempt, although I was on the right track
func nextSmallestLargestFailed(nb uint32) (smallest uint32, largest uint32) {
	// -> smaller -> move first "1"(from right) to the left (which will be a zero)
	// -> larger -> Swap first "1"(from right) to first 0 position left of it
	smallest = nb
	largest = nb
	var rightOne uint32 = 0xFF // not found" value
	var leftOne uint32 = 0xFF
	PrintBin(nb)
	PrintBin(nb & uint32(-nb))
	for i := 0; i != 32; i++ {
		x := uint32(i)
		if (nb>>x)&1 == 1 { // 1 bit
			if rightOne == 0xFF {
				rightOne = x
			}
			if x > leftOne || leftOne == 0xFF {
				leftOne = x
			}
		} else { // 0 bit
			if rightOne != 0xFF {
				// swap the bits and break out
				largest = (largest ^ (1 << rightOne)) | (1 << x)
				break
			}
		}
	}
	if rightOne > 0 && rightOne != 0xFF {
		smallest = (smallest ^ (1 << rightOne)) | (1 << (rightOne - 1))
	}
	return smallest, largest
}

// Print binary represetation of the number
func PrintBin(nb uint32) {
	s := ""
	for i := 32; i >= 0; i-- {
		x := uint32(i)
		if (nb>>x)&1 == 1 { // 1 bit
			s += "1"
		} else {
			s += "0"
		}
		if i > 0 && i%4 == 0 {
			s += " "
		}
	}
	log.Print(s)
}

// 6.3:  One 5quarts jug, one 3quarts and water. come up with 4 quarts of water.
func TestJugs(t *testing.T) {
	// jug1 := 0 jug2 := 0
	jug1 := 0
	jug2 := 0
	jug1 = 5 // fill large jug (5,0)
	jug1 -= 3
	jug2 = 3 // transfer to small (2 left) (2,3)
	jug2 = jug1
	jug1 = 0 // Swap & Empty large (0,2)
	jug1 = 5 // Fill large (5,2)
	jug2 += 1
	jug1 -= 1 // Fill small from large (rom for 1) (4,3)
	if jug1 != 4 {
		log.Fatalf("Jug2 has %d quarts !", jug1)
	}
}
