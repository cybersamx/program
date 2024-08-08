package main

import (
	"fmt"

	"github.com/cybersamx/a"
	"github.com/cybersamx/b"
)

func main() {
	sample := []float64{610, 450, 160, 420, 310}

	fmt.Println(a.StdDev(sample))
	fmt.Println(b.Variance(sample))

	fmt.Println(a.Version())
	fmt.Println(b.Version())
}
