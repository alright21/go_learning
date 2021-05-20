package main

import "fmt"


func main() {

	// Range - Go Programming Tutorial #4 -> https://www.youtube.com/watch?v=iIiqoOa4IzE&list=PL7yh-TELLS1FRTABHY972IpZR73rqEL5j&index=4

	numbers:= []int{1,45,3,67,8,3,90,54}

	for index, number := range numbers {
		fmt.Printf("value at index %d is %d\n", index, number)
	}

	// without index
	for _, number := range numbers {
		fmt.Printf("Value without index: %d\n", number)
	}

	// with maps

	mymap:= map[string]int{"A": 1, "B": 3, "C": 5}

	for key, value := range mymap {
		fmt.Printf("Key %s has value %d\n", key, value)
	}
}