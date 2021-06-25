/*
Given a dictionary of words and a string made up of those words (no spaces), return the original sentence in a list.
If there is more than one possible reconstruction, return any of them. If there is no possible reconstruction, then return null.

For example, given the set of words 'quick', 'brown', 'the', 'fox', and the string "thequickbrownfox", you should return ['the', 'quick', 'brown', 'fox'].

Given the set of words 'bed', 'bath', 'bedbath', 'and', 'beyond', and the string "bedbathandbeyond",
return either ['bed', 'bath', 'and', 'beyond] or ['bedbath', 'and', 'beyond'].

*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	result := wordBreak("bedbathandbeyond", []string{"bed", "bath", "bedbath", "and", "beyond"}, []string{})
	fmt.Println(result)
	result = wordBreak("thequickbrownfox", []string{"quick", "brown", "fox", "the"}, []string{})
	fmt.Println(result)
}

func wordBreak(s string, unusedWords []string, usedWords []string) []string {
	fmt.Println(s, unusedWords, usedWords)
	if len(s) == 0 {
		return usedWords
	}
	for i, word := range unusedWords {
		if strings.HasPrefix(s, word) {
			unused := append(unusedWords[:i], unusedWords[i+1:]...)
			return wordBreak(s[len(word):], unused, append(usedWords, word))
		}
	}
	return nil
}
