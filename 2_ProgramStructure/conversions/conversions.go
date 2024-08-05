package conversions

import (
	lengthconv2 "The-Go-Programming-Language/2_ProgramStructure/conversions/lengthconv"
	tempconv2 "The-Go-Programming-Language/2_ProgramStructure/conversions/tempconv"
	weightconv2 "The-Go-Programming-Language/2_ProgramStructure/conversions/weightconv"
	"fmt"
)

// basicFormat receives 5 convertable to string arguments and returns nice-looking string to print.
func basicFormat(name string, v1, v2, v3, v4 any) string {
	return fmt.Sprintf("%v:\n %v = %v, while %v = %v.", name, v1, v2, v3, v4)
}

// Weight receives value and converts kilograms to pounds and back.
func Weight(v float64) string {
	return basicFormat("Weight",
		weightconv2.Kilogram(v), weightconv2.KgToLb(weightconv2.Kilogram(v)),
		weightconv2.Pound(v), weightconv2.LbToKg(weightconv2.Pound(v)))
}

// Length receives value and converts meters to feet and back.
func Length(v float64) string {
	return basicFormat("Length",
		lengthconv2.Meter(v), lengthconv2.MToFt(lengthconv2.Meter(v)),
		lengthconv2.Foot(v), lengthconv2.FtToM(lengthconv2.Foot(v)))
}

// Temperature receives value and converts Celsius to Fahrenheit and back.
func Temperature(v float64) string {
	return basicFormat("Temperature",
		tempconv2.Celsius(v), tempconv2.CToF(tempconv2.Celsius(v)),
		tempconv2.Fahrenheit(v), tempconv2.FToC(tempconv2.Fahrenheit(v)))
}
