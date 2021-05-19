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

	//INPUT

	var myinput int

	fmt.Println("Please enter a number: ")
	fmt.Scan(&myinput)

	fmt.Printf("Your number was %d\n", myinput)

	//CONDITIONS

	// if statements
	if myinput < 100 {
		fmt.Println("less than 100")
	}else if myinput < 200 {
		fmt.Println("less than 200")
	}else {
		fmt.Println("too large")
	}

	// switch statements
	switch myinput {
	case 10:
		fmt.Println("numer is 10")
	case 20:
		fmt.Println("number is 20")
	default:
		fmt.Println("unlucky")
	}

	// LOOPS

	// while loops
	i:=0 
	for i < 10 {
		fmt.Println(i)
		i++
	}

	// for loops
	for x:=0; x < 5; x++ {

		fmt.Println(x)
	}


	// COLLECTIONS

	// Arrays
	var arr1[5] int 
	// OR arr1 := [5]int 
	// OR arr1:= [5]int{1,2}
	

	arr1[1] = 20
	arr1[4] = 5
	fmt.Println(arr1)
	fmt.Println(arr1[4])
	//length
	fmt.Println(len(arr1))

	// 2D

	var arr2D[4][5] int

	for i:=0; i<4; i++{
		for j:=0; j<5; j++{
			arr2D[i][j] = i*j
		}
	}

	fmt.Println(arr2D)


	// slices -> lists in Python

	s := make([] int, 5)

	fmt.Println(len(s))

	s = append(s, 5)

	fmt.Println(len(s))

	// copy a slice
	c := make([]int, len(s))

	copy(c,s)

	fmt.Println(c)

	fmt.Println(c[1:4])

}