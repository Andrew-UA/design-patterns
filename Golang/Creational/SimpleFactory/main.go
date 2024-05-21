package main

import "fmt"

func main() {
	ak47, _ := getGun("ak47")
	m16, _ := getGun("m16")

	printDetails(ak47)
	printDetails(m16)
}

func printDetails(g IGun) {
	fmt.Printf("Gun: %s \n", g.getName())
	fmt.Printf("Power: %d \n", g.getPower())
}
