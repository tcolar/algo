// History: Nov 25 13 tcolar Creation

package algo

import ()

// Binary Tree
type Tree struct {
	root *TreeNode
}

// A (binary) tree node
type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   interface{}
}
