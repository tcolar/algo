// History: Dec 11 13 tcolar Creation

package algo

import (
	"log"
	"math/rand"
	"testing"
	"time"
)

//4.5 Implement a function to check if a binary tree is a binary search tree
//
// Def: For every node,the node to the left of it has a less value(key) and the
// node to the right of it has a greater value(key)
func IsBinary(tree *Tree) {
	// go down from root, if NOT (left <= x > right) fail ?
	// otherwise recurse to left, right
}

func TestTree(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())

	tree := Tree{
		root: BuildTestBinSearchTree(0, 10),
	}
	log.Print(TreeDump(tree.root, ""))
}
