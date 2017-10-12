package main

import (
	"log"
	"os"

	"github.com/kr/pretty"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Expected a test file as the first argument")
		os.Exit(1)
	}

	// Read the shipment list
	shipments, err := newShipments(os.Args[1])
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// Create a bundler for a given test file
	bundler, err := NewBundler(shipments)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// find our bundles and print them to stdout
	bundles := bundler.FindBundles()
	for _, bundle := range bundles {
		pretty.Println(bundle)
	}
}
