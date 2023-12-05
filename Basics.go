package main

import "fmt"		//Importing the formatter

func man()	{
	fmt.Println("Hello World")

	//Declaration of Variables: var name type = expression
	var integer1 int = 1
	fmt.Println(integer1)

	var a, b = 2, 3
	fmt.Println(a, b)

	var x, y = 4, "Ansh"	//It is benefit of using go Language i.e we can define multiple type of variables in single line but for this we have to initialise the element.
	fmt.Println(x, y)

	var q, w float32
	fmt.Println(q, w) 

	//Another type of declaration of Variables i.e. Short Declaration:
	//	name := expression

	d := false
	fmt.Println(d)

	// Declaration of Pointer Variables
	s := 6		
	//var s int = 6

	p := &s		
	//var p *int = &s
	
	fmt.Println(p, *p)
}
