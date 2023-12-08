package fis

import (
	"fmt"
	"math"
)

//	apply implication method
//
// we currently have 243 rules to inference so we are expecting 243 outpute
func (input *FISInput) inference() []float64 {
	// variable to store fuzzy set
	inputSet := make(map[string][]float64)

	// loop through all existing rules in rule base
	for _, rule := range Rules {
		//  since we are currently only making use of the 'and' operator we would be finding the minimum
		//  we do this to generate a fuzzy set
		var i []float64 = []float64{
			input.Income[rule.Income],
			input.CreditScore[rule.CreditScore],
			input.EmploymentTerm[rule.EmploymentTerm],
			input.CriminalRecord[rule.CriminalRecord],
		}

		//  Antecedent .. minimum was used because all rules are currently ANDed
		andedValues := minimum(i) // the minimum of anded values

		// find max between anded values and collateral (we use OR because collateral is not a compulory field)
		oredValues := math.Max(input.Collateral[rule.Collateral], andedValues)

		fmt.Printf("INPUTSET:: %v \n", oredValues)

		inputSet[rule.Creditworthiness] = append(inputSet[rule.Creditworthiness], oredValues)

		// set = append(set, min)
	}

	// apply implication
	bad := maximum(inputSet["bad"])
	avg := maximum(inputSet["average"])
	good := maximum(inputSet["good"])

	outputSet := []float64{bad, avg, good}

	return outputSet
}
