package main

import (
	"fmt"
	"sync"
)

var once sync.Once

type singleOnce struct {
}

var singleOnceInstance *singleOnce

func getSingleOnceInstance() *singleOnce {
	if singleOnceInstance == nil {
		once.Do(func() {
			fmt.Println("Creating singleOnce instance now")
			singleOnceInstance = &singleOnce{}
		})
	} else {
		fmt.Println("Single instanceOnce already created")
	}

	return singleOnceInstance
}
