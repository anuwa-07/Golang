
package main
import (
	"fmt"
)



func main() {
	words := []string{"a", "cow", "smile", "gopher", "octopus", "anthropologist"}
	for _, word := range words {
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Println(word, "is a short word")
		case 5:
			wordLen := len(word)
			fmt.Println(word, "is exactly the right length: ", wordLen)
		case 6, 7, 8, 9:
		default:
			fmt.Println(word, "is a long word!")
		}
	}


	/*
		If you need to break from switch case in loop,
		You have to define outer scope
	*/
	
	for i := 0; i < 10; i++ {
		switch {
		case i % 2 == 0:
			fmt.Println(i, " is Even")
		case i % 3 == 0:
			fmt.Println(i, " is divisable from 3 not from 2")
		case i % 7 == 0:
			fmt.Println("Exit from the Loop!")
			break // this brek will not exit from the loop
		default:
			fmt.Println(i, " is boring!")
		}
	}

	loop:
		for i := 0; i < 10; i++ {
			switch {
			case i % 2 == 0:
				fmt.Println(i, " is Even")
			case i % 3 == 0:
				fmt.Println(i, " is divisable from 3 not from 2")
			case i % 7 == 0:
				fmt.Println("Exit from the Loop!")
				break loop // this brek will exit from the loop
			default:
				fmt.Println(i, " is boring!")
			}
		}

	
	// Blank Switch
	// In these switch case you can use confitions < , > , == , etc ...
	myWords := []string{"hi", "HelloWorld", "byee"}
	for _, word := range myWords {
		switch worldLen := len(word) {
		case worldLen < 5:
			fmt.Print(word, "is a short word!")
		case worldLen > 10:
			fmt.Println(word, "is a long word!")
		default:
			fmt.Println(word, "is exactly the ring length!")
		}
	}
}




