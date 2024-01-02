package main

import "fmt"

//Capital Letter of the Word LoginToken signifies that the it is the Public Variable.
const LoginToken string = "ajhuidhi"	//Public

func main() {

	//Firstly we have to initialise the modules by using go mod variables(file name) this is good practice of initialising the modules
	var username string = "Ansh"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isloggedIn bool = true
	fmt.Println(isloggedIn)
	fmt.Printf("Variable is of type: %T \n", isloggedIn)

	var rollno uint8 = 255
	fmt.Println(rollno)
	fmt.Printf("Variable is of type: %T \n", rollno)

	var smallfloat float32 = 244.7883393849494048
	fmt.Println(smallfloat)
	fmt.Printf("Variable is of type: %T \n", smallfloat)

	//default values and some aliases
	var anotherVariables int
	fmt.Println(anotherVariables)
	fmt.Printf("Variable is of type: %T \n", anotherVariables)

	//implicit type
	var website = "www.anshshori.com"
	fmt.Println(website)

	//No var style
	//  := this operator is used insidethe function and it cannot used to initilise the global variable or in any global expression
	numberofUser := 3000.90890
	fmt.Println(numberofUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}
