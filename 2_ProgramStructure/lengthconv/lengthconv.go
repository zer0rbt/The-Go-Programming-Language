// Package lengthconv performs Meter and Feet conversions.
package lengthconv

import "fmt"

type Meter float64
type Foot float64

func (m Meter) String() string { return fmt.Sprintf("%gm", m) }
func (ft Foot) String() string { return fmt.Sprintf("%gft", ft) }
