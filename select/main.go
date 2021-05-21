package main

import (
	"fmt"
	"time"
)

func main(){

	one := make(chan string)
	two := make(chan string)

	go func(){
		time.Sleep(2 * time.Second)
		one <- "hey"
	}()

	go func(){
		time.Sleep(3 * time.Second)
		two <- "Hello"
	}()

	for x := 0; x < 2; x++ {

		select{
			case rec1 := <-one:
				fmt.Println("I received from channel 1", rec1)
			case rec2 := <- two:
				fmt.Println("I receive from channel 2", rec2)
		}
	}


}