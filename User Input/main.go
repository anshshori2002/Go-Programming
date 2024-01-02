package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to the user Input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the Rating for our Pizza: ")

	// comma ok || err of

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for Rating: ", input)
	fmt.Printf("Type of this Rating is %T", input)
}
