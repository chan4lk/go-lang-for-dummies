package main

import (
	"fmt"
	"strconv"
)

func convert() {
	var input string
	fmt.Print("Please enter your age: ")
	fmt.Scanf("%s", &input)
	fmt.Println("You entered: ", input)

	age, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("your age is: ", age)
	}

	b, err := strconv.ParseBool("t")
	fmt.Println(b)        // true
	fmt.Println(err)      // <nil>
	fmt.Printf("%T\n", b) // bool

	f, err := strconv.ParseFloat("3.1415", 64)
	fmt.Println(f)        // 3.1415
	fmt.Println(err)      // <nil>
	fmt.Printf("%T\n", f) // float64

	i, err := strconv.ParseInt("-18.56", 10, 64) // base 10, 64-bit

	fmt.Println(i)        // 0
	fmt.Println(err)      // strconv.ParseInt: parsing "-18.56": invalid syntax
	fmt.Printf("%T\n", i) //int64

	u1, err := strconv.ParseUint("18", 10, 64)
	fmt.Println(u1)        // 18
	fmt.Println(err)       // <nil>
	fmt.Printf("%T\n", u1) //uint64

}
