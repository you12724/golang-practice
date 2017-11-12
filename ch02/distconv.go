package distconv

import "fmt"

type Feet float64
type Meter float64

func (f Feet) String() string  { return fmt.Sprintf("%gf", f) }
func (m Meter) String() string { return fmt.Sprintf("%gm", m) }

func FToM(f Feet) Meter { return Meter(f / 3.28) }
func MToF(m Meter) Feet { return Feet(m * 3.28) }
