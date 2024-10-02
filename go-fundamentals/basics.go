package main

import "fmt"

func main() {
	// Basic logic: if-else
	fmt.Println("If conditional:")
	x := 10
	if x > 5 {
		fmt.Println("x is greater than 5")
	} else if x == 5 {
		fmt.Println("x is equal to 5")
	} else {
		fmt.Println("x is less than 5")
	}

	// For loop
	fmt.Println("For loop:")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// While-like loop (Go doesn't have an explicit while loop, but for can act like one)
	fmt.Println("While-like loop:")
	y := 0
	for y < 5 {
		fmt.Println(y)
		y++
	}

	// Range loop for iterating over a slice
	fmt.Println("Range loop over a slice:")
	numbers := []int{1, 2, 3, 4, 5}
	for i, num := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", i, num)
	}

	// Using break and continue in a loop
	fmt.Println("Break and continue example:")
	for i := 0; i < 10; i++ {
		if i == 3 {
			continue // Skip the rest of the loop when i is 3
		}
		if i == 7 {
			break // Stop the loop when i is 7
		}
		fmt.Println(i)
	}
}
