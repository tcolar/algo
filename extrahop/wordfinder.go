// History: May 22 14 tcolar Creation

package main

import (
	"index/suffixarray"
	"log"
	"strings"
)

// WordFinder helps us find the longest word in a given grid of letters
// by making "knights" moves only, see Readme.md for more details.
type WordFinder struct {
	grid       [][]rune
	rows, cols int // number of rows, cols in grid
}

// NewWordFinder creates a new word finder wth the given letter grid.
func NewWordFinder(grid [][]rune) *WordFinder {
	finder := WordFinder{
		grid: grid,
	}
	finder.rows = len(grid)
	if finder.rows > 0 {
		finder.cols = len(finder.grid[0])
	}
	return &finder
}

// FindLongest finds the longest of the given words that is found in the grid
// If none found, return an empty string
func (w WordFinder) FindLongest(words []string) string {

	// Create a suffix tree for faster word/path lookup
	joinedStrings := "\x00" + strings.Join(words, "\x00") + "\x00"
	sa := suffixarray.New([]byte(joinedStrings))

	longest := ""
	// Find the longest word from each position in the grid, one by one.
	// Keep the longest overall.
	for row := 0; row != w.rows; row++ {
		for col := 0; col != w.cols; col++ {
			found := w.longestFrom(sa, col, row, "")
			if len(found) > len(longest) {
				longest = found
			}
		}
	}

	return longest
}

// longestFrom finds the longest word
// that we can find in the grid starting at a given position.
// The word must be one that exists in the word list (suffixarray)
//
// Note: If needed we could optimize this to use []byte instead of creatng strings.
func (w WordFinder) longestFrom(sa *suffixarray.Index, col, row int, prefix string) string {
	longest := ""
	r := w.grid[row][col]
	word := prefix + string(r)
	if len(word) > 6 {
		log.Print(word)
	}
	// Are there any words starting with this
	hits := sa.Lookup([]byte("\x00"+word), 1)
	if len(hits) > 0 {
		// If so, then we should keep looking for longer words (recursive)
		for _, pos := range w.validMoves(col, row) {
			sub := w.longestFrom(sa, pos.col, pos.row, word)
			if len(sub) > len(longest) {
				// longest word so far
				longest = sub
			}
		}
	}
	// We did not find a longer word while recursing,
	// then check if "word" itself is aword from the word list.
	if len(longest) == 0 {
		hits = sa.Lookup([]byte("\x00"+word+"\x00"), 1)
		if len(hits) > 0 {
			//log.Printf("Found word: %s", string(word))
			longest = word
		}
	}

	return longest
}

// validMoves returns a list of valid (knight) positions in the grid we can
// move to from the position (row, col)
//
// Note: We could make this part of an interface if we wanted to provide different implementations
// later say "Queen moves" or such.
func (w WordFinder) validMoves(col, row int) []Pos {
	valids := []Pos{}
	// Check all 8 possible knigt moves
	pos := Pos{col - 2, row - 1}
	if w.isInGrid(pos) {
		valids = append(valids, pos)
	}
	pos = Pos{col - 1, row - 2}
	if w.isInGrid(pos) {
		valids = append(valids, pos)
	}
	pos = Pos{col - 2, row + 1}
	if w.isInGrid(pos) {
		valids = append(valids, pos)
	}
	pos = Pos{col - 1, row + 2}
	if w.isInGrid(pos) {
		valids = append(valids, pos)
	}
	pos = Pos{col + 1, row + 2}
	if w.isInGrid(pos) {
		valids = append(valids, pos)
	}
	pos = Pos{col + 2, row + 1}
	if w.isInGrid(pos) {
		valids = append(valids, pos)
	}
	pos = Pos{col + 2, row - 1}
	if w.isInGrid(pos) {
		valids = append(valids, pos)
	}
	pos = Pos{col + 1, row - 2}
	if w.isInGrid(pos) {
		valids = append(valids, pos)
	}
	return valids
}

// isInGrid tells us if the given position would be valid in the current grid (in bounds)
func (w WordFinder) isInGrid(p Pos) bool {
	if p.col < 0 || p.row < 0 {
		return false
	}
	if p.col >= w.cols || p.row >= w.rows {
		return false
	}
	return true
}

// Pos represents a position (col, row pair)
type Pos struct {
	col, row int
}
