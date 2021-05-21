package main

import (
	"errors"
	"fmt"
)

// how to retunr an error in a function using errors library
func someFunction(param int)(int, error){
	if param == 60{
		return -1, errors.New("i don't like 60")
	}else{
		return param + 10, nil
	}
}
func main(){

	var number int

	// catch the error
	_, error := fmt.Scan(&number)

	if error != nil {
		fmt.Println("We got an error")
		fmt.Println(error)
	}else{
		fmt.Println(number)
	}

	ret, error2 := someFunction(60)

	if error2 != nil {
		fmt.Println("We got an error")
		fmt.Println(error)
	}else{
		fmt.Println(ret)
	}
}