package main

import (
	"fmt"
)


func main() {
	fmt.Println("GO MAP")

	// make a map
	var nilMap map[string]int // this is a nill map

	// using map literal
	map_01 := map[string]int{}

	// map with valus
	animals := map[string][]string {
		"Fish": []string{"Oscar", "Arowana", "Piranha"},
		"Dog": []string{"Bingo", "Rotwiller", "Sheperd"},
	}

	// using the make()
	ages := make(map[int][]string, 10)

	fmt.Println(nilMap)
	fmt.Println(map_01)
	fmt.Println(animals)
	fmt.Println(ages)


	// reading & write values in Map
	map_03 := map[string]int{}

	map_03["Hello"] = 100
	map_03["workd"] = 0

	val, ok := map_03["Hello"]
	fmt.Println(val, ok)

	// if you try to read key not in Map, you will get Zero as default,
	vall, ok := map_03["WORLDX"]
	fmt.Println(vall, ok)






}


