// History: May 21 14 tcolar Creation

package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

// Main method currently runs the requested problem using.
// data/grid.txt and data/input_text.txt
// In real workd we could make the files arguments or use a confile files.
func main() {

	grid := readGrid("data/grid.txt", 8, 8)
	words := readUniqueWords("data/input_text.txt")

	finder := NewWordFinder(grid)
	longest := finder.FindLongest(words)

	log.Printf("Longest word: %s", string(longest))
}

// readUniqueWords is a helper method to find all UNIQUE words in a text file
// words are all normalized to uppercase
// a word is defined as a case insensitive combination of letters 'a' to 'z'
// Hiphens ('-') are also allowed "within" a word.
// For now simply fail eraly if an error occurs
func readUniqueWords(filePath string) []string {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	r, err := regexp.Compile("[a-zA-Z]+(-[a-zA-Z]+)*")
	if err != nil {
		log.Fatal(err)
	}
	all := map[string]bool{}
	for _, s := range r.FindAll(bytes, -1) {
		s := strings.ToUpper(string(s))
		if _, ok := all[s]; !ok {
			all[s] = true
		}
	}
	words := []string{}
	for k, _ := range all {
		words = append(words, k)
	}
	return words
}

// readGrid is a helper method that reads a Grid of w*h letters from a text file
// For now simply fail early if an error occurs
func readGrid(filePath string, w, h int) [][]rune {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	grid := make([][]rune, h)
	for row := 0; row != h; row++ {
		line := make([]rune, w)
		for col := 0; col != w; col++ {
			if !scanner.Scan() {
				log.Fatalf("Not enough data!")
			}
			x := scanner.Text()
			if len(x) != 1 {
				log.Fatal("Expecting only single characters.")
			}
			line[col] = rune(x[0])
		}
		grid[row] = line
	}
	return grid
}
