package main
import (
	"fmt"
)


type People struct {
	name string
	age int
	email string
	isMarried bool
}


func ChangePerson(person People) (People, error) {
	person.email = "newemail@gmail.com"
	person.isMarried = true
	return person, nil

}

func ChangePersonPointer(person *People) {
	person.email = "new-updated-email@gmail.com"
	person.isMarried = true
}





func main() {
	person00 := People{
		name: "Anuruddha Bnadara",
		age: 25,
		email: "anuruddha@gmail.com",
		isMarried: false,
	}

	var person01 People;
	person01 = People{
		name: "Jenny",
		age: 22,
		email: "jenny@gmail.com",
		isMarried: false,
	}
// 
	fmt.Println(person00)
	fmt.Println(person01)

	/*
	person00, err00 := ChangePerson(person00)
	if err00 != nil {
		fmt.Println("Error is Happen!")
	}

	person01, err01 := ChangePerson(person01)
	if err01 != nil {
		fmt.Println("Error is Happen!")
	}
	*/
	ChangePersonPointer(&person00)
	ChangePersonPointer(&person01)

	
	fmt.Println(person00)
	fmt.Println(person01)
}






