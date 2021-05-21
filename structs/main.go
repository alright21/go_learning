package main

import "fmt"


type computer struct{
	manufacturer string
	ram int
	cpuBrand string
	laptop bool
}

// basic constructor
func newComputer(m string, r int, c string, l bool) *computer{
	cmp:= computer{manufacturer: m, ram: r, cpuBrand: c, laptop: l}

	return &cmp
}
// default constructor
func newComputerDefault() *computer{

	cmp := computer{manufacturer: "Lenovo", ram: 4, cpuBrand: "Intel", laptop: true}

	return &cmp
}

// struct methods

func (c *computer) increaseRam(amount int){

	c.ram+=amount
}

func (c *computer) printInfo(){

	fmt.Println("Printing info:")
	// ...
}

func main(){


	computer1 := computer{manufacturer: "Lenovo", ram: 16, cpuBrand: "Intel", laptop: true}

	fmt.Println(computer1)
	fmt.Println(computer1.manufacturer)

	computer2 := newComputer("Dell", 8, "AMD", true)

	fmt.Println(*computer2)

	computer1.increaseRam(16)

	fmt.Println(computer1)

}