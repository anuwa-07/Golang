

package main
import (
	"fmt"
)

/*
	Arrays in GO
		Array are not commonly used in GO. 
		Onece you define array you will not able to extend it or append it any value.
		Go considers the size of the array to be part of the type of the array.
		There for an array that's declared to be [3]int a different type from an array that’s declared to be [4]int.
		You can’t use a type conversion to convert arrays of different sizes to identical types.
		You can’t write a function that works with arrays of any size and you can’t assign arrays of different sizes to the same variable.

	How to define
		
		* without values
		var arr [3]int
			- len() = 3
			- arr[0] = arr[1] = arr[2] = 0
		
		* with values
		var arr = [3]int{1, 2, 3}
		var arr = [...]int{1, 2, 3}

	Compaire array:
		==, != are can be used with arrays.
		ex:
			arr_01 == arr_02 
			arr_01 != arr_02

	
	Slice in GO
		Slice will use commonly in GO.
		Unlike arrays, length is not a part of a slice type
	
	How to define
		* without values
			this will make a "nil" slice.
		var slice_01 []int;
		
		* with values
		var slice_01 = []int{10, 20, 30}
	
	append() 
		

	
*/

func main() {
	var arr [3]int
	fmt.Println(len(arr))
	fmt.Println("Hello World!")
}












