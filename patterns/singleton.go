package main

import (
	"fmt"
	"sync"
)

var (
	once sync.Once
)

func RunSingleton() {
	once.Do(func() {
		fmt.Println("Inside")
	})
	fmt.Println("Outside")
}

func main() {
	for i := 0; i < 5; i++ {
		RunSingleton()
	}
}