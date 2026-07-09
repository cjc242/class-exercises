package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println("Hello World\n")

	input := "There once was a cat named Barry. He was a very good cat. This cat lived in Boston. He loved doing Boston-related activities (that were good for cats). He walked the esplanade. He shopped on Newbury. He ate at Tatte. He sometimes even went to TD Garden. Did you know that cats are not allowed in TD Garden?"

	kitty := []string{"c", "a", "t"}
	kat := make([]string, len(kitty))
	for i := 0; i < len(input); i++ { //want to make slice from the input to compare with slice of kitty which = "cat"
		kat[0] = input[i]

		if slices.Equal(kitty, kat) {
			fmt.Printf("found cat @ %v\n", i)
		}
	}
}
