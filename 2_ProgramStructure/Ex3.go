//Exercise 2.3: Rewrite PopCount to use a loop instead of a single expression. Compare the performance of the two
// versions. (Section 11.4 shows how to compare the performance of different implementations systematically.)

package main

import (
	"The-Go-Programming-Language/2_ProgramStructure/popcount"
	"fmt"
)

func main() {
	fmt.Println(popcount.PopCount(uint64(129)))
	fmt.Println(popcount.PopCount2(uint64(129)))
}
