// Exercise 2.4: Write a version of PopCount that counts bits by shifting its argument through 64
//bit positions, testing the rightmost bit each time. Compare its performance to the table-lookup version.

package main

import (
	"The-Go-Programming-Language/2_ProgramStructure/popcount"
	"fmt"
)

func main() {
	fmt.Println(popcount.PopCount(uint64(127)))
	fmt.Println(popcount.PopCount3(uint64(127)))
}
