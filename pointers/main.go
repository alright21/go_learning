package main

import "fmt"


func changeValue(val int){
	val = 200
}

func changeReference(val *int){
	*val = 500
}
func main(){

	// Pointers - Go Programming Tutorial #6 -> https://www.youtube.com/watch?v=496xSrA6QQ8&list=PL7yh-TELLS1FRTABHY972IpZR73rqEL5j&index=6

	a := 10

	fmt.Println(&a)

	ptr := &a

	fmt.Println(ptr)

	fmt.Println(*ptr)

	// pass by value
	fmt.Printf("Before changeValue: %d\n", a)
	changeValue(a)
	fmt.Printf("After changeValue: %d\n", a)

	// pass by reference
	fmt.Printf("Before changeReference: %d\n", a)
	changeReference(&a)
	fmt.Printf("After changeReference: %d\n", a)

	// using var keyword

	var ptr2 *int

	ptr2 = &a

	fmt.Println(ptr2)

	// pointer to pointer

	var pptr **int

	pptr = &ptr

	fmt.Println(pptr)
	
}