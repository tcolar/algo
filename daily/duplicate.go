package main

func firstDuplicate(a []int) int {
	for _, e := range a {
		if e < 0 {
			e = -e
		}
		index := e - 1
		if a[index] < 0 {
			return e
		}
		// fmt.Print(e, index, a)
		a[index] = -a[index]
	}
	return -1
}
