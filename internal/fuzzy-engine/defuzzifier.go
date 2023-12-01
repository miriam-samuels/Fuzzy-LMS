package fis

func defuzzify(output []float64) float64 {
	// Simple defuzzification - centroid method
	numerator := 0.0
	denominator := 0.0

	for i, value := range output {
		numerator += float64(i) * value
		denominator += value
	}

	return numerator / denominator
}