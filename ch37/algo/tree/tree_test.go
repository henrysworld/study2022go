package tree

import (
	"fmt"
	"testing"
)

type TreeNode struct {
	Left  *TreeNode
	Val   int
	Right *TreeNode
}

type Node struct {
	Val      int
	Children []*Node
}

// 二叉树的前中后序遍历
func TestTraversal(t *testing.T) {
	// root := initNode()
	// res := inorderTraversal1(root)

	root := initNode()
	res := bstTraversal1(root)

	fmt.Println(res)
}

func inorderTraversal(root *TreeNode) []int {
	var inorder func(node *TreeNode)
	var res []int
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		res = append(res, node.Val)
		inorder(node.Left)

		inorder(node.Right)
	}

	inorder(root)
	return res
}

func inorderTraversal1(root *TreeNode) []int {
	var inorder func(node *TreeNode)
	var res []int
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		inorder(node.Left)

		inorder(node.Right)
		res = append(res, node.Val)
	}

	inorder(root)

	return res
}

func perTraversal(root *Node) []int {
	var inorder func(node *Node)
	var res []int
	inorder = func(node *Node) {
		if node == nil {
			return
		}

		res = append(res, node.Val)
		for _, cur := range node.Children {
			inorder(cur)
		}
	}

	inorder(root)

	return res
}

func bstTraversal(root *Node) [][]int {
	var ans = [][]int{}
	if root == nil {
		return ans
	}

	q := []*Node{root}
	for q != nil {
		level := []int{}
		tmp := q
		q = nil
		for _, node := range tmp {
			level = append(level, node.Val)
			q = append(q, node.Children...)
		}

		ans = append(ans, level)
	}

	return ans
}

func bstTraversal1(root *Node) [][]int {
	var ans = [][]int{}
	if root == nil {
		return ans
	}

	q := []*Node{root}
	for q != nil {
		level := []int{}
		tmp := q
		q = nil
		for _, node := range tmp {
			level = append(level, node.Val)
			q = append(q, node.Children...)
		}

		ans = append(ans, level)
	}

	return ans
}

func initBTNode() *TreeNode {
	bt := &TreeNode{}
	bt.Left = &TreeNode{}
	bt.Val = 13
	bt.Right = &TreeNode{}

	bt.Left.Val = 8

	bt.Right.Val = 18

	bt.Left.Left = &TreeNode{}

	bt.Left.Right = &TreeNode{}
	bt.Left.Left.Val = 6
	bt.Left.Right.Val = 10

	bt.Right.Left = &TreeNode{}
	bt.Right.Left.Val = 16

	bt.Right.Right = &TreeNode{}
	bt.Right.Right.Val = 20

	return bt
}

func initNode() *Node {
	root := &Node{
		Val:      0,
		Children: []*Node{},
	}

	node1 := &Node{
		Val:      1,
		Children: []*Node{},
	}
	node2 := &Node{
		Val:      2,
		Children: []*Node{},
	}
	node3 := &Node{
		Val:      3,
		Children: []*Node{},
	}

	node11 := &Node{
		Val:      11,
		Children: []*Node{},
	}

	node12 := &Node{
		Val:      12,
		Children: []*Node{},
	}

	node13 := &Node{
		Val:      13,
		Children: []*Node{},
	}

	node21 := &Node{
		Val:      21,
		Children: []*Node{},
	}

	node22 := &Node{
		Val:      22,
		Children: []*Node{},
	}

	node1.Children = append(node1.Children, node11)
	node1.Children = append(node1.Children, node12)
	node1.Children = append(node1.Children, node13)

	node2.Children = append(node2.Children, node21)
	node2.Children = append(node2.Children, node22)

	root.Children = append(root.Children, node1)
	root.Children = append(root.Children, node2)
	root.Children = append(root.Children, node3)

	return root
}
