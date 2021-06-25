package main

import (
	"fmt"
)

/*
You are given an M by N matrix consisting of booleans that represents a board.
Each True boolean represents a wall. Each False boolean represents a tile you can walk on.

Given this matrix, a start coordinate, and an end coordinate, return the minimum number of
steps required to reach the end coordinate from the start. If there is no possible path, then return null.
You can move up, left, down, and right. You cannot move through walls. You cannot wrap around the edges of the board.
*/

func main() {
	b := board{
		{false, false, false, false},
		{true, true, false, true},
		{false, false, false, false},
		{false, false, false, false},
	}
	fmt.Println(movesTo(b, coords{3, 0}, coords{0, 0}))
}

var queue []coords

func movesTo(b board, from, to coords) int {
	// sounds like dfs or Dijkstra, but i only remember how to do dfs, so using that
	depth := 0
	queue = append(queue, from)
	for len(queue) > 0 {
		ql := len(queue)
		for i := 0; i != ql; i++ {
			at := queue[i]
			if at.x == to.x && at.y == to.y {
				return depth
			}
			// mark as a wall so we wont revisit, but that mutates board so kinda ugly
			b[at.x][at.y] = true
			// enqueue all non wall neighbors (no diagonals)
			enqueue(b, at, -1, 0)
			enqueue(b, at, 0, -1)
			enqueue(b, at, 0, +1)
			enqueue(b, at, 1, 0)
		}
		queue = queue[ql:]
		fmt.Println(queue)
		depth++
	}
	return -1
}

func enqueue(b board, pos coords, xoffset, yoffset int) {
	tox := pos.x + xoffset
	toy := pos.y + yoffset
	if tox >= len(b) || tox < 0 {
		return
	}
	if toy >= len(b[0]) || toy < 0 {
		return
	}
	if b[tox][toy] {
		return // wall
	}
	at := coords{tox, toy}
	queue = append(queue, at)
}

type coords struct {
	x int
	y int
}
type board [][]bool
