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
func IsBinary(tree *TreeNode) bool {
	// recurse down from root, if NOT (left <= x > right) fail ?
	// otherwise recurse to left, right
	if tree.Left != nil {
		if tree.Left.Val > tree.Val {
			return false
		}
		if !IsBinary(tree.Left) {
			return false
		}
	}
	if tree.Right != nil {
		if tree.Right.Val <= tree.Val {
			return false
		}
		if !IsBinary(tree.Right) {
			return false
		}
	}
	return true
}

func TestTree(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())

	tree := Tree{
		root: BuildTestBinSearchTree(0, 10),
	}
	if !IsBinary(tree.root) {
		log.Print(TreeDump(tree.root, ""))
		log.Fatal("Tree should be a binary search tree")
	}
	node := tree.root
	for node.Right != nil {
		node = node.Right
	}
	node.Val = 0
	if IsBinary(tree.root) {
		log.Print(TreeDump(tree.root, ""))
		log.Fatal("Tree should *NOT* be a binary search tree")
	}
}
