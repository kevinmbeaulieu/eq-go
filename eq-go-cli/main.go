package main

import (
	"fmt"
	"github.com/kevinmbeaulieu/eq-go/eq-go"
	"os"
)

func main() {
	dirA := os.Args[1]
	dirB := os.Args[2]
	eq, msg := eqgo.DirectoriesEquivalent(dirA, dirB)
	fmt.Printf("%s & %s equivalent? %t (%s).\n", dirA, dirB, eq, msg)
}
