package tempconv

import (
	"fmt"
)

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZero Celsius = -273.15
	FreezingC 	 Celsius = 0
	BoillingC    Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g C", c)
}

func (c Fahrenheit) String() string {
	return fmt.Sprintf("%g C", c)
}
