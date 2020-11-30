package main

import (
	"fmt"
	"github.com/kevinmbeaulieu/echo/echo"
	"os"
)

func main() {
	dirA := os.Args[1]
	dirB := os.Args[2]
	eq, msg := echo.DirectoriesEquivalent(dirA, dirB)
	fmt.Printf("%s & %s equivalent? %t (%s).\n", dirA, dirB, eq, msg)
}
