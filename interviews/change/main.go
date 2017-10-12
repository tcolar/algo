// To execute Go code, please declare a func main() in a package "main"

package main

import "fmt"

type coin byte
type change map[coin]int // coint -> count
var fsCoins = []coin{25, 20, 10, 5, 1}

func main() {
	change := tryChange(41, fsCoins)
	fmt.Println(change)
}

func tryChange(cents int, coins []coin) change {
	change := change{}
	coinIndex := 0

	if cents == 0 {
		return change
	}

	for _, coin := range coins {
		partialChange := tryChange(cents-int(coin), coins)
		partialChange[coin] += 1
		nbCoins := partialChange.CoinCount()
		if nbCoins > change.CoinCount() {
			change = partialChange
		}
	}

	return change
}

// Given a number of cents, write a function to make change with the
// fewest number of coins.
// using US demination: {25, 10, 5, 1}
// Output something that tells me how to break it into quarters, dimes and // nickles and cents
