package fis

import "fmt"

//	apply implication method
//
// we currently have 243 rules to inference so we are expecting 243 outpute
func (input *FISInput) inference() float64 {
	// variable to store fuzzy set
	var set []float64

	// loop through all existing rules in rule base
	for idx, rule := range Rules {
		//  since we are currently only making use of the 'and' operator we would be finding the minimum
		//  we do this to generate a fuzzy set
		var i []float64 = []float64{
			input.Collateral[rule.Collateral],
			input.Income[rule.Income],
			input.CreditScore[rule.CreditScore],
			input.LoanAmount[rule.LoanAmount],
			input.CriminalRecord[rule.CriminalRecord],
		}

		fmt.Printf("\n rule:: %d", idx)

		//  pass slice into aggregation function
		min := minimum(i)

		set = append(set, min)
	}

	fmt.Printf("\n Fuzzy Set %v", set)

	// aggregate output
	output := maximum(set)
	return output
}
