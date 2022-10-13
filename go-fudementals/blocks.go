package main

import (
	"fmt"
)

func main() {
	if n := 10; n == 0 {
		fmt.Println("N is Zero!", n)
	}else if n < 5 {
		fmt.Println("N is smaill than 5", n)
	}else{
		fmt.Println("N is bigger than 5", n)
	}
}


