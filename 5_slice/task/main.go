// Declare a nil slice of integers. Create a loop that appends 10 values to the
// slice.
// Iterate over the slice and display each value.
// Declare a slice of five strings and initialize the slice with string literal
// values.
// Display all the elements. Take a slice of index one and two
// and display the index position and value of each element in the new slice

package main

import "fmt"

func main() {
	var numbers []int

	for i := 0; i < 10; i++ {
		numbers = append(numbers, i*10)
	}
	for id, numbers := range numbers {
		fmt.Println(id, numbers)
	}

	names := []string{"Adam", "John", "Jane", "Thomas", "Jim"}

	for i, name := range names {
		fmt.Printf("Index: %d  Name: %s\n", i, name)
	}
	// Take a slice of index 1 and 2.
	slice := names[1:3]
	// Display the value of the new slice.
	for i, name := range slice {
		fmt.Printf("Index: %d  Name: %s\n", i, name)
	}
}
