/*
	Methods for nil Instances

	GO allows you to call a method on a nil instance
	- It actually tries to invoke the method.
	- If it’s a method with a value receiver, you’ll get a panic.
	- If it’s a method with a pointer receiver,
	  it can work if the method is written to handle the possibility of a nil instance.

	Example: Implementaion of Binary Tree
*/

package main

import (
	"fmt"
)

type IntTree struct{
	val int
	left, right *IntTree
}

// Insert a value
func (it *IntTree) Insert(val int) *IntTree {
	if it == nil {
		return &IntTree{val: val}
	}
	if val < it.val {
		it.left = it.left.Insert(val)
	} else if val > it.val {
		it.right = it.right.Insert(val)
	}
	//
	return it
}

// check the avalibilty
func (it *IntTree) Contains(val int) bool {
	switch {
	case it == nil:
		return false
	case val < it.val:
		return it.left.Contains(val)
	case val > it.val:
		return it.right.Contains(val)
	default:
		return true
	}
}

// 
type MySelf struct {
	name string
	email string
}
func (ms MySelf) getName() string {
	return ms.name;
}


// Start Calling!
func main() {
	fmt.Println("Implement Binary Tree!")
	var it *IntTree;
	it = it.Insert(5)
	it = it.Insert(15)
	it = it.Insert(25)
	it = it.Insert(35)
	it = it.Insert(45)
	fmt.Println(it.Contains(1200))
	fmt.Println(it.Contains(45))

	// You can call methos also with MainType, But when you calling it, 
	// you need to pass Instance as fist argument.
	// But this valid only non pointer methods.
	me00 := MySelf{
		name: "Anuruddha Bnadara",
		email: "anurddha@codimite.com",
	}
	getNameExample := me00.getName
	name := getNameExample(me00)
	fmt.Println(name)
}




