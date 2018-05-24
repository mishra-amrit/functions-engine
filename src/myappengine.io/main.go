package main

import (
	. "myappengine.io/app"
	. "myappengine.io/util"
)

func showBanner() {
	Log("===========")
	Log("MyAppEngine")
	Log("===========")
}

func main() {

	showBanner()
	Log("Initializing MyAppEngine..")
	Initialize()
	Log("MyAppEngine initialized.")

}
