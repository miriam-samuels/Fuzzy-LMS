package service

// function to calculate the degree to which rating conforms to a subset x
//
//	where 0 is no membership
//	1 is full membership
//	and  rating /  maxValue get the exact degree
func calcMembershipDegree(rating float64, minValue float64, maxValue float64) float64 {
	var degree float64
	if rating < minValue {
		degree = 0
	} else if (rating > minValue) && (rating < maxValue) {
		degree = rating / maxValue
	} else {
		degree = 1
	}
	return degree
}

// method to fuzzify input variable income
func (i *Income) Fuzzify() map[string]float64 {
	// var to store variable value for fuzzification
	var rating float64

	// linguistic variables
	var lower float64
	var middle float64
	var upper float64

	// min and max of liguistic variables
	minLower := 0.0
	maxLower := 4.0
	minMiddle := 3.0
	maxMiddle := 7.0
	minUpper := 6.4
	maxUpper := 10.0

	const max float64 = 120000000 // max criteria for wealth in real world form

	// convert income level to scale of 0 - 10
	// check if user earns more than max criteria
	if i.amount >= float64(max) {
		rating = 10 // give maximum rating
	} else {
		rating = (i.amount * 10) / max
	}

	// fuzify data
	//get degree of membershio of rating to lower class set
	lower = calcMembershipDegree(rating, minLower, maxLower)

	//get degree of membershio of rating to middle class set
	middle = calcMembershipDegree(rating, minMiddle, maxMiddle)

	//get degree of membershio of rating to middle class set
	upper = calcMembershipDegree(rating, minUpper, maxUpper)

	// return the degree to which the user belarges to each linguistic variable set
	return map[string]float64{"lower": lower, "middle": middle, "upper": upper}
}

// method to fuzzify input variable EMPLOYMENT TERM
func (a *LoanAmount) Fuzzify() map[string]float64 {
	// var to store variable value for fuzzification
	var rating float64

	// linguistic variables of employment period
	var small float64
	var medium float64
	var large float64

	// min and max of liguistic variables (years)
	minSmall := 0.0
	maxSmall := 9.0
	minMedium := 5.0
	maxMedium := 20.0
	minLarge := 15.0
	maxLarge := 25.0

	const max float64 = 200000000 // max amount to borrow

	// convert income level to scale of 0 - 10
	// check if user earns more than max criteria
	if a.amount >= float64(max) {
		rating = 10 // give maximum rating
	} else {
		rating = (a.amount * 10) / max
	}

	//get degree of membershio of employment term to small employment term set
	small = calcMembershipDegree(rating, minSmall, maxSmall)

	//get degree of membershio of employment term to medium employment term set
	medium = calcMembershipDegree(rating, minMedium, maxMedium)

	//get degree of membershio of employment term to large employment term set
	large = calcMembershipDegree(rating, minLarge, maxLarge)

	// return the degree to which the user belarges to each linguistic variable set
	return map[string]float64{"small": small, "medium": medium, "large": large}
}

// method to fuzzify input variable credit score
func (c *CreditScore) Fuzzify() map[string]float64 {
	// linguistic variables of credit score
	var bad float64
	var average float64
	var good float64

	// min and max of liguistic variables (uisng PICO scoring system)
	minBad := 300.0
	maxBad := 600.0
	minAverage := 550.0
	maxAverage := 750.0
	minGood := 700.0
	maxGood := 850.0

	//get degree of membershio of score to bad credit score set
	bad = calcMembershipDegree(float64(c.score), minBad, maxBad)

	//get degree of membershio of score to average credit score set
	average = calcMembershipDegree(float64(c.score), minAverage, maxAverage)

	//get degree of membershio of score to good credit score set
	good = calcMembershipDegree(float64(c.score), minGood, maxGood)

	// return the degree to which the user belarges to each linguistic variable set
	return map[string]float64{"bad": bad, "average": average, "good": good}
}

func (c *Criminal) Fuzzify() map[string]float64 {
	// var to store variable value for fuzzification
	var rating float64

	// linguistic variables of criminal record
	var bad float64
	var average float64
	var good float64

	// min and max of liguistic variables (uisng PICO scoring system)
	minBad := 0.0
	maxBad := 4.0
	minAverage := 2.3
	maxAverage := 7.0
	minGood := 7.0
	maxGood := 10.0

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
	if c.hasCriminalRec {
		// check then number of offences
		if len(c.Offences) > 1 {
			rating = 0
		} else {
			rating = offences[c.Offences[0]] // check rating of first item in user offences arr
		}
	} else {
		//  if user has no criminal record automatically score user good
		rating = 10
	}

	//get degree of membershio of score to bad criminal record set
	bad = calcMembershipDegree(rating, minBad, maxBad)

	//get degree of membershio of score to averagecriminal record set
	average = calcMembershipDegree(rating, minAverage, maxAverage)

	//get degree of membershio of score to good criminal record set
	good = calcMembershipDegree(rating, minGood, maxGood)

	// return the degree to which the user belarges to each linguistic variable set
	return map[string]float64{"bad": bad, "average": average, "good": good}
}

func (brw *Collateral) Fuzzify() uint8 {
	return 0
}
