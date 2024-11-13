package main

import (
	"fmt"
	"github.com/whoAbhishekSah/golang_import_poc/calculator"
	"github.com/whoAbhishekSah/golang_import_poc/internal"
)

func main() {
	fmt.Println(calculator.Add(1, 2))
	fmt.Println(internal.Add(1, 2))
	return
}
