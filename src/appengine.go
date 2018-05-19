package main

import (
	"boot"
	"fmt"
)

func main() {
	fmt.Println("===========")
	fmt.Println("MyAppEngine")
	fmt.Println("===========")
	fmt.Println("\nInitializing MyAppEngine..")
	boot.InitializeAppEngine();
	fmt.Println("MyAppEngine initialized.")

}
