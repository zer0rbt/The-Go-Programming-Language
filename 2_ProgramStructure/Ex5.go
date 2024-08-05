// Exercise 2.5: The expression x&(x-1) clears the rightmost non-zero bit of x. Write a version
// of PopCount that counts bits by using this fact, and assess its performance.

package main

import (
	"The-Go-Programming-Language/2_ProgramStructure/popcount"
	"fmt"
)

func main() {
	fmt.Println(popcount.PopCount(uint64(127)))
	fmt.Println(popcount.PopCount4(uint64(127)))
}
