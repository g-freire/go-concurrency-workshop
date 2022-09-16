package main

import (
	"fmt"
)

func Race() {
	var data int
	go func() {
		// A trying to access the variable data
		data++
		fmt.Printf("A -the value is %v. \n", data)
	}()

	// B trying to access the variable data
	//time.Sleep(time.Second)
	if data == 0 {
		fmt.Printf("B - the value is %v. \n", data)
	}
}

func main() {
	//Race()
	for i := 0; i < 1000; i++ {
		Race()
	}
}

//	RACE CONDITION FIRST EXAMPLE

//	Here, A and B are both trying to access the variable data, but there is no
//	guarantee what order this might happen in. There are three possible outcomes
//	to running this code:

//	Nothing is printed. In this case, A was executed before B.

//	“The value is 0” is printed. In this case, B were executed
//	before A.

//	“the value is 1” is printed. In this case, B was executed before A. But A was executed
//  before the print line


//var memoryAccess sync.Mutex
//var value int
//go func() {
//	memoryAccess.Lock()
//	value++
//	memoryAccess.Unlock()
//}()
//memoryAccess.Lock()
//if value == 0 {
//fmt.Printf("the value is %v.\n", value)
//} else {
//fmt.Printf("the value is %v.\n", value)
//}
//memoryAccess.Unlock()