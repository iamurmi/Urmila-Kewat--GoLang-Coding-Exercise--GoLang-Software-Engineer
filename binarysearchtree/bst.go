package binarysearchtree

import "fmt"

// Node ...
type Node struct {
	data int
	l    *Node
	r    *Node
}

// Bst ...
type Bst struct {
	root *Node
}

// derefferncing will done by the compiler automatically

// InsertRecursively ...
func (n *Node) InsertRecursively(node Node) {
	// function to find the correct place for a node in a tree

	// when new node is lesser than n -> go to left subtree of n and
	//again
	//start the function by making - left child of n ->as n
	// restore Point(Previous Execution State) stored at callstack.. it restored again when
	// newly function call finshed..
	// when new node is greater than n -> go to right subtree of n and
	//again
	//start the function by making - right child of n ->as n
	// restore Point(Previous Execution State) stored at callstack.. it restored again when
	// newly function call finshed..

	if n == nil {
		return //  base condition of Recursive Problem->    if(n==null) return ;
	} else if node.data <= n.data {
		if n.l == nil {
			n.l = &node
		} else {
			n.l.InsertRecursively(node)
		}
	} else {
		if n.r == nil {
			n.r = &node
		} else {
			n.r.InsertRecursively(node)
		}
	}
}

// Insert ...
func (bst *Bst) Insert(value int) *Bst {
	// insert function create Node then add the node to its correct position by calling insertRecursively()

	node := &Node{data: value}
	if bst.root == nil {
		bst.root = &Node{data: value}
	} else {
		bst.root.InsertRecursively(*node)
	}
	return bst
}

// PostOrderRecursively ...
func (n *Node) PostOrderRecursively() {
	if n == nil {
		return
	}
	n.l.PostOrderRecursively()
	n.r.PostOrderRecursively()
	fmt.Print(n.data, " ")
}

// PostOrder ...
func (bst *Bst) PostOrder() {
	fmt.Print("Postorder traversal: ")
	bst.root.PostOrderRecursively()
	fmt.Println("")
}

// InOrderRecursively ...
func (n *Node) InOrderRecursively() {
	if n == nil {
		return
	}
	n.l.InOrderRecursively()
	fmt.Print(n.data, " ")
	n.r.InOrderRecursively()
}

// InOrder ...
func (bst *Bst) InOrder() {
	fmt.Print("Inorder traversal: ")
	bst.root.InOrderRecursively()
	fmt.Println("")
}

// PreOrderRecursively ...
func (n *Node) PreOrderRecursively() {
	if n == nil {
		return
	}
	fmt.Print(n.data, " ")
	n.l.PreOrderRecursively()
	n.r.PreOrderRecursively()
}

// PreOrder ...
func (bst *Bst) PreOrder() {
	fmt.Print("Preorder traversal: ")
	bst.root.PreOrderRecursively()
	fmt.Println("")
}
