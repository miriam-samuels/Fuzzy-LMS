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

type Employment struct {
	term uint8
}

// input to fuzzy inference system
type FISInput struct {
	CreditScore    map[string]float64
	Collateral     map[string]float64
	Income         map[string]float64
	CriminalRecord map[string]float64
	EmploymentTerm     map[string]float64
}

// output of fuzzy inference system
type FISOutput struct {
	Creditworthiness []float64
}

// fuzzy fules for inference
type FISRules struct {
	CreditScore, Collateral, Income, CriminalRecord, EmploymentTerm string
	// operator string // currently using and operator for now
}

type DefuzzifiedOutput struct {
	Creditworthiness string
}

// triangulare membership function
type TriangularMF struct {
	A, B, C float64 // vertices of a triangle
}

type TrapezoidalMF struct {
	A, B, C, D float64 // vertices of a trapezoid
}
