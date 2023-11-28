package fis

import "math"

// calculate membership degree on a triangular membership func
func triangularMF(vertice TriangularMF, x float64) float64 {
	// calculate for left hand side of triagle
	distA2B := vertice.B - vertice.A // distance betwwen A and B
	distA2X := x - vertice.A         // distance betwwen A and x
	slopeLeft := distA2X / distA2B   // position of x on the lhs

	// calculate for right hand side of triagle
	distB2C := vertice.C - vertice.B // distance betwwen B and C
	distX2C := vertice.C - x         // distance betwwen x and C
	slopeRight := distX2C / distB2C  // position of x on the rhs

	// go with minimum
	membership := math.Min(slopeLeft, slopeRight)
	membership = math.Max(0, membership) // for scenarios with no membership and falls in negative value

	return membership
}

// calculate membership degree on a triangular membership func
func trapezoidalMF(vertice TrapezoidalMF, x float64) float64 {
	// calculate for left hand side of trapezoid
	distA2B := vertice.B - vertice.A
	distA2X := x - vertice.A
	slopeLeft := distA2X / distA2B

	// the constant area of trapezoid
	slopeB2C := 1

	// calculate for right hand side of trapezoid
	distC2D := vertice.D - vertice.C
	distX2D := vertice.D - x
	slopeRight := distX2D / distC2D

	membership := math.Min(math.Min(slopeLeft, float64(slopeB2C)), slopeRight)
	membership = math.Max(0, membership)

	return membership
}

// fuzzy operator and aggregation  methods
//
//	used normal for loops because for...range copies and the larger this array gets the more troublesome copying becomes
//
// find the maximum of all values passed
func maximum(val []float64) float64 {
	var maxValue float64
	for i := 0; i < len(val); i++ {
		if i == 0 {
			maxValue = val[i]
		} else {
			if val[i] > maxValue {
				maxValue = val[i]
			}
		}
	}
	return maxValue
}

// find the minimum of all values passed
func minimum(val []float64) float64 {
	var minValue float64
	for i := 0; i < len(val); i++ {
		if i == 0 {
			minValue = val[i]
		} else {
			if val[i] < minValue {
				minValue = val[i]
			}
		}
	}
	return minValue
}
