package main

import (
	"myappengine.io/boot"
	"fmt"
)

func showBanner(){
	fmt.Println("===========")
	fmt.Println("MyAppEngine")
	fmt.Println("===========")
}

func main() {

	showBanner()
	fmt.Println("Initializing MyAppEngine..")
	boot.InitializeAppEngine();
	fmt.Println("MyAppEngine initialized.")

}
