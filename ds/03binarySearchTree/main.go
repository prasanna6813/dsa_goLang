package main

import "fmt"

type Node struct{
	key int
	Left *Node
	Right *Node
}

var count int

// Insert will add a node to the tree
func (n *Node) Insert(k int){
	
	if n.key < k{
		// move right
		if n.Right == nil{
			n.Right = &Node{key: k}
		} else {
			n.Right.Insert(k)
		}
	}else if n.key > k{
		// move left
		if n.Left == nil{
			n.Left = &Node{key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}


// Search will take in a key value
// and return true if there is a node with that value

func(n *Node) Search(k int) bool{
	count++ 
	if n == nil{
		return false
	}
	if n.key < k{
		//  move right
		return n.Right.Search(k)
	} else if n.key > k {
		//  move left
		n.Left.Search(k)
	}
	return true
}

func main(){
	tree:=&Node{key: 100}
	tree.Insert(200)
	tree.Insert(300)
	fmt.Println(tree)

	tree.Insert(52)
	tree.Insert(203)
	tree.Insert(19)
	tree.Insert(76)
	tree.Insert(150)
	tree.Insert(310)
	tree.Insert(7)
	tree.Insert(24)
	tree.Insert(88)
	tree.Insert(276)

	fmt.Println(tree.Search(7656),"\n",count )

}