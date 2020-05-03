package main

import "fmt"

//Var has package level scope.
var global_variable_1 = 54

// if not initialised , then variable has value of 0
var global_variable_int_zero_init int
//Output : I am in Foo now , printing global_variable 0

func main() {
	fmt.Println("hello Go!! world !")
	variable1 := 4
	fmt.Println(variable1)
	variable2 := 4*2
	fmt.Println(variable2)
	variable3 := "4"
	fmt.Println(variable3)
	variable4  := variable1 + variable2
	fmt.Println(variable4)
	//Finding Type of variable
	fmt.Printf("Type of variable 1 is : ")
	fmt.Printf("%T\n",variable1)
	//Type of variable 1 is : int
	foo()
	conversion_func()
}

func foo() {
	fmt.Println("I am in Foo now , printing global_variable",global_variable_1)
	fmt.Println("I am in Foo now , printing global_variable",global_variable_int_zero_init)
	// global_variable_int_zero_init = "Trying to assign string value "
	//Output : Error : .\hello_Variables.go:29:32: cannot use "Trying to assign string value " (type untyped string) as type int in assignment
}

func conversion_func()  {
	int_variable := 42
	fmt.Printf("Int Variable Details : %v\t%T\n",int_variable,int_variable)
	bool_variable :=  false
	string_variable := "3"
	fmt.Printf("String Variable Details : %v\t%T\n",string_variable,string_variable)
	// converted_variable := int(string_variable)
	fmt.Printf("Bool Variable Details : %v\t%T\n",bool_variable,bool_variable)
	// fmt.Printf("Converted Variable Details : %v\t%T\n",converted_variable,converted_variable)
	//.\hello_Variables.go:42:27: cannot convert string_variable (type string) to type int

}
