// Exercise 2.1: Add types, constants, and functions to tempconv for processing temperatures in
//the Kelvin scale, where zero Kelvin is −273.15°C and a difference of 1K has the same magnitude as 1°C.

package main

import (
	"The-Go-Programming-Language/2_ProgramStructure/tempconv"
	"fmt"
)

func main() {
	fmt.Printf("Absolute zero is %v, %v, or %v\n",
		tempconv.AbsoluteZeroK, tempconv.KToF(tempconv.AbsoluteZeroK), tempconv.KToC(tempconv.AbsoluteZeroK))
	fmt.Printf("But the lowest registred temperature on earth is %v, %v, or %v\n",
		tempconv.Celsius(-93.2), tempconv.CToK(tempconv.Celsius(-93.2)), tempconv.KToF(tempconv.CToK(tempconv.Celsius(-93.2))))
}
