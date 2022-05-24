package main

import "fmt"

func main() {
	// 1. way of creating map
	colors := map[string]string{  
		"red": "#ff0000",
		"green": "#4bf745",
	}

	// 2. create empty map then insert keys into it later
	// var colors map[string]string
	// colors := make(map[string]string)
	
	// ** Note, cannot do .notation like structName.white - only can do with struct
	// in maps, the keys are statically typed
	colors["white"] = "#ffffff" 

	// delete keys in map
	// delete(colors, "white")

	// fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}


// map
// 1. reference type
// 2. all keys must be the same type
// 3. all values must be the same type
// 4. keys are indexed - we can iterate over them
// 5. don't need to know all the keys at compile time - can add more keys after inititation
// 6. USE TO represent a collection of related properties

// struct
// 1. value type
// 2. values can be of different type
// 3. keys don't support indexing (can't interate over them?)
// 4. you need to know all the different fields at compile time
// 5. USE TO represent a "thing" with a lot of different properties