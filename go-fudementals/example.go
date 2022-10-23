
package main
import (
	"fmt"
)


func main() {
	/*
	var slice_01[]int;
	var slice_02 = []int{10, 20, 30, 40, 50, 60, 70}
	slice_03 := make([]int, 10, 20)

	fmt.Println(slice_01)
	fmt.Println(slice_02)
	fmt.Println(slice_03)

	slice_sub := slice_02[:2]
	fmt.Println(slice_sub)

	slice_sub = append(slice_sub, 100, 200, 300, 400)

	fmt.Println(slice_sub)
	fmt.Println(slice_02)
	*/

	// Copy
	slice_01 := []int{1, 2, 3, 4}
	slice_02 := make([]int, 4)

	num := copy(slice_02, slice_01)

	fmt.Println(slice_01)
	fmt.Println(slice_02)
	fmt.Println(num)

	// Appned some values to slice_02
	slice_02 = append(slice_02, 200, 300, 400, 500)
	fmt.Println(slice_02)

}




