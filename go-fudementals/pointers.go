
package main
import (
	"fmt"
)

func failedUpdate(g *int) {
	x := 100
	g = &x
}

type Foo struct {
	field01 string
	field02 int
}

func MakeFoo01(foo *Foo) error {
	foo.field01 = "Hello World"
	foo.field02 = 10000
	return nil
}

func MakeFoo02() (Foo, error) {
	foo := Foo{
		field01: "Hello Hello",
		field02: 1010,
	}
	return foo, nil
}


func main() {
	/*
		& - Is use to get the Address of a variable
		* - Is use to get the value from a pointer
	*/
	x := 100
	pointerX := &x

	// fmt.Println(x)
	fmt.Println(pointerX)
	// fmt.Println(*pointerX) // This is called dereferencing

	// Before dereferencing must make sure that the pointer is non-nil.

	// Important Stuff
	/*
	Even we pass pointer to a Function, We can't change that pointer's address.
	We only can change that pointer's value only. 
	*/

	var f *int // f is nill
	failedUpdate(f)
	// fmt.Println(f) // prints nill


	var xx Foo;
	MakeFoo01(&xx)
	fmt.Println(xx)



}






