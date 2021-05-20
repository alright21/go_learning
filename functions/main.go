package main

import "fmt"

func main() {

	// Functions - Go Programming Tutorial #5 -> https://www.youtube.com/watch?v=tDKsqia-sYA&list=PL7yh-TELLS1FRTABHY972IpZR73rqEL5j&index=5

	hello()

	fmt.Println(add(1,4))

	fmt.Println(sum(1,4,5,7,13))

	arr := []int{1,4,5,7,6345,54,3,3}
	// use ... to treat array as single number list
	fmt.Println(sum(arr...))


	// Closures

	nextEvenNumber := evenNumbers()

	fmt.Println(nextEvenNumber())
	fmt.Println(nextEvenNumber())
	fmt.Println(nextEvenNumber())
	fmt.Println(nextEvenNumber())
	fmt.Println(nextEvenNumber())
	fmt.Println(nextEvenNumber())

}

func hello(){
	fmt.Println("Hello World")
}

func add(x int, y int) int {

	return x + y
}

// variable length parameter functions
func sum(numbers ...int) int{
	
	mysum := 0

	for _, value := range numbers{

		mysum += value
	}

	return mysum
}

// closure
func evenNumbers() func() int {
	i := 0

	return func() int {
		i+=2

		return i
	}
}