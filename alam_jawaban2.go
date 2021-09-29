package main

import (
	"fmt"
	"math"
)

type Crcl struct {
	n float64
}

func (r *Crcl) Luas() float64 {
	return math.Pi * r.n * r.n
}
func (r *Crcl) Keliling() float64 {
	return 2 * math.Pi * r.n
}
func main() {
	var n float64
	fmt.Printf("Input radius: ")
	fmt.Scan(&n)
	c := Crcl{n: n}
	fmt.Printf("Keliling : %.2f\n", c.Luas())
	fmt.Printf("Luas : %.2f\n", c.Keliling())
}
