package main

import (
	"errors"
	"fmt"
)

// Norml funtion
func simpleFunc() {
	fmt.Println("Simple Function!")
}

// function with arguments
func argumentFuntion(name string, age int){
	fmt.Println("Name is: ", name)
	fmt.Println("Age is: ", age)
}

// funtion with return value
func returnFunc01() int {
	return 100
}

func returnFunc02(name string, hobbies []string) ( string, int) {
	fmt.Println("Hello World")
	xx := 100
	if xx >= 70 {
		return "Valid", xx
	}	
	return "Invalid value", xx
}


// variadic function:
/*
	variadic are use to have mutiple aragument, when you do not know how many will have.
	inside of funtion variadic value will convert as a slice. Then you can iterate over it and get the values.

	If you pass a slice for variadic value, you must need to add ... after the value.  Like in Below Example.
*/
func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}



// named return values.
/*
When you supply names to your return values, what you are doing is
pre-declaring variables that you use within the function to hold the
return values.
*/
func divAndReminder(numerator int, denominator int) (result int, reminder int, err error) {
	if denominator == 0 {
		err = errors.New("Can't devide from Zero!")
		return result, reminder, err
	}
	result, reminder = numerator/denominator, numerator%denominator
	return result, reminder, err
}


// Anonymous Functions
func HelloAnonymous() {
	for i := 0; i < 5; i++ {
		func (j int)  {
			fmt.Println("Inside of an Anonymous Func", j)
		}(i)
	}
}




func main() {
	fmt.Println(addTo(3))
	fmt.Println(addTo(3, 2))
	fmt.Println(addTo(3, 4, 5, 6, 7))
	fmt.Println(addTo(3, []int{1, 2, 3, 4, 5}...))

	// calling fnuc
	HelloAnonymous()
}







