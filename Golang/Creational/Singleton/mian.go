package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 10; i++ {
		go getInstance()
	}

	fmt.Println()

	for j := 0; j < 10; j++ {
		go getSingleOnceInstance()
	}

	fmt.Scanln()

	//time.Sleep(5 * time.Second)
}
