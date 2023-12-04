package fis

import "fmt"

// fuzzify input variable income
func (i *Income) fuzzify() map[string]float64 {
	var rating float64

	// linguistic variables
	var lower float64
	var middle float64
	var upper float64

	const max float64 = 120000000 // max criteria for wealth in real world form

	// convert income level to scale of 0 - 10
	// check if user earns more than max criteria
	if i.Amount >= float64(max) {
		rating = 10 // give maximum rating
	} else {
		rating = (i.Amount * 10) / max
	}

	//get degree of membership of rating to lower class set
	vertice := TrapezoidalMF{A: 0, B: 0, C: 2, D: 3.7} // the point of each vertice
	lower = trapezoidalMF(vertice, rating)

	//get degree of membership of rating to middle class set
	vertice = TrapezoidalMF{A: 3, B: 4.583, C: 5.417, D: 7}
	middle = trapezoidalMF(vertice, rating)

	//get degree of membership of rating to upper class set
	vertice = TrapezoidalMF{A: 6.25, B: 8.5, C: 10, D: 10}
	upper = trapezoidalMF(vertice, rating)

	// return the degree to each linguistic variable set
	return map[string]float64{"lower": lower, "middle": middle, "upper": upper}
}

// method to fuzzify input variable loan amount
func (e *Employment) fuzzify() map[string]float64 {
	// linguistic variables of loan amount
	var short float64
	var medium float64
	var long float64

	//get degree of membership of rating to small
	vertice := TrapezoidalMF{A: 0, B: 0, C: 2, D: 8} // the point of each vertice
	short = trapezoidalMF(vertice, float64(e.term))

	//get degree of membership of employment term to medium employment term set
	vertice = TrapezoidalMF{A: 5, B: 10, C: 15, D: 20}
	medium = trapezoidalMF(vertice, float64(e.term))

	//get degree of membershio of employment term to large employment term set
	vertice = TrapezoidalMF{A: 15, B: 20, C: 25, D: 25}
	long = trapezoidalMF(vertice, float64(e.term))

	// return the degree to which the user belarges to each linguistic variable set
	return map[string]float64{"short": short, "medium": medium, "long": long}
}

// method to fuzzify input variable credit score
func (c *CreditScore) fuzzify() map[string]float64 {
	// linguistic variables of credit score
	var bad float64
	var average float64
	var good float64

	//get degree of membershio of score to bad credit score set
	vertice := TrapezoidalMF{A: 300, B: 300, C: 400, D: 600} // the point of each vertice
	bad = trapezoidalMF(vertice, float64(c.Score))

	//get degree of membershio of score to average credit score set
	vertice = TrapezoidalMF{A: 550, B: 650, C: 680, D: 749}
	average = trapezoidalMF(vertice, float64(c.Score))

	//get degree of membershio of score to good credit score set
	vertice = TrapezoidalMF{A: 700, B: 750, C: 850, D: 850}
	good = trapezoidalMF(vertice, float64(c.Score))

	// return the degree to which the user belarges to each linguistic variable set
	return map[string]float64{"bad": bad, "average": average, "good": good}
}

func (c *Criminal) fuzzify() map[string]float64 {
	var rating float64

	// linguistic variables of criminal record
	var bad float64
	var average float64
	var good float64

	offences := map[string]float64{
		"Homicide":                   10,
		"Theft And Robbery":          7,
		"Domestic Violence":          6,
		"Burglary":                   6,
		"Kidnapping":                 9,
		"SexualOffences":             8,
		"Drug Offences":              5,
		"Assault":                    7,
		"Cybercrime":                 6,
		"Terrorism":                  10,
		"Human Trafficking":          9,
		"Corruption And Bribery":     8,
		"Perjury":                    4,
		"Forgery And Counterfeiting": 5,
		"Other Offences":             2,
	}

	// check if user has criminal record
	if c.HasCriminalRec {
		// check then number of offences
		if len(c.Offences) > 1 {
			rating = 1
		} else {
			rating = offences[c.Offences[0]] // check rating of first item in user offences arr
		}
	} else {
		//  if user has no criminal record automatically score user good
		rating = 10
	}

	//get degree of membershio of score to bad criminal record set
	vertice := TrapezoidalMF{A: 0, B: 0, C: 1, D: 4} // the point of each vertice
	bad = trapezoidalMF(vertice, rating)

	//get degree of membershio of score to averagecriminal record set
	vertice = TrapezoidalMF{A: 2.2, B: 4.2, C: 5.8, D: 7.8}
	average = trapezoidalMF(vertice, rating)

	//get degree of membershio of score to good criminal record set
	vertice = TrapezoidalMF{A: 7, B: 9, C: 10, D: 10}
	good = trapezoidalMF(vertice, rating)

	// return the degree to which the user belarges to each linguistic variable set
	return map[string]float64{"bad": bad, "average": average, "good": good}
}

func (c *Collateral) fuzzify() map[string]float64 {
	var rating float64

	// linguistic variables of collateral
	var bad float64
	var average float64
	var good float64

	assets := map[string]float64{
		"RealEstat":              1,
		"RealEsta":               2,
		"RealEst":                3,
		"RealEs":                 4,
		"RealE":                  5,
		"RealEstate":             6,
		"Vehicles":               7,
		"SavingsOrFixedDeposits": 8,
		"StocksAndBonds":         9,
		"JewelryAndValuables":    10,
	}

	// check if user has collateral
	if c.HasCollateral {
		rating = assets[c.Collateral] // check rating of collateral

	} else {
		//  if user has no collateral automatically score user bad
		rating = 0 // TODO: add or rules to rulebase so we don't have to do this
	}
	//get degree of membershio of score to bad criminal record set
	vertice := TrapezoidalMF{A: 0, B: 0, C: 1, D: 4} // the point of each vertice
	bad = trapezoidalMF(vertice, rating)

	//get degree of membershio of score to averagecriminal record set
	verticeT := TriangularMF{A: 3, B: 5, C: 7}
	average = triangularMF(verticeT, rating)

	//get degree of membershio of score to good criminal record set
	vertice = TrapezoidalMF{A: 7, B: 9, C: 10, D: 10}
	good = trapezoidalMF(vertice, rating)

	fmt.Printf("Collateral")
	// return the degree to which the user belarges to each linguistic variable set
	return map[string]float64{"bad": bad, "average": average, "good": good}
}
