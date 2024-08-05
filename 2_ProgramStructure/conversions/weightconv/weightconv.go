// Package weightconv performs Kilogram and Pound conversions.
package weightconv

import "fmt"

type Kilogram float64
type Pound float64

func (kg Kilogram) String() string { return fmt.Sprintf("%gkg", kg) }
func (lb Pound) String() string    { return fmt.Sprintf("%glb", lb) }
