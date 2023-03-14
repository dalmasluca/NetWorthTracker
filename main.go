package main

import (
	"os"
)

func main() {

	if _, err := os.Stat("log.log"); os.IsNotExist(err) {
		Initilization()
	}

	MainLoop()

}
