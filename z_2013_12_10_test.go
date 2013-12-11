// History: Dec '▩'' ' '▩'3 tcolar Creation

package algo

import (
	"log"
	"testing"
)

var Maze [][]int = [][]int{
	{'▩', '▩', '▩', ' ', '▩', '▩', '▩', '▩', '▩', '▩', '▩', '▩', '▩', '▩', '▩'},
	{'▩', ' ', '▩', ' ', '▩', '▩', ' ', ' ', ' ', '▩', '▩', '▩', '▩', ' ', ' '},
	{'▩', ' ', ' ', ' ', ' ', '▩', ' ', ' ', ' ', '▩', '▩', '▩', ' ', ' ', '▩'},
	{'▩', ' ', ' ', ' ', ' ', '▩', '▩', ' ', ' ', ' ', ' ', ' ', ' ', ' ', '▩'},
	{'▩', ' ', '▩', '▩', ' ', '▩', '▩', ' ', '▩', '▩', ' ', '▩', '▩', '▩', '▩'},
	{'▩', ' ', '▩', '▩', ' ', '▩', '▩', ' ', '▩', '▩', ' ', ' ', ' ', '▩', '▩'},
	{' ', ' ', '▩', '▩', ' ', '▩', '▩', ' ', '▩', '▩', '▩', '▩', '▩', '▩', '▩'},
	{'▩', ' ', '▩', ' ', ' ', ' ', ' ', ' ', '▩', '▩', '▩', '▩', '▩', '▩', '▩'},
	{'▩', ' ', '▩', ' ', ' ', '▩', '▩', ' ', '▩', '▩', '▩', '▩', '▩', '▩', '▩'},
	{'▩', '▩', '▩', '▩', '▩', '▩', '▩', ' ', '▩', '▩', '▩', '▩', '▩', '▩', '▩'},
}

func TestMaze(t *testing.T) {
	startx, starty := 0, 6
	endx, endy := 14, 1
	robot := Robot{
		X:       startx,
		Y:       starty,
		Heading: 'E',
	}
	maze := Maze
	for robot.X != endx || robot.Y != endy {
		robot.move(maze)

		// If back at starting point we failed to find an exit
		if robot.X == startx && robot.Y == starty {
			log.Fatal("Could not find exit.")
		}
	}
	// Draw the maze
	for y := 0; y != len(maze); y++ {
		txt := ""
		for x := 0; x != len(maze[y]); x++ {
			txt += string(maze[y][x])
		}
		txt += string('\n')
		log.Print(txt)
	}
}

type Robot struct {
	X       int
	Y       int // cur pos in maze
	Heading int
}

func (r *Robot) move(maze [][]int) {
	// Attempt to move the robot
	// Trick : Following the right wall should find the exit
	// Actually will follow left wall as more eficient overall with a right turn only robot

	// choice 1 is to the left
	r.Right()
	r.Right()
	r.Right()
	if r.Forward(maze) {
		return
	}
	// choice 2 is straight
	r.Right()
	if r.Forward(maze) {
		return
	}
	// choice 3 is to the right
	r.Right()
	if r.Forward(maze) {
		return
	}
	// choice 4 is turn around
	r.Right()
	if r.Forward(maze) {
		return
	}
}

// Turn the robot to the right (of current heading.)
func (r *Robot) Right() {
	switch r.Heading {
	case 'N':
		r.Heading = 'E'
	case 'E':
		r.Heading = 'S'
	case 'S':
		r.Heading = 'W'
	case 'W':
		r.Heading = 'N'
	}
}

// Move forward if it can(no wall)
// Return wether it was availabe to move or not
func (r *Robot) Forward(maze [][]int) bool {
	x, y := r.X, r.Y
	switch r.Heading {
	case 'N':
		y--
	case 'E':
		x++
	case 'S':
		y++
	case 'W':
		x--
	}

	//log.Printf("Checking %d, %d", x, y)

	if x < 0 || y < 0 || y >= len(maze) || x >= len(maze[0]) {
		return false // can't go out of maze
	}

	if maze[y][x] != '▩' { // open space, move to it
		log.Printf("Moved %s to %d, %d", string(r.Heading), x, y)
		r.X = x
		r.Y = y
		maze[y][x] = '*'
		return true
	}

	// could not move
	return false
}
