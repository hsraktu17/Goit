package main

import (
	"fmt"
	"os"

	"github.com/hsraktu17/goit/core"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: goit <command>")
		return
	}

	switch os.Args[1] {
	case "init":
		if err := core.InitRepo(); err != nil {
			fmt.Println("Erorr:", err)
		}
	default:
		fmt.Println("Unknown command:", os.Args[1])
	}
}
