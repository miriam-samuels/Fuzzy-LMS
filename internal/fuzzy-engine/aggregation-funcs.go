package fis

// fuzzy operator and aggregation  methods
//  used normal for loops because for...range copies and the larger this array gets the more troublesome copying becomes

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
