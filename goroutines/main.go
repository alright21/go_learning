package main

import (
	"fmt"
	"time"
)


func printHello(){
	fmt.Println("Hello World")
	// time.Sleep(2 * time.Second)
	fmt.Println("Done waiting")
}

func printBye(){
	fmt.Println("Bye World")
}
// send into the channel
func sendInfo(channel chan<- string, info string){
	channel <- info
}

// extract from the channel
func printInfo(channel <-chan string){
	fmt.Println(<-channel)
}

func main(){

	// async function example
	go printHello()
	go printBye()

	information := make(chan string)
	// ordinary channels example
	go printInfo(information)
	go sendInfo(information,"Hello")


	time.Sleep(2 * time.Second)

	// buffer channel
	buffer := make(chan string, 5)

	buffer <- "Hello world"
	buffer <- "Hi"
	buffer <- "Bye"
	buffer <- "123"
	buffer <- "dsfdsa"
	// buffer <- "Deadlock"


}