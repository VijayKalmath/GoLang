package main

import "fmt"

func main() {
	fmt.Printf("hello Go!! world ! \n")
	foo()
	fmt.Printf("Back in Main")
}

func foo() {
	fmt.Println("I am in Foo now ")
}
