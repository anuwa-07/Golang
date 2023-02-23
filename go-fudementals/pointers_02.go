package main
import (
	"fmt"
)

func main()  {
	fmt.Println("THIS IS ABOUT POINTERS")

	// These are usual variables, When you pass these to functions, 
	// Function will take copy of these, Therefor the modifications doing for these variables will not affect.
	var age int32 = 23;
	var name string = "Anuruddha"

	// & : use to get the address of the variable.
	// * : Use to define variable as pointer type
	//
	// example:
	var agePointer *int32;
	agePointer = &age;

	var namePointer *string;
	namePointer = &name;

	fmt.Println("Age is", age);
	fmt.Println("Age Address is: ", agePointer);

	fmt.Println();
	fmt.Println("Name is: ", name);
	fmt.Println("Name Address is: ", namePointer);
	fmt.Println()

	// Now how we get actual value from a Pointer
	fmt.Print()
	fmt.Println(":: -------------- Get actual value from Pointer --------------- ::");
	fmt.Println("Name", *namePointer);
	fmt.Println("Age", *agePointer);
	fmt.Println();


	// You can only reassign the value if there was a value already assigned to the pointer.
	var gloable_name *string;
	fmt.Println("THE FIRST VALUE: ", gloable_name)
	fmt.Println("THE FIRST VALUE ADDRESS: ", &gloable_name);
	fmt.Println()
	updateName(gloable_name);
	fmt.Println("THE SECOND VALUE: ", gloable_name)
	fmt.Println("THE SECOND VALUE ADDRESS: ", &gloable_name);
	fmt.Println()

	// 
}

func updateName(name *string) {
	var _name string;
	_name = "Hello World";
	//
	name = &_name
}