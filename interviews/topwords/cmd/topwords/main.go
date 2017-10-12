package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/tcolar/topwords"
)

func main() {
	if len(os.Args) > 1 && (os.Args[1] == "-h" || os.Args[1] == "--help") {
		log.Println("topwords reads from the input stream and returns the words with the highest occurence")
		return
	}
	finder := topwords.NewFinder()
	topWords, err := finder.Find(os.Stdin)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Println(strings.Join(topWords.Words(), " "))
}
