package main

import (
	"fmt"
)

func main() {
	var p portoflio

	p = p.CreatePortoflio()

	p.save("file/p.json")
	fmt.Printf("p: %v\n", p)

}
