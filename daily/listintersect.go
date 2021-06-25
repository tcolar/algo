package main

import "fmt"

/*
Given two singly linked lists that intersect at some point, find the intersecting node. The lists are non-cyclical.

For example, given A = 3 -> 7 -> 8 -> 10 and B = 99 -> 1 -> 8 -> 10, return the node with value 8.

In this example, assume nodes with the same value are the exact same node objects.

Do this in O(M + N) time (where M and N are the lengths of the lists) and constant space.

*/

func main() {
	// same object means same list suffix ? so we need to find common suffix then so we can start from back with 2 pointers
	// it s simple link list though, so we would want to reverse it first so we can traverse in order, but thats not efficient
	// cut the list to same length because only that part can have a matching suffix,
	// actually, dont need to cut just move the pointer there

	a := []int{2, 4, 77, 3, 7, 8, 10}
	b := []int{99, 1, 8, 10}
	ia := 0
	ib := 0
	// move indexes until tails are the same length
	for i := 0; i < len(a)-len(b); i++ {
		ia++ // next
	}
	for i := 0; i < len(b)-len(a); i++ {
		ib++ // next
	}
	// iterate the two pointers until we get the same node
	for ia < len(a) && ib < len(b) {
		if a[ia] == b[ib] {
			fmt.Printf("found intersection at %d \n", a[ia])
			return
		}
		ia++ // next
		ib++ // next
	}
	fmt.Println("No intersection")
}
