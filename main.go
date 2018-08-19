package main

import (
	"fmt"
	"os"
)

func main() {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) < 1 {
		fmt.Println("Usage: ./main <Application> [Options]")
		return
	}

	application := os.Args[1]
	
	fmt.Println(application)
}
