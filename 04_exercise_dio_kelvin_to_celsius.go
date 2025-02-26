// simple go script to convert degrees in kelvin to celsius
package main

import (
	"fmt"
)

func main() {
	kelvin := 373.15
	celsius := kelvinToCelsius(kelvin)
	fmt.Printf("%g°K is %g°C\n", kelvin, celsius)
}

func kelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}
