package main

import (
	"fmt"
	//"goloveu/goinactioncode"
	_ "goloveu/goinactioncode/executers"
	"goloveu/utils"
)

func main() {
	fmt.Println("Let US ROCK")
	//goinactioncode.RunSample("fdfsdfsdf")
	n := utils.TreeNode{
		Name: "root",
	}
	l1 := n.AddLeftNode("l1")
	r1 := n.AddRightNode("r1")
	_ = l1.AddLeftNode("ll2")
	_ = l1.AddRightNode("lr2")
	_ = r1.AddRightNode("rr2")

	n.Trans()
}
