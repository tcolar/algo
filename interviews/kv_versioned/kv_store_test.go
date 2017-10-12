package main

import "testing"

func TestSimple(t *testing.T) {
}

func testInputOutput(inputFile, outputFile string) {
	kv := NewKeyValueStore()
	//log.SetOutput(os.Stderr)
	kv.ProcessInput("./samples/input.txt")
	// TODO check output
}
