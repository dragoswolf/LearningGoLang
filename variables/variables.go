package main

import "fmt"

// We can declare global variables (at the global scope) but we cannot declare global variables using the short variable declaration
// eleven := 233 -> this will give us an error, we cannot use short variable declaration (:=)
var exEleven = 233

func main() {

	//There are various forms of declaring variables in GoLang.
	/*First, let's see the variable types:
	int -> integer type
	float64 -> float type
	bool -> boolean type
	string -> string type

	In order to declare a variable, we must first indicate that we will do such thing.
	After this, we must name our variable and indicate what type it will be.

	var integerVar int
	var floatVar float64
	var boolVar bool
	var stringVar string
	*/

	// ALL UNDECLARED VARIABLES WILL BE INITIALIZED AT 0 (OR THEIR EQUIVALENT)

	// There are some ways to declare variables. Here are a few:

	fmt.Println("Checking if the first wave of variables is correctly declared...")
	var exOne int
	var exTwo int
	var exThree float64
	var exFour bool

	fmt.Println(exOne, exTwo, exThree, exFour)

	// Now let's see another way of declaring multiple variables of the same type.

	fmt.Println("Checking if the second wave of variables is correctly declared...")
	var exFive, exSix int

	fmt.Println(exFive, exSix)

	fmt.Println("Checking if the third wave of variables is correctly declared...")

	var (
		exSeven int
		exEight float64
		exNine  bool
		exTen   string
	)

	fmt.Println(exSeven, exEight, exNine, exTen)

	// NOTE: When executed, you must've been returned with 0s, false statements and blank spaces. If so, then it is correct.

	// There is another way of declaring a variable IF WE KNOW WHAT VALUE IT WILL HAVE SINCE IT'S DECLARATION.
	// This type of declaration is called "Short Declaration Statement" or "Short Variable Declaration"

	// So far, we've seen that in order to declare a variable, we need "var <variableName> <variableType>"
	// Here is how you declare a bool without all the fuss we needed before.

	fmt.Println("Checking if the fourth wave of variables is correctly declared...")

	boolStatement := true
	intStatement := 21
	floatStatement := 23.5
	stringStatement := "Why"

	fmt.Println(boolStatement, intStatement, floatStatement, stringStatement)

	// As you can see, the ":=" eliminates the need for declaring the declaration (var) and declaring the variable type.
	// This is ONLY to be used when we initialize a variable with a value other than 0.

	// Another form of declaring variables if we already know what value it will have at its declaration is:

	fmt.Println("Checking if the fifth wave of variables is correctly declared...")

	var integerVar = 15
	var floatVar = 27.1
	var stringVar = "Hello"
	var boolVar = true

	fmt.Println(integerVar, floatVar, boolVar, stringVar)

	// We can use the Short Variable Declaration to initialize multiple parallel variables.

	fmt.Println("Checking if the sixth wave of variables is correctly declared...")

	name, lastName := "Nikola", "Tesla"

	// We can also initialize two variables of different types this way.

	success, speed := true, 50

	// As we can see, the statement from above declares a bool type variable and an integer type variable.

	fmt.Println(name, lastName, success, speed)
}
