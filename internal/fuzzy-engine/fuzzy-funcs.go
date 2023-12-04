package fis

import (
	"math"
)

// calculate membership degree on a triangular membership func
func triangularMF(vertice TriangularMF, x float64) float64 {
	var membership float64
	if x <= vertice.A || x >= vertice.C || x == 0 {
		membership = 0
	} else if vertice.A <= x && x <= vertice.B {
		// calculate for left hand side of triangle
		distA2B := vertice.B - vertice.A // distance betwwen A and B
		distA2X := x - vertice.A         // distance betwwen A and x
		membership = distA2X / distA2B   // position of x on the lhs
	} else if vertice.B <= x && x <= vertice.C {
		// calculate for right hand side of triagle
		distB2C := vertice.C - vertice.B // distance betwwen B and C
		distX2C := vertice.C - x         // distance betwwen x and C
		membership = distX2C / distB2C   // position of x on the rhs
	}

	membership = math.Max(0, membership) // for scenarios with no membership and falls in negative value

	return membership
}

// calculate membership degree on a triangular membership func
func trapezoidalMF(vertice TrapezoidalMF, x float64) float64 {
	var membership float64
	if x < vertice.A || x > vertice.D || x == 0 {
		membership = 0
	} else if vertice.A <= x && x <= vertice.B {
		// calculate for left hand side of trapezoid
		distA2X := x - vertice.A
		distA2B := vertice.B - vertice.A
		membership = distA2X / distA2B

	} else if vertice.B <= x && x <= vertice.C {
		// the constant area of trapezoid
		membership = 1

	} else if vertice.C <= x && x <= vertice.D {
		// calculate for right hand side of trapezoid
		distC2D := vertice.D - vertice.C
		distX2D := vertice.D - x
		membership = distX2D / distC2D
	}

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
