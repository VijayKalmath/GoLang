// Hands-on exercise #4
// For this exercise
// Create your own type. Have the underlying type be an int.
// create a VARIABLE of your new TYPE with the IDENTIFIER “x” using the “VAR” keyword
// in func main
// print out the value of the variable “x”
// print out the type of the variable “x”
// assign 42 to the VARIABLE “x” using the “=” OPERATOR
// print out the value of the variable “x”
package main

import "fmt"

type newType int

var  x newType

func main() {
    fmt.Printf("Value of variable x %v\n",x)
    fmt.Printf("Type of variable x %T\n",x)
    x=42
    fmt.Printf("New Value of variable x %v\n",x)

}
