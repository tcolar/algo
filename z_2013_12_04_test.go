// History: Dec 04 13 tcolar Creation

package algo

import (
	"log"
	"testing"
)

//How will you find if a string is a substring of another string in O(n) complexity.
// For example, "tl" is substring of bottle.
func TestSubstring(t *testing.T) {
	if IsSubtring("bottle", "") != true {
		log.Panic("err1")
	}
	if IsSubtring("bottle", "abot") != false {
		log.Panic("err2")
	}
	if IsSubtring("bottle", "bot") != true {
		log.Panic("err3")
	}
	if IsSubtring("bottle", "tle") != true {
		log.Panic("err4")
	}
	if IsSubtring("", "bot") != false {
		log.Panic("err5")
	}
	if IsSubtring("bottltle", "tle") != true {
		log.Panic("err6")
	}
}

// This is failry linear (0n) except if a ton of backtracking is needed (near matches)
// Otherwise we could use a suufix tree / Ukkonen tree but that's quite complex
func IsSubtring(full string, sub string) bool {
	if len(sub) == 0 {
		return true
	}
	index := 0
	sl := len(sub)
	for i := 0; i < len(full); i++ {
		c := rune(full[i])
		if c == rune(sub[index]) {
			index++
			if index >= sl {
				log.Printf("Found %s at index %d", sub, i-sl+1)
				return true
			}
		} else {
			// "bactrack" to charcater after last match started
			i -= index
			index = 0
		}
	}
	return false
}

//Design a datastructure for the below functions.
//void add(T elem);
//void remove(T elem);
//T elem removeRandom();
//The operation time requirement : O(1)
func DataStruct() {
	// -> would do hashtable with keys kept in an array so removeRandom could pick one by random index
}
