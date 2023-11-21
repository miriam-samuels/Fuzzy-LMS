package fis

// input variables
type CreditScore struct {
	Score uint16
}

type Collateral struct {
	HasCollateral bool
	Collateral    string
}

type Income struct {
	Amount float64
}

type Criminal struct {
	HasCriminalRec bool
	Offences       []string
}

type LoanAmount struct {
	Amount float64
}

// input to fuzzy inference system
type FISInput struct {
	CreditScore    map[string]float64
	Collateral     map[string]float64
	Income         map[string]float64
	CriminalRecord map[string]float64
	LoanAmount     map[string]float64
}

// output of fuzzy inference system
type FISOutput struct {
	Creditworthiness []float64
}

// fuzzy fules for inference
type FISRules struct {
	CreditScore      string
	Collateral       string
	Income           string
	CriminalRecord   string
	LoanAmount       string
	Creditworthiness string
	// operator string // currently using and operator for now
}

type DefuzzifiedOutput struct {
	Creditworthiness string
}
