package main

import "fmt"

func main() {
	fmt.Println("Hello, welcome to the playground")
	forLoops()

}

func forLoops() {
	fmt.Println("Travers the playground")

	// where I do simple traversals
	// simple for loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("")
	dataset := []string{"a", "b", "c", "d"}
	for _, value := range dataset {
		fmt.Println(value)
	}
}
