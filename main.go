package main

import (
	bstree "assignment/binarysearchtree"
	"assignment/houserobber"
	"fmt"
)

func bst() {
	fmt.Println("Assignment 1: ")
	v := bstree.Bst{}
	v.Insert(20)
	v.Insert(15)
	v.Insert(29)
	v.Insert(10)
	v.Insert(18)
	v.Insert(22)
	v.Insert(50)
	fmt.Println("Inserted Sequence: 20 ->  15 ->  29 ->  10 ->  18 ->  22 ->  50")

	v.InOrder()
	v.PreOrder()
	v.PostOrder()

}
func robhouse() {
	fmt.Println("")
	fmt.Println("Assignment 2: ")
	a := []int{2, 3, 4, 5, 6, 7}
	max := houserobber.Rob(a)
	fmt.Println("Array is: ", a)
	fmt.Println(" House Rober:  maximum amount of money: ", max)

}
func main() {

	bst()

	robhouse()
}
