package main

import (
	"fmt"

	"github.com/doctordesh/goodness"
)

func main() {
	g := goodness.New("", 53)
	err := g.Serve()
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	}
}
