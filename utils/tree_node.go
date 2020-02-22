package utils

import (
	"fmt"
)

type TreeNode struct {
	Name  string
	Left  *TreeNode
	Right *TreeNode
}

func (n *TreeNode) AddLeftNode(name string) *TreeNode {
	left := &TreeNode{
		Name:  name,
		Left:  nil,
		Right: nil,
	}
	n.Left = left
	return left
}

func (n *TreeNode) AddRightNode(name string) *TreeNode {
	right := &TreeNode{
		Name:  name,
		Left:  nil,
		Right: nil,
	}
	n.Right = right
	return right
}

func (n *TreeNode) Trans() {
	//前序遍历：先遍历根节点再遍历左子树，再遍历右子树
	fmt.Println(n)
	if n.Left == nil || n.Right == nil {
		return
	}
	n.Left.Trans()
	n.Right.Trans()
}
