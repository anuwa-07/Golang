
package main
import (
	"fmt"
)

type Person struct {
	FirstName string
	LastName string
	Age int
}

// Pointer Reciver method
func (p *Person) Update() {
	p.FirstName = "Jhone"
	p.LastName = "Wik"
	p.Age = 45
}

// Value reciver method
func (p Person) String() string {
	return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
}


func main() {
	var prson01 Person
	prson01 = Person{
		FirstName: "Jenny",
		LastName: "Holy",
		Age: 25,
	}
	result := prson01.String()
	fmt.Println(result)
	prson01.Update()
	result01 := prson01.String()
	fmt.Println(result01)
}

