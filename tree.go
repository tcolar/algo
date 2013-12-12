// History: Nov 25 13 tcolar Creation

package algo

import (
	"fmt"
	"math/rand"
)

// Binary Tree
type Tree struct {
	root *TreeNode
}

func TreeDump(t *TreeNode, indent string) (txt string) {
	txt = fmt.Sprintf("\n%s%d", indent, t.Val)

	if t.Left != nil {
		txt += TreeDump(t.Left, fmt.Sprintf("%s(%d)", indent, t.Val))
	}
	if t.Right != nil {
		txt += TreeDump(t.Right, fmt.Sprintf("%s(%d)", indent, t.Val))
	}

	return txt
}

// A (binary) tree node
type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func BuildTestBinSearchTree(min int, max int) *TreeNode {
	if (min + 1) >= max {
		return nil
	}
	mid := min + 1 + rand.Int()%(max-min-1)
	left := BuildTestBinSearchTree(min, mid)
	right := BuildTestBinSearchTree(mid, max)
	t := &TreeNode{
		Left:  left,
		Right: right,
		Val:   mid,
	}
	return t
}
