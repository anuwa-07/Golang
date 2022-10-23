
/*
	For loop in Four ways
		1. A complete, C-style for
		2. A condition-only for
		3. An infinite for
		4. for-range for
*/
package main
import (
	"fmt"
)


func main() {

	// 01.  A complete, C-style for
	for i := 0; i < 10; i++ {
		fmt.Println("Hello World > ", i)
	}

	// 02. A condition-only for
	j := 1
	for j < 10 {
		fmt.Println("Hello > ", j)
		j = j * 2
	}

	// 03. An infinite for
	/*
		You can use " break and continue " keyword to exit and continue Infinity loops
	*/
	for {
		fmt.Println("Run Forever!")
		break
	}

	// 04. For-range
	evenValus := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	for i, v := range evenValus {
		fmt.Println(i, v)
	}

}





