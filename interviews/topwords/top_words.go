package topwords

import (
	"bufio"
	"io"
	"unicode"
)

// Finder can be used to search the top words in a text.
// Note: "top"here means recurring the most often.
type Finder struct {
}

// NewFinder creates a new TopWords finder
func NewFinder() *Finder {
	return &Finder{}
}

// Find finds the top words in a text stream
func (f *Finder) Find(input io.Reader) (TopWords, error) {
	reader := bufio.NewReader(input)
	wordCounts := map[string]int{} // stores count per individual word (# of occurences)
	topWords := TopWords{}         // Store on the fly the current top words
	mostUsed := 0                  // number of occurence of the current top word
	for {
		word, err := f.nextWord(reader)
		if err == io.EOF {
			break // end of stream
		}
		if err != nil {
			return nil, err
		}
		// increases the word occurence count
		count, _ := wordCounts[word]
		count++
		wordCounts[word] = count

		// Build on the fly of a map of the current top words(most occurences)
		// This uses a bit more memory but saves us from scanning through all of wordCounts at the end
		if count > mostUsed {
			// if it's a new "record" as far as most occurences, since it's now the only
			// top word, the existing top words are no longer valid, so drop them
			mostUsed = count
			topWords = map[string]int{}
		}
		if count >= mostUsed {
			// this is currently a top word, store it (automaticaly deduped since we use a map)
			topWords[word] = count
		}
	}
	// return the list of top words
	return topWords, nil
}

// Reads the next word from the stream, skipping non-word characters
// May return :
//  - a word (string)
//  - an io.EOF error if we reached the end of the stream
//  - some other unexpected error
func (f *Finder) nextWord(reader *bufio.Reader) (string, error) {
	var word []rune
	for {
		r, _, err := reader.ReadRune()
		if err == io.EOF {
			if len(word) != 0 { // return the last word of the stream
				return string(word), nil
			}
			return "", err
		}
		if err != nil {
			return "", err // unexpected error
		}
		if unicode.IsLetter(r) {
			word = append(word, r) // a plain letter or number
			continue
		}
		if r == '-' || r == '\'' { // allow ' and - within a word
			next, _, _ := reader.ReadRune() // "peek" at next rune to see if within a word
			reader.UnreadRune()
			if len(word) > 0 && unicode.IsLetter(next) {
				word = append(word, r) // we are whithin a word, allow those characters
				continue
			}
		}
		// we reached a "non-word" character, if we had built a word p to this point, return it
		if len(word) > 0 {
			return string(word), nil
		}
	}
}

// TopWords is a map of top words (with ocuurence count)
type TopWords map[string]int

// Words returns the words(keys) of TopWords
func (t *TopWords) Words() []string {
	words := make([]string, len(*t))
	for word := range *t {
		words = append(words, word)
	}
	return words
}
