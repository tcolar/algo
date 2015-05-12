// History: May 22 14 tcolar Creation

package main

import "testing"

// Test the example given by ExtraHop
func TestExtraHop(t *testing.T) {
	grid := readGrid("data/test_grid.txt", 8, 8)
	words := []string{"EXTRA", "HOP", "NETWORK"}
	finder := NewWordFinder(grid)
	checkString(t, finder.FindLongest(words), "EXTRA")
}

// Test the Loves Labours Lost exercise
func TestLovesLabourLost(t *testing.T) {
	grid := readGrid("data/grid.txt", 8, 8)
	words := readUniqueWords("data/input_text.txt")
	finder := NewWordFinder(grid)
	// If my algo worked roperly that is the longest word in the grid
	checkString(t, finder.FindLongest(words), "HONORIFICABILITUDINITATIBUS")
}

// Test a few other grid word searches
func TestGridSerach(t *testing.T) {
	grid := [][]rune{
		[]rune{'B', 'C', 'D', 'E', 'N'},
		[]rune{'X', 'Y', 'A', 'Z', 'Z'},
	}
	words := []string{"P", "BANGO"}
	finder := NewWordFinder(grid)
	checkString(t, finder.FindLongest(words), "") // no match
	words = []string{"A", "K"}
	finder = NewWordFinder(grid)
	checkString(t, finder.FindLongest(words), "A")
	words = []string{"X", "XD", "XDZ", "XDZN"}
	finder = NewWordFinder(grid)
	checkString(t, finder.FindLongest(words), "XDZ")
	words = []string{"XDZ", "BANANA", "BANANAC"}
	finder = NewWordFinder(grid)
	checkString(t, finder.FindLongest(words), "BANANA")
}

// Test the word reader that should retunr unique words normalized to Upercase
// see readUniqueWords for details
func TestWordReader(t *testing.T) {
	words := readUniqueWords("data/test_text.txt")
	if len(words) != 6 {
		t.Errorf("Expected 6 words, got %d", len(words))
	}
	checkString(t, words[0], "FUU")
	checkString(t, words[1], "BOO")
	checkString(t, words[2], "FOO-BAR")
	checkString(t, words[3], "BAR")
	checkString(t, words[4], "BAZ")
	checkString(t, words[5], "A")
}

// Test the readGrid method
func TestReadGrid(t *testing.T) {
	grid := readGrid("data/test_grid.txt", 8, 8)
	if len(grid) != 8 {
		t.Errorf("Expected 8 rows, got %d", len(grid))
	}
	if len(grid[0]) != 8 {
		t.Errorf("Expected 8 cols, got %d", len(grid[0]))
	}
	checkString(t, string(grid[0][0]), "Q")
	checkString(t, string(grid[7][7]), "S")
	checkString(t, string(grid[2][0]), "R")
	checkString(t, string(grid[5][6]), "X")
}

func checkString(t *testing.T, got, expected string) {
	if got != expected {
		t.Errorf("Expected %s, got %s", expected, got)
	}
}
