package service

//	apply implication method
//
// we currently have 243 rules to inference so we are expecting 243 outpute
func (input *FISInput) InferenceRules() float64 {
	// variable to store output of fis
	var output []float64

	// loop through all existing rules in rule base
	for _, rule := range Rules {
		//  since we are currently only making use of the and operator we would be finding just the minimum
		//  we do this to infer creditworthiness for each rule based on given inputs
		min := minimum(
			input.Collateral[rule.Collateral],
			input.Income[rule.Income],
			input.CreditScore[rule.CreditScore],
			input.LoanAmount[rule.LoanAmount],
			input.CriminalRecord[rule.CriminalRecord],
		)

		output = append(output, min)
	}

	// aggregate output
	o := maximum(output)
	return o
}

// fuzzy operator methods
func maximum(output []float64) float64 {
	var maxValue float64
	for i := 0; i < len(output); i++ {
		if i == 0 {
			maxValue = output[i]
		} else {
			if output[i] > maxValue {
				maxValue = output[i]
			}
		}
	}
	return maxValue
}

// finf the minimum of all values passed
func minimum(output ...float64) float64 {
	var minValue float64
	for i := 0; i < len(output); i++ {
		if i == 0 {
			minValue = output[i]
		} else {
			if output[i] < minValue {
				minValue = output[i]
			}
		}
	}
	return minValue
}
