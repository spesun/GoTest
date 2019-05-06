package main

import (
	"fmt"
	"gopl.io/ch2/tempconv"
)

func main() {
	s := float64(1.0)
	fmt.Print(s)
	fmt.Printf("Brrrr! %v\n", tempconv.AbsoluteZeroC) // "Brrrr! -273.15°C"
	fmt.Println(tempconv.CToF(tempconv.BoilingC)) // "212°F"
}
