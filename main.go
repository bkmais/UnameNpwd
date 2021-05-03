package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: [Feet]")
		return
	}

	feet := os.Args[1]
	val, err := strconv.ParseFloat(feet, 64)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	mtrs := val * 0.3048
	fmt.Printf("%v feet = %g meters.\n", feet, mtrs)
}
