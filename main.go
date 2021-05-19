package main

import (
	"fmt"
	"reflect"
)


func main(){
	// simple hello world in GO
	fmt.Println("Hello World")

	// variables
	var name = "Hello"
	fmt.Println(reflect.TypeOf(name))

	var number = 10
	fmt.Println(reflect.TypeOf(number))

	// initial definitio operator
	x:= 11
	fmt.Println(x)

	//constants
	const PI = 3.14 
	fmt.Println(PI)

	//booleans
	t:=true
	fmt.Println(t)

	//OPERATORS

	//logical operators
	fmt.Println(true && true)
	fmt.Println(true || false)
	fmt.Println(!true)

	//assignemt
	a:=10

	a = 20

	fmt.Println(a)

	//incerement/decrement

	a++
	a--

	fmt.Println(a)

	//aritmetic operators

	a+=1
	a-=2
	a*=3
	a/=9


	
}