package main

import (
	"flag"
	"log"
	"os"

	"github.com/tcolar/pantopoda"
)

func main() {
	config := pantopoda.DefaultConfig
	outFile := ""

	flag.BoolVar(&config.Verbose, "verbose", true, "Verbose output")
	flag.IntVar(&config.Routines, "routines", 5, "Number of routines o spawn.")
	flag.IntVar(&config.PauseBetweenRequests, "delay", 100, "Delay(ms) between each request (per routine)")
	flag.IntVar(&config.HttpTimeout, "timeout", 10, "Http timeout (in seconds)")
	flag.IntVar(&config.MaxLinksPerPage, "maxlinks", 500, "Max links per page")
	flag.StringVar(&outFile, "out", "", "Output file (Json).")
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		log.Fatal("Expected a domain name as a parameter.")
	}
	domain := args[0]

	spider := pantopoda.NewSpider(&config)

	// Crawl
	spider.CrawlDomain(domain)

	// Once dne, write json if an output file was specified.
	if len(outFile) > 0 {
		f, err := os.Create(outFile)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		spider.ToJson(f)
	}
}
