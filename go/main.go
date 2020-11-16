package main

import (
	"fmt"

	"github.com/yiping-allison/intro-to-go/go/cafe"
	"github.com/yiping-allison/intro-to-go/go/person"
)

func main() {
	fmt.Println("Welcome to Go's main function!")

	fmt.Println("Starting our sample cafe...")
	cafe := cafe.Cafe{}
	// initialize our sample cafe
	cafe.Init()

	// let's create some sample customers
	fmt.Println("Fred's Turn")
	customer := person.New("Fred", 100)
	fmt.Println(customer.Buy(cafe.Drinks[0], cafe.Drinks[2], cafe.CafeFood[2]))
	fmt.Println(customer.Money)

	fmt.Println("Jess' Turn")
	lostWallet := person.New("Jess", 0)
	fmt.Println(lostWallet.Buy(cafe.CafeFood[0]))
	fmt.Println(lostWallet.Money)
}
