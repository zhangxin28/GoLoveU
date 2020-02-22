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

func TransNodeDLR(root *TreeNode) {
	//前序遍历：先遍历根节点再遍历左子树，再遍历右子树
	if root == nil {
		return
	}
	fmt.Println(root)
	TransNodeDLR(root.Left)
	TransNodeDLR(root.Right)
}

func TransNodeLDR(root *TreeNode) {
	//中序遍历:先遍历左子树再遍历根，再遍历根节点再遍历右子树
	if root == nil {
		return
	}
	TransNodeLDR(root.Left)
	fmt.Println(root)
	TransNodeLDR(root.Right)
}

func TransNodeLRD(root *TreeNode) {
	//后序遍历:先遍历左子树再遍历根，再遍历右子树，再遍历根节点
	if root == nil {
		return
	}
	TransNodeLRD(root.Left)
	TransNodeLRD(root.Right)
	fmt.Println(root)
}
