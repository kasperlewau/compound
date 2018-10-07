// Package compound provides a method to calculate the sum of
// an initial investment, monthly payments & compounded interest over time
package compound

import "math"

// Interest returns the sum of p + pmt/month for t years at an interest
// rate of r%. interest is compounded n times/year.
func Interest(p, pmt, r, n, t float64) float64 {
	rn := (r * .01) / n
	pow := math.Pow(1+rn, n*t)

	principal := p * pow
	series := pmt * (12 / n * (pow - 1) / rn)

	return principal + series
}
