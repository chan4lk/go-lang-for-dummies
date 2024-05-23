package main

import "fmt"

func scan() {
	var age int
	fmt.Print("Please enter your age: ")
	fmt.Scanf("%d", &age)
	fmt.Println("You entered: ", age)
}
