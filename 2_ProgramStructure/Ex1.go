// Exercise 2.1: Add types, constants, and functions to tempconv for processing temperatures in
//the Kelvin scale, where zero Kelvin is −273.15°C and a difference of 1K has the same magnitude as 1°C.

package main

import (
	tempconv2 "The-Go-Programming-Language/2_ProgramStructure/conversions/tempconv"
	"fmt"
)

func main() {
	fmt.Printf("Absolute zero is %v, %v, or %v\n",
		tempconv2.AbsoluteZeroK, tempconv2.KToF(tempconv2.AbsoluteZeroK), tempconv2.KToC(tempconv2.AbsoluteZeroK))
	fmt.Printf("But the lowest registred temperature on earth is %v, %v, or %v\n",
		tempconv2.Celsius(-93.2), tempconv2.CToK(tempconv2.Celsius(-93.2)), tempconv2.KToF(tempconv2.CToK(tempconv2.Celsius(-93.2))))
}
