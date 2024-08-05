// Exercise 2.2: Write a general-purpose unit-conversion program analogous to cf that reads
//numbers from its command-line arguments or from the standard input if there are no arguments, and converts each number into units like temperature in Celsius and Fahrenheit,
//length in feet and meters, weight in pounds and kilograms, and the like.

// P.S. I have tried to use map[string]reflect.Type in order to make the program well-scalable, but figured out that it's
// hardly achievable without dynamic typing.
package main

import (
	"The-Go-Programming-Language/2_ProgramStructure/conversions"
	"fmt"
	"os"
	"strconv"
)

var units = []func(float64) string{conversions.Weight, conversions.Length, conversions.Temperature}

func main() {
	var input []string = os.Args[1:]
	if len(os.Args[1:]) == 0 {
		var s string
		_, err := fmt.Scanln(&s)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Ex2.2: %v\n", err)
			os.Exit(1)
		}
		input = append(input, s)
	}
	for _, arg := range input {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Ex2.2: %v\n", err)
			os.Exit(1)
		}

		for _, unit := range units {
			fmt.Println(unit(t))
		}
	}

}
