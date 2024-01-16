package fis

import (
	"fmt"
	"math"
)

//	apply implication method
//
// we currently have 243 rules to inference so we are expecting 243 outpute
func (input *FISInput) inference() (float64, float64, float64) {
	fmt.Printf("%+v", input)
	// variable to store fuzzy set
	set := make(map[string][]float64)

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

		set[rule.Creditworthiness] = append(set[rule.Creditworthiness], oredValues)
	}

	// merge the rule strength of each linguistic term
	rssBad := rootSumSquare(set["bad"])
	rssAvg := rootSumSquare(set["average"])
	rssGood := rootSumSquare(set["good"])

	return rssBad, rssAvg, rssGood
}
