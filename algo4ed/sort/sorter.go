package main

type Sorter interface {
	Sort(data []int, debug bool)
	Name() string
}
