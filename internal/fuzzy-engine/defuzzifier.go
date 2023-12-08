package fis

func defuzzify(output []float64) float64 {
	// Simple defuzzification - centroid method
	numerator := 0.0
	denominator := 0.0

	//get degree of membershio of score to bad criminal record set
	vertice := TrapezoidalMF{A: 0, B: 0, C: 1, D: 4.5} // the point of each vertice
	bad := trapezoidalMF(vertice, output[0])

	//get degree of membershio of score to averagecriminal record set
	verticeT := TriangularMF{A: 3, B: 5, C: 7}
	avg := triangularMF(verticeT, output[1])

	//get degree of membershio of score to good criminal record set
	vertice = TrapezoidalMF{A: 6.5, B: 9.5, C: 10, D: 10}
	good := trapezoidalMF(vertice, output[2])

	numerator = ((output[0] * bad) + (output[1] * avg) + (output[2] * good))
	for _, value := range output {
		denominator += value
	}

	return numerator / denominator
}
