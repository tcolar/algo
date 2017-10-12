package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatalf("Expected an input file path as an argument")
		os.Exit(1)
	}
	inputFile := os.Args[1]
	kv := NewKeyValueStore()
	err := kv.ProcessInput(inputFile)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
