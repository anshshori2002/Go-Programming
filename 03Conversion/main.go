package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to Our app")
	fmt.Println("Please rate our pizza from 1 to 5: ")
	
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	
	fmt.Println("Thanks for Rating: ", input)
	// numRating, err := strconv.ParseFloat(input, 64)
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err!=nil {
		fmt.Println(err)

	} else {
		fmt.Println("Adding one to the rating: ", numRating+1)

	}


}
