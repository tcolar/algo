package topwords

import (
	"strings"
	"testing"
)

// various basic tests of the TopWords finder
func TestTopWords(t *testing.T) {
	test(t, "", []string{}) // nothing

	test(t, ";;..", []string{}) // no words

	test(t, "single", []string{"single"})

	test(t, "one two", []string{"one", "two"})

	test(t, "one two three two", []string{"two"})

	test(t, ";;one four \r\n three four,.,", []string{"four"})

	text := "A blue shirt cost is twenty-four dollars but a white shirt is only twenty so I bought the white shirt."
	test(t, text, []string{"shirt"})
}

// Tests for "special" characters ' and - which are only allowed "within" a word
func TestSpecials(t *testing.T) {
	test(t, "ain't 'nope ain't nope'", []string{"ain't", "nope"}) // apostrophes

	test(t, "twenty-two -twentytwo twenty-two twentytwo-", []string{"twenty-two", "twentytwo"}) // dashes
}

// Test a large text file
func TestLarge(t *testing.T) {
	text := "A blue shirt cost is twenty-four dollars but a white shirt is only twenty so I bought the white shirt."
	// Repeat this text a hundred thousand times
	large := make([]rune, len(text)*100000)
	for i := 0; i != 100000; i++ {
		copy(large[i*len(text):], []rune(text))
	}
	test(t, string(large), []string{"shirt"})
}

func TestWithLoremIpsum(t *testing.T) {
	text := `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna
 aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. 
 Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint 
 occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.`
	test(t, text, []string{"in"})
}

// TestWorstCase tests a case where all the words are the same length and unique
// This is the worst case in terms of memory usage
func TestWorstCase(t *testing.T) {
	text := "aaa bbb ccc ddd eee fff ggg hhh iii jjj kkk lll mmm nnn ooo ppp qqq rrr sss ttt uuu vvv www xxx yyy zzz"
	test(t, text, strings.Fields(text))
}

func test(t *testing.T, text string, expected []string) {
	finder := NewFinder()
	input := strings.NewReader(text)
	topWords, err := finder.Find(input)
	if err != nil {
		t.Fatal(err)
	}
	words := []string{}
	for word := range topWords {
		words = append(words, word)
	}
	if len(words) != len(expected) {
		t.Fatalf("Expected %v, got %v", expected, words)
	}
	for _, word := range expected {
		if _, found := topWords[word]; !found {
			t.Fatalf("Expected to find '%s' in %v", word, words)
		}
	}
}
