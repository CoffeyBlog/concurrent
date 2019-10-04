
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup  				//Global WaitGroup


func main() {
	fmt.Println("OS\t", runtime.GOOS)    						//Operating system
	fmt.Println("ARCH\t", runtime.GOARCH )  					//Architecture
	fmt.Println("CPUs\t", runtime.NumCPU())  					//# of cores on CPU's
	fmt.Println("Goroutines\t", runtime.NumGoroutine()) 		//Goroutines

	wg.Add(1)
	go foo() 			// launches a single go routine
	go bar ()			// launches a second go routine

	fmt.Println("CPUs\t", runtime.NumCPU())  					//shows how many cores are on the CPU's
	fmt.Println("Goroutines\t", runtime.NumGoroutine()) 		//Goroutines how many of those ___ CPU cores are being used --- now we are using 2 of the cores on that CPU
	wg.Wait()						//Wait group function
}


func foo(){									// just a regular loop - named foo
	for i := 0; i< 10; i ++ {
		fmt.Println("foo:\t", i)
	}
	wg.Done()
}

func bar(){
	for i := 0; i < 10; i ++ {				// just a regular loop - named bar
		fmt.Println("bar:\t", i)
	}
}



// WaitGroup is a primitive "task / item" built into the language
// When developing concurrent and parallel programs -- you are like a police officer directing traffic

// First you go --- then you go --- then you can go ...etc.