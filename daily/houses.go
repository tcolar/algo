package main

/*
A builder is looking to build a row of N houses that can be of K different colors.
He has a goal of minimizing cost while ensuring that no two neighboring houses are of the same color.

Given an N by K matrix where the nth row and kth column represents the cost to build the nth house with kth color,
return the minimum cost which achieves this goal.
*/

func main() {
	costs := [][]int{
		// 1  2  3  4 <- house number
		{3, 5, 7, 9}, // red
		{1, 5, 3, 2}, // blue
		{2, 5, 2, 3}, // green
		{9, 3, 8, 9}, // yellow
	}

	// meh don't want to do yet another DP problem, so sick of these
}
