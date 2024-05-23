package main

import (
	"fmt"
	"reflect"
	"time"
	"unicode/utf8"
)

const publisher = "Wiley"

func main() {
	var num1 = 6     // type infered
	var num2 int     // explicity type
	var num3 float32 // floating point variable
	var name string  // string variable
	var raining bool // boolean variable

	firstName := "Chandima"
	first, last, age := "chandima", "ranaweera", 33

	var (
		middleName  string = "Sri"
		lastName    string = "Manoj"
		ageLastYear        = 32
	)

	var unused = 5
	_ = unused // compiler is happy

	adress := "The White House\n1600 Pennsylvania Avenue NW\nWashington, DC 20500"

	quote := `"Anyone who has never made
a mistake has never tried
anything new."
--Albert Einstein`

	firstNameSinhala := "චන්දිම"
	lastNameSinhala := "රණවීර"

	start := time.Now()

	fmt.Println(num1)        // 6
	fmt.Println(num2)        // 0
	fmt.Println(num3)        // 0
	fmt.Println(name)        // "" (empty string)
	fmt.Println(raining)     // false
	fmt.Println(firstName)   // Chandima
	fmt.Println(first)       // chandima
	fmt.Println(last)        // ranaweera
	fmt.Println(age)         // 33
	fmt.Println(middleName)  // Sri
	fmt.Println(lastName)    // Manoj
	fmt.Println(ageLastYear) // 32
	fmt.Println(publisher)   // Wiley
	fmt.Println(adress)
	fmt.Println(quote)
	fmt.Println(firstNameSinhala)                         // "චන්දිම"
	fmt.Println(lastNameSinhala)                          //"රණවීර"
	fmt.Println(len(firstNameSinhala))                    // 18
	fmt.Println(len(lastNameSinhala))                     // 15
	fmt.Println(utf8.RuneCountInString(firstNameSinhala)) //6
	fmt.Println(utf8.RuneCountInString(lastNameSinhala))  //5
	fmt.Printf("%T\n", start)                             // time.Time
	fmt.Println(reflect.TypeOf(start))                    // time.Time
	fmt.Println(reflect.ValueOf(start).Kind())            // struct
}
