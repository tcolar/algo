This is an implemntation to solve the given problem (See problem description in problem.txt)

I Only used standard Go (no dependencies)

The main method solves the given problem, example of running it:

$> go run main.go wordfinder.go
   Longest word: HONORIFICABILITUDINITATIBUS

There are also various tests in wordfinder_test.go
$> go test

----

How it works:

First I created a couple utility methods in main.go

- readUniqueWords : It reads the input text from file and return a list of UNIQUE words
normalized to uppercase. I used the pattern [a-zA-Z]+(-[a-zA-Z]+)* as the "word" definition
, see the method comments for more details.

- readGrid : Reads a grid from a text file and creates a [][]rune from it.

Now Here is what the algo does (Wordfinder.go):

1) Loads all the unique words into a suffixarray as that allows quick words lookup (by prefix)
2) For each postion in the grid find the longest word we can make starting there (recursive)
3) Return the logest word overall that we found in the grid.


