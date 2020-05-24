package main

import "fmt"

var x bool

func main() {
	fmt.Println("=====================Playing with Bool Variables=====================")
	fmt.Println(x)
	x = true
	fmt.Println(x)
	a := 42
	b := 7
	fmt.Println("Value of a is : ", a)
	fmt.Println("Value of b is : ", b)
	fmt.Println("Is a equal than b :", a == b)
	fmt.Println("Is a greater than b :", a > b)
	fmt.Println("Is a lesser than b :", a < b)
	fmt.Println("=====================Playing with Int Variables=====================")
	var x1 uint
	x1 = 2
	x2 := 42.123
	fmt.Printf("%T\n", x1)
	fmt.Printf("%T\n", x2)
	fmt.Println(x2)
	x2 = 42
	fmt.Printf("%T\n", x2)
	var x3 complex64
	x3 = 2 + 3i
	fmt.Printf("%v\t%T\n", x3, x3)
	fmt.Println("=====================Playing with String Variables=====================")
	s := "Go Language"
	fmt.Printf("%s\n", s)
	fmt.Printf("Raw String %q\n", s)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}

	fmt.Println("")
    fmt.Println("")

	for i := 0; i < len(s); i++ {
		fmt.Printf("%q \n", s[i])
	}

    fmt.Println("")
	fmt.Println("")
	for i, v := range s {
		fmt.Printf("%d \t %d \n",  v, i)
	}
}
