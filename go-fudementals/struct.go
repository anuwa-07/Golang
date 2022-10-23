
package main
import (
	"fmt"
)


// define struct
type Person struct {
	name string
	age int
	email string
}

// anonymous structs
var employee struct {
	name string
	age int
	email string
}

func main() {
	person_01 := Person{}
	person_01.age = 100
	person_01.email = "anuruddha@gmail.com"
	person_01.name = "anuruddha"

	fmt.Println(person_01)
	fmt.Println(employee.age == 400)
	fmt.Println(employee.email == "")
}




