package main
import (
	"fmt"
)



// Define a Person struct
type Person struct {
	id int32
	name string
	email string
	job string
	isMarried bool
}

// Define the method for person
func (p Person) getInfo() string {
	return fmt.Sprintf("%s %s, age %d", p.name, p.email, p.id)
}


func main() {
	var person Person;
	person = Person{
		id: 123,
		name: "Anuruddha Bandara",
		email: "bluezydev@gmail.com",
		job: "Software Engineer",
		isMarried: false,
	}

	var info string;
	info = person.getInfo();

	fmt.Println(info);
}


