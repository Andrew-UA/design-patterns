package main

import "fmt"

func main() {
	adidasFactory, _ := GetSportFactory("adidas")
	nikeFactory, _ := GetSportFactory("nike")

	adidasShoe := adidasFactory.makeShoe()
	adidasShirt := adidasFactory.makeShirt()

	nikeShoe := nikeFactory.makeShoe()
	nikeShirt := nikeFactory.makeShirt()

	printShoeDetails(adidasShoe)
	printShirtDetails(adidasShirt)

	printShoeDetails(nikeShoe)
	printShirtDetails(nikeShirt)

}

func printShoeDetails(s IShoe) {
	fmt.Printf("Logo: %s\n", s.getLogo())
	fmt.Printf("Size: %d\n", s.getSize())
}

func printShirtDetails(s IShirt) {
	fmt.Printf("Logo: %s\n", s.getLogo())
	fmt.Printf("Size: %d\n", s.getSize())
}
