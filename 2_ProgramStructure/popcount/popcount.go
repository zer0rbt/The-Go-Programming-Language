//Exercise 2.3: Rewrite PopCount to use a loop instead of a single expression. Compare the performance of the two
// versions. (Section 11.4 shows how to compare the performance of different implementations systematically.)

package popcount

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// PopCount2 returns the population count (number of set bits) of x, but works different from PopCount.
func PopCount2(x uint64) int {
	var out byte
	for i := 0; i < 8; i++ {
		out += pc[byte(x>>(i*8))]
	}
	return int(out)
}
