package main

import "fmt"

const (
    _= iota
    a=iota
    b=iota
    c=a+b
    d=c+1+iota
)

func main() {

fmt.Printf("%T\t%v\n", a,a)
fmt.Printf("%T\t%v\n", b,b)
fmt.Printf("%T\t%v\n", c,c)
fmt.Printf("%T\t%v\n", d,d)

}
