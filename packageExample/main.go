//Compile this as "go build main.go aux1.go aux2.go" or "go run main.go aux1.go aux2.go"
//After this, compile them again like this "go build main.go" or "go run main.go" and see what happens

package main

import "fmt"

func main() {
	fmt.Println("This comes from the MAIN file.")
	aux1()
	aux2()
}
