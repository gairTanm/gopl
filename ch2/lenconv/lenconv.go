package lenconv

import "fmt"

type Feet float64
type Meter float64

func (m Meter) String() string { return fmt.Sprintf("%g meters", m) }
func (f Feet) String() string  { return fmt.Sprintf("%g ft", f) }
