package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

var dataSets map[string][]int

func main() {
	rand.Seed(time.Now().UnixNano())

	initDataSets(25000)

	sorts := []Sorter{
		BubbleSort{},
		InsertionSort{},
		SelectionSort{},
		MergeSort{},
		QuickSort{},
		HeapSort{},
		// TODO Heap
		// TODO Shell
	}

	// execute sorts
	results := map[string]map[string]time.Duration{}
	for i, sorter := range sorts {
		// make sure the sorter impl actually works
		data := dataSets["small"]
		if err := checkSort(sorter, data, true); err != nil {
			panic(err)
		}
		// measure durations
		results[sorter.Name()] = map[string]time.Duration{}
		for setName, _ := range dataSets {
			if setName == "vhuge" && i < 3 { // first 3 algos are too slow
				results[sorter.Name()][setName] = time.Duration(0)
				continue
			}
			data := dataSets[setName]
			d := timeSort(sorter, data, false)
			results[sorter.Name()][setName] = d
		}
	}

	// show results table
	sets := make([]string, 0, len(dataSets))
	for set, _ := range dataSets {
		sets = append(sets, set)
	}
	sort.Strings(sets)

	fmt.Print(format("", 15))
	for _, set := range sets {
		fmt.Print(format(set, 8))
	}
	fmt.Println()

	for _, sorter := range sorts {
		fmt.Print(format(sorter.Name(), 15))
		for _, set := range sets {
			fmt.Print(formatDur(results[sorter.Name()][set], 8))
		}
		fmt.Println()
	}
}

func initDataSets(size int) {
	sorted := []int{}
	inverted := []int{}
	scrambled := []int{}
	random := []int{}
	dups := []int{}
	for i := 0; i != size; i++ {
		sorted = append(sorted, i)
		scrambled = append(scrambled, i)
		inverted = append(inverted, size-1-i)
		random = append(random, rand.Intn(size)) // random
		dups = append(dups, i%3)                 // lots of dups
	}
	for i := 0; i != size; i++ {
		a := rand.Intn(len(scrambled))
		b := rand.Intn(len(scrambled))
		scrambled[a], scrambled[b] = scrambled[b], scrambled[a]
	}

	huge := []int{}
	for i := 0; i != 10; i++ {
		huge = append(huge, random...)
	}

	dataSets = map[string][]int{
		"empty":   []int{},
		"single":  []int{1},
		"small":   []int{3, 9, -1, 0, 7, 4, 8, 8, 2, 6, 5, -2},
		"sorted":  sorted,                        //already sorted
		"invert":  inverted,                      //sorted backward
		"scrambl": scrambled,                     //uniques but scrambled
		"dups":    dups,                          // lots of dup values
		"rand":    random,                        // fully random
		"rand/2":  dup(random, 0, len(random)/2), // 1/2 size random
		"rand/4":  dup(random, 0, len(random)/4), // 1/4 size random
		"vhuge":   huge,
	}
}
