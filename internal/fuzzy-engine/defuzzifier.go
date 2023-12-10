package fis

func defuzzify(rssBad, rssAvg, rssGood float64) float64 {
	// Simple defuzzification - centroid method

	verticesBad := TrapezoidalMF{A: 0, B: 0, C: 1, D: 4.5} // the point of each vertice
	badCenter := (verticesBad.D - verticesBad.A) / 2

	verticesAvg := TriangularMF{A: 3, B: 5, C: 7}
	avgCenter := (verticesAvg.C - verticesAvg.A) / 2

	verticesGood := TrapezoidalMF{A: 6.5, B: 9.5, C: 10, D: 10}
	goodCenter := (verticesGood.D - verticesGood.A) / 2

	numerator := (badCenter * rssBad) + (avgCenter * rssAvg) + (goodCenter * rssGood)
	denominator := rssBad + rssGood + rssAvg

	return numerator / denominator
}
