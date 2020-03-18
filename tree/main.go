package main

import (
	"fmt"
	"math"
)

type treeNode struct {
	Val   int
	Left  *treeNode
	Right *treeNode
}

func main() {
	root := treeNode{Val: 5, Left: nil, Right: nil}

	root.addLeft(4)
	root.addLeft(11)
	root.addLeft(7)
	root.Left.Left.addRight(2)
	root.addRight(8)
	root.addRight(4)
	root.Right.addLeft(13)
	root.addRight(1)

	fmt.Println(hasPathSum(&root, 22))
}

func (tn *treeNode) add(value int) {
	if tn.Left == nil {
		tn.Left = &treeNode{Val: value, Left: nil, Right: nil}
		return
	} else if tn.Right == nil {
		tn.Right = &treeNode{Val: value, Left: nil, Right: nil}
		return
	}

	if tn.Left != nil {
		tn.Left.add(value)
	} else if tn.Right != nil {
		tn.Right.add(value)
	}
}

func (tn *treeNode) addLeft(value int) {
	if tn.Left != nil {
		tn.Left.addLeft(value)
		return
	}

	tn.Left = &treeNode{Val: value, Left: nil, Right: nil}
}

func (tn *treeNode) addRight(value int) {
	if tn.Right != nil {
		tn.Right.addRight(value)
		return
	}

	tn.Right = &treeNode{Val: value, Left: nil, Right: nil}
}

func maxDepth(root *treeNode) int {
	if root == nil {
		return 0
	}
	depth := 1
	depth = int(math.Max(float64(depth), float64(maxDepth(root.Left)+1)))
	depth = int(math.Max(float64(depth), float64(maxDepth(root.Right)+1)))

	return depth
}

func hasPathSum(root *treeNode, sum int) bool {
    if root == nil {
        return false
	}
	
	if root.Left == nil && root.Right == nil {
		return sum == root.Val
	}

	if hasPathSum(root.Left, sum - root.Val) {
		return true
	}

	if hasPathSum(root.Right, sum - root.Val) {
		return true
	}

    return false
}
